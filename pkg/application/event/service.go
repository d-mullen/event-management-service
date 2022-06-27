package event

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"github.com/zenoss/event-management-service/pkg/domain/eventts"
	"github.com/zenoss/event-management-service/pkg/domain/scopes"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zingo/v4/interval"

	"github.com/ryboe/q"
	"github.com/zenoss/event-management-service/internal/frequency"
	"github.com/zenoss/event-management-service/pkg/domain/event"
)

type Service interface {
	Add(context.Context, *event.Event) (*event.Event, error)
	Find(context.Context, *event.Query) (*event.Page, error)
	Get(context.Context, *event.GetRequest) ([]*event.Event, error)
	Count(context.Context, *event.CountRequest) (*eventts.CountResponse, error)
	Frequency(context.Context, *event.FrequencyRequest) (*eventts.FrequencyResponse, error)
}

type service struct {
	eventContext event.Repository
	eventTS      eventts.Repository
	entityScopes scopes.EntityScopeProvider
}

func NewService(
	r event.Repository,
	eventTSRepo eventts.Repository,
	entityScopeProvider scopes.EntityScopeProvider) Service {
	return &service{eventContext: r, eventTS: eventTSRepo, entityScopes: entityScopeProvider}
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
	log := zenkit.ContextLogger(ctx)
	occurrenceMap, occInputMap := eventResultsToOccurrenceMaps(results)
	log.Tracef("occurrenceInputMap: %#v", occInputMap)
	occurrencesWithMD, err := eventTS.Get(ctx, &eventts.GetRequest{
		EventTimeseriesInput: eventts.EventTimeseriesInput{
			ByOccurrences: struct {
				OccurrenceMap map[string][]*eventts.OccurrenceInput
			}{
				OccurrenceMap: occInputMap,
			},
		},
	})
	log.WithField("occurrencesWithMetadata", occurrencesWithMD).Trace("got results from event-ts-svc")
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

func getEventTSFrequency(ctx context.Context,
	freqReq *event.FrequencyRequest,
	results []*event.Event,
	eventTS eventts.Repository) (*eventts.FrequencyResponse, error) {
	log := zenkit.ContextLogger(ctx)
	latest := 1
	if freqReq.CountInstances {
		latest = 0
	}
	_, occurrenceInputMap := eventResultsToOccurrenceMaps(results)
	req := &eventts.FrequencyRequest{
		EventTimeseriesInput: eventts.EventTimeseriesInput{
			ByOccurrences: struct {
				OccurrenceMap map[string][]*eventts.OccurrenceInput
			}{
				OccurrenceMap: occurrenceInputMap,
			},
			Latest: uint64(latest),
		},
		Fields:        freqReq.Fields,
		GroupBy:       freqReq.GroupBy,
		Downsample:    freqReq.Downsample,
		PersistCounts: freqReq.PersistCounts,
	}
	if freqReq.Filter != nil {
		filters, err := EventFilterToEventTSFilter(freqReq.Filter)
		if err != nil {
			log.WithError(err).Warn("failed to convert filters")
		}
		req.Filters = filters
	}
	resp, err := eventTS.Frequency(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func TransformScopeFilter(ctx context.Context, filter *event.Filter, provider scopes.EntityScopeProvider) (*event.Filter, bool, error) {
	log := zenkit.ContextLogger(ctx)
	if filter == nil {
		return nil, false, errors.New("got nil filter")
	}
	if provider == nil {
		return nil, false, errors.New("got nil provider")
	}
	if filter.Op != event.FilterOpScope {
		return nil, false, nil
	}
	if scope, ok := filter.Value.(*event.Scope); ok {
		eventIds, err := provider.GetEntityIDs(ctx, scope.Cursor)
		if err != nil {
			log.WithError(err).Error("failed to get entity ids")
			return nil, false, err
		}
		return &event.Filter{
			Field: "entity",
			Op:    event.FilterOpIn,
			Value: eventIds,
		}, true, nil
	}
	return nil, false, nil
}

func (svc *service) Find(ctx context.Context, query *event.Query) (*event.Page, error) {
	q.Q("service.Find: got query", query)
	err := applyScopeFilterTransform(ctx, query, svc.entityScopes)
	if err != nil {
		return nil, err
	}
	q.Q(query)
	results, err := svc.eventContext.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if len(results.Results) > 0 && svc.eventTS != nil && shouldGetEventTSDetails(query) {
		err = getEventTSResults(ctx, results.Results, svc.eventTS)
		if err != nil {
			return nil, err
		}
	}
	return results, nil
}

func applyScopeFilterTransform(ctx context.Context, query *event.Query, provider scopes.EntityScopeProvider) error {
	if query.Filter != nil {
		if query.Filter.Op != event.FilterOpAnd {
			inFilter, ok, err := TransformScopeFilter(ctx, query.Filter, provider)
			if err != nil {
				return err
			}
			q.Q("service.Find: called TransformScopeFilter", query.Filter, ok, inFilter)
			if ok {
				query.Filter = inFilter
			}
		} else {
			filterValue := query.Filter.Value
			if filters, ok := filterValue.([]*event.Filter); ok {
				newFilters := make([]*event.Filter, len(filters))
				for i, filter := range filters {
					inFilter, didTransform, err := TransformScopeFilter(ctx, filter, provider)
					if err != nil {
						return err
					}
					q.Q("service.Find: called TransformScopeFilter", query.Filter, ok, inFilter)
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

func shouldGetEventTSCounts(req *event.FrequencyRequest) bool {
	if req.CountInstances {
		return true
	}
	for _, field := range req.Fields {
		if !event.IsSupportedField(field) {
			return true
		}
	}
	for _, groupBy := range req.GroupBy {
		if !event.IsSupportedField(groupBy) {
			return true
		}
	}
	filters, _ := EventFilterToEventTSFilter(req.Filter)
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
	err := applyScopeFilterTransform(ctx, &req.Query, svc.entityScopes)
	if err != nil {
		log.WithError(err).Error("failed to get frequency")
		return nil, err
	}
	resp, err := svc.eventContext.Find(ctx, &req.Query)
	if err != nil {
		return nil, err
	}
	if !shouldGetEventTSCounts(req) {
		freqMap := frequency.NewFrequencyMap(req.GroupBy, req.TimeRange.Start, req.TimeRange.End, req.Downsample, req.PersistCounts)
		for _, result := range resp.Results {
			for _, occurrence := range result.Occurrences {
				log.WithField("occurrence", occurrence).Trace("processing occurrence")
				occAsMap := toMap(occurrence)
				atOrAboveStartTime := interval.AtOrAbove(uint64(occurrence.StartTime))
				if occurrence.Status != event.StatusClosed && occurrence.EndTime > 0 { // occurrence is active
					atOrAboveStartTime = interval.Closed(uint64(occurrence.StartTime), uint64(occurrence.EndTime)) // occurrence is closed
				}
				for _, field := range req.Fields {
					normalizedField := strings.ToLower(field)
					if v, ok := occAsMap[normalizedField]; ok {
						freqMap.Put(atOrAboveStartTime, map[string][]any{normalizedField: {v}})
					} else if occurrence.Metadata != nil {
						if values, ok2 := occurrence.Metadata[normalizedField]; ok2 {
							freqMap.Put(atOrAboveStartTime, map[string][]any{normalizedField: values})
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
	if len(resp.Results) > 0 && svc.eventTS != nil {
		freqResp, err := getEventTSFrequency(ctx, req, resp.Results, svc.eventTS)
		if err != nil {
			return nil, err
		}
		return freqResp, nil
	}
	return &eventts.FrequencyResponse{}, nil
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
	err := applyScopeFilterTransform(ctx, &req.Query, svc.entityScopes)
	if err != nil {
		log.WithError(err).Error("failed to get count")
		return nil, err
	}
	resp, err := svc.eventContext.Find(ctx, &req.Query)
	if err != nil {
		return nil, err
	}
	if !shouldGetEventTSCounts(&event.FrequencyRequest{
		Query:  req.Query,
		Fields: req.Fields,
	}) {
		results := make(countResult)
		for _, result := range resp.Results {
			for _, occurrence := range result.Occurrences {
				log.WithField("occurrence", occurrence).Trace("processing occurrence")
				occAsMap := toMap(occurrence)
				for _, field := range req.Fields {
					normalizedField := strings.ToLower(field)
					result, ok := results[normalizedField]
					if !ok {
						result = make(map[fieldValueKey]uint64)
						results[normalizedField] = result
					}
					if v, ok := occAsMap[normalizedField]; ok {
						valueKey := fieldValueKey{field: normalizedField, value: v}
						cnt := result[valueKey]
						result[valueKey] = cnt + 1
					} else if occurrence.Metadata != nil {
						if values, ok2 := occurrence.Metadata[normalizedField]; ok2 {
							for _, v := range values {
								valueKey := fieldValueKey{field: normalizedField, value: v}
								cnt := result[valueKey]
								result[valueKey] = cnt + 1
							}
						}
					}
				}
			}
		}
		resp := &eventts.CountResponse{
			Results: make(map[string][]*eventts.CountResult),
		}
		for field, countResult := range results {
			results := make([]*eventts.CountResult, 0)
			for vk, counts := range countResult {
				results = append(results, &eventts.CountResult{
					Value: vk.value,
					Count: counts,
				})
			}
			resp.Results[field] = results
		}
		return resp, nil
	}
	return nil, errors.New("unimplemented")
}
