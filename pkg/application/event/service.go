package event

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/zenoss/event-management-service/pkg/models/eventts"
	"github.com/zenoss/event-management-service/pkg/models/scopes"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zingo/v4/interval"

	"github.com/zenoss/event-management-service/internal/frequency"
	"github.com/zenoss/event-management-service/pkg/models/event"
)

const (
	defaultFrequencyDownsampleRate = int64(5 * 60 * 60 * 1e3)
)

type Service interface {
	Add(context.Context, *event.Event) (*event.Event, error)
	Find(context.Context, *event.Query) (*event.Page, error)
	Get(context.Context, *event.GetRequest) ([]*event.Event, error)
	Count(context.Context, *event.CountRequest) (*eventts.CountResponse, error)
	Frequency(context.Context, *event.FrequencyRequest) (*eventts.FrequencyResponse, error)
}

type service struct {
	eventContext   event.Repository
	eventTS        eventts.Repository
	entityScopes   scopes.EntityScopeProvider
	activeEntities scopes.ActiveEntityRepository
}

func NewService(
	r event.Repository,
	eventTSRepo eventts.Repository,
	entityScopeProvider scopes.EntityScopeProvider,
	activeEntityAdapter scopes.ActiveEntityRepository) Service {
	return &service{
		eventContext:   r,
		eventTS:        eventTSRepo,
		entityScopes:   entityScopeProvider,
		activeEntities: activeEntityAdapter}
}

func (svc *service) Add(_ context.Context, _ *event.Event) (*event.Event, error) {
	panic("not implemented") // TODO: Implement
}

func eventResultsToOccurrenceMaps(
	results []*event.Event) (map[string]*event.Occurrence, map[string][]*eventts.OccurrenceInput) {

	occurrenceMap := make(map[string]*event.Occurrence)
	occInputMap := make(map[string][]*eventts.OccurrenceInput)
	for _, eventResult := range results {
		occInputSlice := make([]*eventts.OccurrenceInput, len(eventResult.Occurrences))
		for i, occ := range eventResult.Occurrences {
			occInputSlice[i] = &eventts.OccurrenceInput{
				ID:      occ.ID,
				EventID: occ.EventID,
				TimeRange: eventts.TimeRange{
					Start: occ.StartTime,
					End:   occ.EndTime,
				},
				IsActive: occ.Status != event.StatusClosed,
			}
			occurrenceMap[occ.ID] = occ
		}
		occInputMap[eventResult.ID] = occInputSlice
	}
	return occurrenceMap, occInputMap
}

func getEventTSResults(ctx context.Context, results []*event.Event, eventTS eventts.Repository) error {
	occurrenceMap, occInputMap := eventResultsToOccurrenceMaps(results)
	occurrencesWithMD, err := eventTS.Get(ctx, &eventts.GetRequest{
		EventTimeseriesInput: eventts.EventTimeseriesInput{
			ByOccurrences: struct {
				OccurrenceMap map[string][]*eventts.OccurrenceInput
			}{
				OccurrenceMap: occInputMap,
			},
		},
	})
	if err != nil {
		return err
	}
	for _, occurrenceWithMD := range occurrencesWithMD {
		occResult, ok := occurrenceMap[occurrenceWithMD.ID]
		if !ok {
			occID, didMatch := matchUnassignedOccurrences(occurrenceWithMD, occInputMap)
			if !didMatch || len(occID) == 0 {
				continue
			}
			ok2 := false
			occResult, ok2 = occurrenceMap[occID]
			if !ok2 {
				continue
			}
		}
		if len(occResult.Metadata) == 0 {
			occResult.Metadata = occurrenceWithMD.Metadata
			continue
		}
		for k, v := range occurrenceWithMD.Metadata {
			occResult.Metadata[k] = v
		}
	}
	return nil
}

func matchUnassignedOccurrences(
	occurrenceWithMetadata *eventts.Occurrence,
	eventToOccurrenceInputMap map[string][]*eventts.OccurrenceInput,
) (string, bool) {
	occurrenceInputs, ok := eventToOccurrenceInputMap[occurrenceWithMetadata.EventID]
	if !ok {
		return "", false
	}
	for _, occInput := range occurrenceInputs {
		ts := occurrenceWithMetadata.Timestamp
		if occInput.IsActive && occInput.TimeRange.Start <= ts ||
			occInput.TimeRange.Start <= ts && ts <= occInput.TimeRange.End {
			return occInput.ID, true
		}
	}
	return "", false
}

