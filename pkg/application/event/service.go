package event

import (
	"context"
	"encoding/json"
	"math"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/pkg/models/eventts"
	"github.com/zenoss/event-management-service/pkg/models/scopes"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zingo/v4/interval"

	"github.com/zenoss/event-management-service/internal"
	"github.com/zenoss/event-management-service/internal/batchops"
	"github.com/zenoss/event-management-service/internal/frequency"
	"github.com/zenoss/event-management-service/internal/instrumentation"
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
	results []*event.Event) (map[string]*event.Occurrence,
	map[string][]*eventts.OccurrenceInput, map[string]*event.Event) {

	occurrenceMap := make(map[string]*event.Occurrence)
	occInputMap := make(map[string][]*eventts.OccurrenceInput)
	eventMap := make(map[string]*event.Event)
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
			eventMap[occ.ID] = eventResult
		}
		occInputMap[eventResult.ID] = occInputSlice
	}
	return occurrenceMap, occInputMap, eventMap
}

func getOccurrenceDetails(ctx context.Context, origOccurs []*event.Occurrence, q *event.Query, eventTSRepo eventts.Repository) ([]*event.Occurrence, error) {
	var (
		err                       error
		log                       = zenkit.ContextLogger(ctx)
		eventIDs                  = make([]string, 0)
		occurrenceInputsByEventID = make(map[string][]*eventts.OccurrenceInput)
		minStart, maxEnd          = int64(math.MaxInt64), int64(math.MinInt64)
		tsFilters                 = make([]*eventts.Filter, 0)
	)
	if !shouldGetEventTSDetails(q) {
		return origOccurs, nil
	}
	for _, occ := range origOccurs {
		input := &eventts.OccurrenceInput{
			ID:      occ.ID,
			EventID: occ.EventID,
			TimeRange: eventts.TimeRange{
				Start: occ.StartTime,
				End:   occ.EndTime,
			},
			IsActive: occ.Status != event.StatusClosed,
		}
		if input.TimeRange.Start < minStart {
			minStart = input.TimeRange.Start
		}
		if input.TimeRange.End > maxEnd {
			maxEnd = input.TimeRange.End
		}
		occSlice, ok := occurrenceInputsByEventID[occ.EventID]
		if !ok {
			occSlice = []*eventts.OccurrenceInput{input}
			eventIDs = append(eventIDs, occ.EventID)
		} else {
			occSlice = append(occSlice, input)
		}
		occurrenceInputsByEventID[occ.EventID] = occSlice
	}
	if maxEnd == 0 {
		maxEnd = time.Now().UnixMilli()
	}
	tsFilters, err = eventFilterToEventTSFilter(q.Filter)
	if err != nil {
		log.WithError(err).Error("failed to covert to event-ts filter")
	}

	if len(q.Severities) > 0 { // severity filter
		tsSeverities := make([]any, len(q.Severities))
		for i, sev := range q.Severities {
			tsSeverities[i] = event.Severity_name[sev]
		}
		tsFilters = append(tsFilters, &eventts.Filter{
			Operation: eventts.Operation_OP_IN,
			Field:     "_zv_severity",
			Values:    tsSeverities})
	}

	if len(q.Statuses) > 0 { // status filter
		tsStatuses := make([]any, len(q.Statuses))
		for i, status := range q.Statuses {
			tsStatuses[i] = event.Status_name[status]
		}
		tsFilters = append(tsFilters, &eventts.Filter{
			Operation: eventts.Operation_OP_IN,
			Field:     "_zv_status",
			Values:    tsStatuses})
	}

	req := &eventts.GetRequest{
		EventTimeseriesInput: eventts.EventTimeseriesInput{
			TimeRange: eventts.TimeRange{
				Start: q.TimeRange.Start,
				End:   q.TimeRange.End,
			},
			ByEventIDs: struct{ IDs []string }{
				IDs: eventIDs,
			},
			ByOccurrences: struct {
				ShouldApplyIntervals bool
				OccurrenceMap        map[string][]*eventts.OccurrenceInput
			}{
				ShouldApplyIntervals: q.ShouldApplyOccurrenceIntervals,
				OccurrenceMap:        occurrenceInputsByEventID,
			},
			Latest:       1,
			ResultFields: q.Fields,
			Filters:      tsFilters,
		},
	}
	if q.ShouldApplyOccurrenceIntervals {
		req.TimeRange.Start = minStart
		req.TimeRange.End = maxEnd
	}
	results := make([]*event.Occurrence, 0, len(origOccurs))
	occurrencesByEventId := make(map[string][]*eventts.Occurrence)
	stream := eventTSRepo.GetStream(ctx, req)
	var streamErr error
	ticker := time.NewTicker(2 * time.Second)
StreamLoop:
	for {
		select {
		case resp := <-stream:
			if resp == nil {
				continue
			}
			if resp.Err != nil {
				log.WithError(resp.Err).Warn("got error during event-ts scan")
				streamErr = resp.Err
				break StreamLoop
			} else if resp.Result != nil {
				log.Debugf("got event-ts result: %s %s %d\n", resp.Result.ID, resp.Result.EventID, resp.Result.Timestamp)
				occSlice, ok := occurrencesByEventId[resp.Result.EventID]
				if !ok || occSlice == nil {
					occSlice = make([]*eventts.Occurrence, 0)
				}
				occSlice = append(occSlice, resp.Result)
				occurrencesByEventId[resp.Result.EventID] = occSlice
				ticker.Reset(2 * time.Second)
			}
		case <-ticker.C:
			break StreamLoop
		case <-ctx.Done():
			ctxErr := errors.Unwrap(ctx.Err())
			switch ctxErr {
			case nil:
				// no op
			case context.DeadlineExceeded:
				log.Debug("deadline exceeded during event-ts scan")
				break StreamLoop
			default:
				log.WithError(ctxErr).Error("got error during event-ts scan")
				streamErr = ctxErr
				break StreamLoop
			}
		}
	}
	if streamErr != nil {
		return nil, errors.Wrap(streamErr, "failed to stream event time-series data")
	}
	for _, occ := range origOccurs {
		if occSlice, ok := occurrencesByEventId[occ.EventID]; ok {
		InnerLoop:
			for _, occWithMD := range occSlice {
				if occ.ID == occWithMD.ID ||
					occ.Status == event.StatusClosed && occ.StartTime <= occWithMD.Timestamp && occWithMD.Timestamp <= occ.EndTime ||
					occ.Status != event.StatusClosed && occ.StartTime <= occWithMD.Timestamp {
					if len(occ.Metadata) == 0 {
						occ.Metadata = occWithMD.Metadata
					} else {
						for k, v := range occWithMD.Metadata {
							occ.Metadata[k] = append(occ.Metadata[k], v...)
						}
					}
					if _, ok3 := occ.Metadata["lastSeen"]; !ok3 {
						occ.Metadata["lastSeen"] = []any{occWithMD.Timestamp}
					}
					results = append(results, occ)
					continue InnerLoop
				}
			}
		}
	}
	return results, nil
}