func TransformScopeFilter(ctx context.Context, timeRange event.TimeRange, tenantID string, filter *event.Filter, entityScopeProvider scopes.EntityScopeProvider, activeEntityAdapter scopes.ActiveEntityRepository) (*event.Filter, bool, error) {
	log := zenkit.ContextLogger(ctx)
	if filter == nil {
		return nil, false, errors.New("got nil filter")
	}
	if entityScopeProvider == nil {
		return nil, false, errors.New("got nil provider")
	}
	if filter.Op != event.FilterOpScope {
		return nil, false, nil
	}
	if scope, ok := filter.Value.(*event.Scope); ok {
		entityIDs, err := entityScopeProvider.GetEntityIDs(ctx, scope.Cursor)
		if err != nil {
			log.WithError(err).Error("failed to get entity ids")
			return nil, false, err
		}
		if activeEntityAdapter != nil {
			entityIDs2, err := activeEntityAdapter.Get(ctx, timeRange, tenantID, entityIDs)
			if err != nil {
				log.WithField(logrus.ErrorKey, err).Warn("failed to get active entity ids")
			} else {
				entityIDs = entityIDs2
			}
		}
		return &event.Filter{
			Field: "entity",
			Op:    event.FilterOpIn,
			Value: entityIDs,
		}, true, nil
	}
	return nil, false, nil
}

func (svc *service) Find(ctx context.Context, query *event.Query) (*event.Page, error) {
	var (
		resultMut sync.Mutex
		results   = new(event.Page)
	)
	err := applyScopeFilterTransform(ctx, query, svc.entityScopes, svc.activeEntities)
	if err != nil {
		return nil, err
	}
	if query.PageInput == nil || (query.PageInput != nil && query.PageInput.Limit == 0 && len(query.PageInput.Cursor) == 0 && len(query.PageInput.SortBy) == 0) {
		queryGroup, gCtx := errgroup.WithContext(ctx)
		queryGroup.SetLimit(5)
		queries := SplitOutQueries(2000, "entity", query)
		for _, query := range queries {
			_q := query
			queryGroup.Go(func() error {
				currPage, err := svc.eventContext.Find(gCtx, _q)
				if err != nil {
					return err
				}
				resultMut.Lock()
				results.Results = append(results.Results, currPage.Results...)
				resultMut.Unlock()
				return nil
			})
		}
		if err = queryGroup.Wait(); err != nil {
			return nil, err
		}
	} else {
		results, err = svc.eventContext.Find(ctx, query)
		if err != nil {
			return nil, err
		}
	}
	if len(results.Results) > 0 && svc.eventTS != nil && shouldGetEventTSDetails(query) {
		err = getEventTSResults(ctx, results.Results, svc.eventTS)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func applyScopeFilterTransform(ctx context.Context, query *event.Query, provider scopes.EntityScopeProvider, activeEntityAdapter scopes.ActiveEntityRepository) error {
	if query.Filter != nil {
		if query.Filter.Op != event.FilterOpAnd && query.Filter.Op != event.FilterOpOr && query.Filter.Op != event.FilterOpNot {
			inFilter, ok, err := TransformScopeFilter(ctx, query.TimeRange, query.Tenant, query.Filter, provider, activeEntityAdapter)
			if err != nil {
				return err
			}
			if ok {
				query.Filter = inFilter
			}
		} else if query.Filter.Op == event.FilterOpNot {
			if filter, ok := query.Filter.Value.(*event.Filter); ok {
				inFilter, didTransform, err := TransformScopeFilter(ctx, query.TimeRange, query.Tenant, filter, provider, activeEntityAdapter)
				if err != nil {
					return err
				}
				if didTransform {
					query.Filter.Value = inFilter
				}
			}
		} else {
			filterValue := query.Filter.Value
			if filters, ok := filterValue.([]*event.Filter); ok {
				newFilters := make([]*event.Filter, len(filters))
				for i, filter := range filters {
					inFilter, didTransform, err := TransformScopeFilter(ctx, query.TimeRange, query.Tenant, filter, provider, activeEntityAdapter)
					if err != nil {
						return err
					}
					if didTransform {
						newFilters[i] = inFilter
					} else {
						newFilters[i] = filter
					}
				}
				query.Filter.Value = newFilters
			}
		}
	}
	return nil
}

func (svc *service) Get(ctx context.Context, req *event.GetRequest) ([]*event.Event, error) {
	results, err := svc.eventContext.Get(ctx, req)
	if err != nil {
		return nil, err
	}
	if len(results) > 0 && svc.eventTS != nil {
		err = getEventTSResults(ctx, results, svc.eventTS)
		if err != nil {
			return nil, err
		}

	}
	return results, nil
}

func shouldGetEventTSDetails(query *event.Query) bool {
	for _, field := range query.Fields {
		if len(field) > 0 && !event.IsSupportedField(field) {
			return true
		}
	}
	filters, _ := EventFilterToEventTSFilter(query.Filter)
	for _, filter := range filters {
		if len(filter.Field) > 0 && !event.IsSupportedField(filter.Field) {
			return true
		}
	}
	return false
}

func toMap(src any) map[string]any {
	results := make(map[string]any)
	b, err := json.Marshal(src)
	if err != nil {
		return results
	}
	_ = json.Unmarshal(b, &results)
	return results
}

func bucketsToFrequencyResult(buckets []*frequency.Bucket) []*eventts.FrequencyResult {
	results := make([]*eventts.FrequencyResult, len(buckets))
	for i, bucket := range buckets {
		results[i] = &eventts.FrequencyResult{
			Key:    bucket.Key,
			Values: bucket.Values,
		}
	}
	return results
}

func (svc *service) Frequency(ctx context.Context, req *event.FrequencyRequest) (*eventts.FrequencyResponse, error) {
	log := zenkit.ContextLogger(ctx)
	// construct query with combined fields from query and frequency request
	requestedFields := make([]string, 0)
	requestedFields = append(requestedFields, req.Fields...)
	requestedFields = append(requestedFields, req.GroupBy...)
	requestedFields = append(requestedFields, req.Query.Fields...)
	updatedQuery := &event.Query{
		Tenant:     req.Query.Tenant,
		TimeRange:  req.TimeRange,
		Severities: req.Severities,
		Statuses:   req.Statuses,
		Fields:     requestedFields,
		Filter:     req.Query.Filter,
		PageInput:  req.Query.PageInput,
	}
	resp, err := svc.Find(ctx, updatedQuery)
	if err != nil {
		log.WithField(logrus.ErrorKey, err).Error("failed to perform frequency")
		return nil, err
	}
	downsample := defaultFrequencyDownsampleRate
	if req.Downsample > 0 {
		downsample = req.Downsample
	}
	freqMap := frequency.NewFrequencyMap(req.GroupBy, req.TimeRange.Start, req.TimeRange.End, downsample, req.PersistCounts)
	for _, result := range resp.Results {
		for _, occurrence := range result.Occurrences {
			log.WithField("occurrence", occurrence).Trace("processing occurrence")
			occAsMap := toMap(occurrence)
			atOrAboveStartTime := interval.AtOrAbove(uint64(occurrence.StartTime))
			if occurrence.Status != event.StatusClosed && occurrence.EndTime > 0 { // occurrence is active
				atOrAboveStartTime = interval.Closed(uint64(occurrence.StartTime), uint64(occurrence.EndTime)) // occurrence is closed
			}
			for _, field := range req.Fields {
				foundFieldValue := false
				if v, ok := occAsMap[field]; ok {
					freqMap.Put(atOrAboveStartTime, map[string][]any{field: {v}})
					foundFieldValue = true
				} else if occurrence.Metadata != nil {
					if values, ok2 := occurrence.Metadata[field]; ok2 {
						freqMap.Put(atOrAboveStartTime, map[string][]any{field: values})
						foundFieldValue = true
					}
				}
				if !foundFieldValue && result.Dimensions != nil {
					if dim, ok := result.Dimensions[field]; ok {
						freqMap.Put(atOrAboveStartTime, map[string][]any{field: {dim}})
					}
				}
			}
		}
	}
	timestamps, buckets := freqMap.Get()
	return &eventts.FrequencyResponse{
		Timestamps: timestamps,
		Results:    bucketsToFrequencyResult(buckets),
	}, nil
}

type (
	fieldValueKey struct {
		field string
		value any
	}
	countResult map[string]map[fieldValueKey]uint64
)

func (svc *service) Count(ctx context.Context, req *event.CountRequest) (*eventts.CountResponse, error) {
	log := zenkit.ContextLogger(ctx)
	// construct query with combined fields from query and count request
	updatedQuery := &event.Query{
		Tenant:     req.Query.Tenant,
		TimeRange:  req.TimeRange,
		Severities: req.Severities,
		Statuses:   req.Statuses,
		Fields:     append(req.Fields, req.Fields...),
		Filter:     req.Query.Filter,
		PageInput:  req.Query.PageInput,
	}
	resp, err := svc.Find(ctx, updatedQuery)
	if err != nil {
		log.WithField(logrus.ErrorKey, err).Error("failed to perform count")
		return nil, err
	}
	countResults := make(countResult)
	for _, eventResult := range resp.Results {
		for _, occurrence := range eventResult.Occurrences {
			log.WithField("occurrence", occurrence).Trace("processing occurrence")
			occAsMap := toMap(occurrence)
			for _, field := range req.Fields {
				countResult, ok := countResults[field]
				if !ok {
					countResult = make(map[fieldValueKey]uint64)
					countResults[field] = countResult
				}
				foundFieldValue := false
				if v, ok := occAsMap[field]; ok {
					foundFieldValue = true
					valueKey := fieldValueKey{field: field, value: v}
					cnt := countResult[valueKey]
					countResult[valueKey] = cnt + 1
				} else if occurrence.Metadata != nil {
					if values, ok2 := occurrence.Metadata[field]; ok2 {
						foundFieldValue = true
						for _, v := range values {
							valueKey := fieldValueKey{field: field, value: v}
							cnt := countResult[valueKey]
							countResult[valueKey] = cnt + 1
						}
					}
				}
				if !foundFieldValue && eventResult.Dimensions != nil {
					if dim, ok := eventResult.Dimensions[field]; ok {
						valueKey := fieldValueKey{field: field, value: dim}
						cnt := countResult[valueKey]
						countResult[valueKey] = cnt + 1

					}
				}
			}
		}
	}
	countResp := &eventts.CountResponse{
		Results: make(map[string][]*eventts.CountResult),
	}
	for field, countResult := range countResults {
		results := make([]*eventts.CountResult, 0)
		for vk, counts := range countResult {
			results = append(results, &eventts.CountResult{
				Value: vk.value,
				Count: counts,
			})
		}
		countResp.Results[field] = results
	}
	return countResp, nil
}