func makeEventTSOccurrenceProcessor(q *event.Query, eventTSRepo eventts.Repository) event.OccurrenceProcessor {
	return func(ctx context.Context, origOccurrences []*event.Occurrence) ([]*event.Occurrence, error) {
		ctx, span := instrumentation.StartSpan(ctx, "occurrenceProcessor/getDetails")
		defer span.End()
		results := make([]*event.Occurrence, 0)
		resultMut := sync.Mutex{}
		batchops.DoConcurrently(ctx, 50, 25,
			origOccurrences,
			func(ctx context.Context, batch []*event.Occurrence) ([]*event.Occurrence, error) {
				batchResult, err := getOccurrenceDetails(ctx, batch, q, eventTSRepo)
				if err != nil {
					return nil, errors.Wrap(err, "failed to get occurrence time-series details")
				}
				return batchResult, nil
			},
			func(ctx context.Context, batch []*event.Occurrence) (bool, error) {
				resultMut.Lock()
				defer resultMut.Unlock()
				results = append(results, batch...)
				return true, nil
			},
		)
		return results, nil
	}
}

func getEventTSResults(ctx context.Context, results []*event.Event, eventTS eventts.Repository, query *event.Query) error {
	occurrenceMap, occInputMap, _ := eventResultsToOccurrenceMaps(results)
	var (
		shouldApplyOccurrenceIntervals bool
		tsFilters                      []*eventts.Filter
		tsFields                       []string
	)
	if query != nil {
		tsFilters, _ = eventFilterToEventTSFilter(query.Filter)
		tsFields = query.Fields

		if len(query.Severities) > 0 { // severity filter
			tsSeverities := make([]any, len(query.Severities))
			for i, sev := range query.Severities {
				tsSeverities[i] = int(sev)
			}
			tsFilters = append(tsFilters, &eventts.Filter{
				Operation: eventts.Operation_OP_IN,
				Field:     "_zv_severity",
				Values:    tsSeverities})
		}

		if len(query.Statuses) > 0 { // status filter
			tsStatuses := make([]any, len(query.Statuses))
			for i, status := range query.Statuses {
				tsStatuses[i] = int(status)
			}
			tsFilters = append(tsFilters, &eventts.Filter{
				Operation: eventts.Operation_OP_IN,
				Field:     "_zv_status",
				Values:    tsStatuses})
		}
		shouldApplyOccurrenceIntervals = query.ShouldApplyOccurrenceIntervals
	}

	occurrencesWithMD, err := eventTS.Get(ctx, &eventts.GetRequest{
		EventTimeseriesInput: eventts.EventTimeseriesInput{
			ByOccurrences: struct {
				ShouldApplyIntervals bool
				OccurrenceMap        map[string][]*eventts.OccurrenceInput
			}{
				ShouldApplyIntervals: shouldApplyOccurrenceIntervals,
				OccurrenceMap:        occInputMap,
			},
			Filters:      tsFilters,
			ResultFields: tsFields,
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
		entityIDs, err := entityScopeProvider.GetEntityIDs(ctx, scope.Cursor, timeRange)
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
		resp      = new(event.Page)
	)
	err := applyScopeFilterTransform(ctx, query, svc.entityScopes, svc.activeEntities)
	if err != nil {
		return nil, err
	}
	eventContextQ := &event.Query{
		ShouldApplyOccurrenceIntervals: query.ShouldApplyOccurrenceIntervals,
		Tenant:                         query.Tenant,
		TimeRange:                      query.TimeRange,
		Severities:                     internal.CloneSlice(query.Severities),
		Statuses:                       internal.CloneSlice(query.Statuses),
		Fields:                         internal.CloneSlice(query.Fields),
		Filter:                         query.Filter.Clone(),
		PageInput:                      query.PageInput,
	}
	// drop any filters on unsupported fields in event context store query
	newFilter, err := transformFilter(query.Filter, func(f *event.Filter) (*event.Filter, bool, error) {
		if event.IsSupportedField(f.Field) {
			return f, true, nil
		}
		return nil, false, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to filter out filters with unsupported fields")
	}
	if newFilter != nil {
		eventContextQ.Filter = newFilter
	}
	occProcessor := makeEventTSOccurrenceProcessor(query, svc.eventTS)
	findOpt := event.NewFindOption()
	findOpt.SetOccurrenceProcessors([]event.OccurrenceProcessor{occProcessor})
	_ = findOpt
	if query.PageInput == nil || (query.PageInput != nil && query.PageInput.Limit == 0 && len(query.PageInput.Cursor) == 0 && len(query.PageInput.SortBy) == 0) {
		queryGroup, gCtx := errgroup.WithContext(ctx)
		queryGroup.SetLimit(5)
		queries := SplitOutQueries(2000, "entity", eventContextQ)
		for _, query := range queries {
			_q := query
			queryGroup.Go(func() error {
				currPage, err := svc.eventContext.Find(gCtx, _q, findOpt)
				if err != nil {
					return err
				}
				resultMut.Lock()
				defer resultMut.Unlock()
				resp.Results = append(resp.Results, currPage.Results...)
				return nil
			})
		}
		if err = queryGroup.Wait(); err != nil {
			return nil, err
		}
	} else {
		resp, err = svc.eventContext.Find(ctx, eventContextQ, findOpt)
		if err != nil {
			return nil, err
		}
	}
	return resp, nil
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
		err = getEventTSResults(ctx, results, svc.eventTS, nil)
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
	filters, _ := eventFilterToEventTSFilter(query.Filter)
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
		Fields:     append(req.Fields, req.Query.Fields...),
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
