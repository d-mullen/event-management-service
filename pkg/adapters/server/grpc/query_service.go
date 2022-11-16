package grpc

import (
	"context"
	"errors"

	"github.com/zenoss/event-management-service/internal/auth"

	appEvent "github.com/zenoss/event-management-service/pkg/application/event"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zing-proto/v11/go/cloud/eventquery"
	eventPb "github.com/zenoss/zing-proto/v11/go/event"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	errGrpcNotImplemented = status.Error(codes.Unimplemented, "unimplemented")
)

type (
	EventQueryService struct {
		eventquery.UnimplementedEventQueryServiceServer
		svr appEvent.Service
	}
)

func NewEventQueryService(eventQuerySvr appEvent.Service) *EventQueryService {
	return &EventQueryService{
		svr: eventQuerySvr,
	}
}

var _ eventquery.EventQueryServiceServer = &EventQueryService{}

func sortOrderFromProto(order eventquery.SortOrder) event.SortOrder {
	if order == eventquery.SortOrder_SORT_ORDER_DESC {
		return event.SortOrderDescending
	}
	return event.SortOrderAscending
}
func fieldNameFromProto(byField *eventquery.SortBy_ByField) string {
	switch v := byField.ByField.SortField.(type) {
	case *eventquery.SortByField_Property:
		return v.Property
	case *eventquery.SortByField_Dimension:
		return v.Dimension
	case *eventquery.SortByField_Metadata:
		return v.Metadata
	default:
		return ""
	}
}

func ProtoSortByToDomain(orig *eventquery.SortBy) (*event.SortOpt, error) {
	switch v := orig.SortType.(type) {
	case *eventquery.SortBy_ByField:
		return &event.SortOpt{
			Field:     fieldNameFromProto(v),
			SortOrder: sortOrderFromProto(v.ByField.GetOrder()),
		}, nil
	case *eventquery.SortBy_BySeries:
		return &event.SortOpt{
			Field:     v.BySeries.SeriesName,
			SortOrder: sortOrderFromProto(v.BySeries.Order),
		}, nil
	default:
		return nil, status.Error(codes.InvalidArgument, "unsupported sort by type")
	}
}

func defaultSetQuery[T eventPb.Status | eventPb.Severity, E event.Status | event.Severity](values []T, max E) []E {
	var (
		retval []E = make([]E, 0)
		e      E
	)

	if len(values) == 0 {
		for e = 0; e <= max; e++ {
			retval = append(retval, E(e))
		}
	} else {
		for _, v := range values {
			retval = append(retval, E(v))
		}
	}
	return retval
}

func QueryProtoToEventQuery(tenantID string, query *eventquery.Query) (*event.Query, error) {
	if query == nil {
		return nil, status.Error(codes.InvalidArgument, "nil query found on request")
	}
	if tr := query.TimeRange; tr == nil || tr.End < tr.Start {
		return nil, status.Error(codes.InvalidArgument, "invalid time range")
	}
	result := &event.Query{
		Tenant:     tenantID,
		Severities: defaultSetQuery(query.Severities, event.SeverityMax),
		Statuses:   defaultSetQuery(query.Statuses, event.StatusMax),
		TimeRange: event.TimeRange{
			Start: query.TimeRange.Start,
			End:   query.TimeRange.End,
		},
		ShouldApplyOccurrenceIntervals: query.ActiveCriteria == eventquery.TemporalFilterCriteria_BY_OCCURRENCES,
	}

	if clause := query.GetClause(); clause != nil {
		filter, err := ClauseProtoToEventFilter(clause)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "failed to convert clause proto: %q", err)
		}
		result.Filter = filter
	}
	if len(query.Fields) > 0 {
		fields := make([]string, 0)
		for _, input := range query.Fields {
			fields = append(fields, input.Field)
		}
		result.Fields = fields
	}
	if len(query.SortBy) > 0 {
		pageInput := &event.PageInput{}
		sortBys := make([]event.SortOpt, 0)
		for _, sortBy := range query.SortBy {
			newSortBy, err := ProtoSortByToDomain(sortBy)
			if err == nil {
				sortBys = append(sortBys, *newSortBy)
			}
		}
		pageInput.SortBy = sortBys
		result.PageInput = pageInput
	}
	if query.PageInput != nil {
		if result.PageInput == nil {
			result.PageInput = &event.PageInput{}
		}
		result.PageInput.Cursor = query.PageInput.Cursor
		result.PageInput.Limit = uint64(query.PageInput.Limit)
	}
	return result, nil
}

func (handler *EventQueryService) Search(ctx context.Context, req *eventquery.SearchRequest) (*eventquery.SearchResponse, error) {
	ctx, span := trace.StartSpan(ctx, "zenoss.cloud/eventQuery/v2/go/EventQueryService.Search")
	defer span.End()
	log := zenkit.ContextLogger(ctx).WithField("req", req)
	ident, err := auth.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	findReq, err := QueryProtoToEventQuery(ident.TenantName(), req.Query)
	if err != nil {
		log.WithError(err).Error("failed to convert request to domain query")
		return nil, err
	}
	log.
		WithField("query", findReq).
		Tracef("executing Find request")
	page, err := handler.svr.Find(ctx, findReq)
	if err != nil {
		log.WithError(err).Errorf("failed to convert request to domain query")
		return nil, status.Errorf(codes.Unknown, "failed to execute search: %q", errors.Unwrap(err))
	}
	totalCount := uint64(0)
	if req.Query.PageInput != nil {
		if len(req.Query.PageInput.Cursor) == 0 {
			if !page.HasNext {
				totalCount = uint64(len(page.Results))
			}
		}
	}

	resp := &eventquery.SearchResponse{
		Results: make([]*eventquery.EventResult, 0),
		PageInfo: &eventquery.PageInfo{
			StartCursor: page.StartCursor,
			EndCursor:   page.EndCursor,
			HasNext:     page.HasNext,
			HasPrev:     page.HasPrev,
			Count:       uint64(len(page.Results)),
			TotalCount:  totalCount,
		},
	}
	for _, result := range page.Results {
		currID := result.ID
		currEntity := result.Entity
		currName := result.Name
		currDim, err := structpb.NewStruct(result.Dimensions)
		if err != nil {
			log.WithError(err).Error("failed to convert result dimensions to struct protocol buffer")
			return nil, err
		}
		curr := &eventquery.EventResult{
			Id:         currID,
			Entity:     currEntity,
			Dimensions: currDim,
			Name:       currName,
		}
		if len(result.Occurrences) > 0 {
			curr.Occurrences = make([]*eventquery.Occurrence, len(result.Occurrences))
			for i, occ := range result.Occurrences {
				curr.Occurrences[i] = &eventquery.Occurrence{
					Id:            occ.ID,
					EventId:       occ.EventID,
					TimeRange:     &eventquery.TimeRange{Start: occ.StartTime, End: occ.EndTime},
					Severity:      eventPb.Severity(occ.Severity), // TODO: fixme
					Status:        eventPb.Status(occ.Status),     // TODO: fixme
					InstanceCount: uint64(occ.InstanceCount),
				}
				currOcc := curr.Occurrences[i]
				if occ.Acknowledged != nil {
					currOcc.Acknowledged = &wrapperspb.BoolValue{Value: *occ.Acknowledged}
				}
				if len(occ.Notes) > 0 {
					currOcc.Notes = make([]*eventquery.Note, len(occ.Notes))
					for j, thisNote := range occ.Notes {
						currOcc.Notes[j] = &eventquery.Note{
							Id:        thisNote.ID,
							Content:   thisNote.Content,
							CreatedBy: thisNote.CreatedBy,
							CreatedAt: thisNote.CreatedAt.UnixMilli(),
							UpdatedAt: thisNote.UpdatedAt.UnixMilli(),
						}
					}
				}
				if len(occ.Metadata) > 0 {
					resultMD := &structpb.Struct{}
					resultMD.Fields = make(map[string]*structpb.Value)
					for k, anyValues := range occ.Metadata {
						listValue, err := structpb.NewList(anyValues)
						if err != nil {
							return nil, status.Errorf(codes.Internal, "failed to convert metatadata to protobuf: %q", err)
						}
						// Insure that we are using event context lastSeen value.
						if k == "lastSeen" {
							listValue, _ = structpb.NewList([]any{occ.LastSeen})
						}
						resultMD.Fields[k] = &structpb.Value{
							Kind: &structpb.Value_ListValue{
								ListValue: listValue,
							},
						}
					}
					currOcc.Metadata = resultMD
				}
			}
		}
		resp.Results = append(resp.Results, curr)

	}
	return resp, nil
}

func (handler *EventQueryService) GetEvent(_ context.Context, _ *eventquery.GetEventRequest) (*eventquery.GetEventResponse, error) {
	return nil, errGrpcNotImplemented
}

func (handler *EventQueryService) GetEvents(_ context.Context, _ *eventquery.GetEventsRequest) (*eventquery.GetEventsResponse, error) {
	return nil, errGrpcNotImplemented
}

func (handler *EventQueryService) Count(ctx context.Context, req *eventquery.CountRequest) (*eventquery.CountResponse, error) {
	ctx, span := trace.StartSpan(ctx, "zenoss.cloud/eventQuery/v2/go/EventQueryService.Count")
	defer span.End()
	log := zenkit.ContextLogger(ctx).WithField("req", req)
	ident, err := auth.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	findReq, err := QueryProtoToEventQuery(ident.TenantName(), req.Query)
	if err != nil {
		log.WithError(err).Error("failed to convert request to domain query")
		return nil, err
	}
	reqFields := make([]string, len(req.Fields))
	for i, f := range req.Fields {
		reqFields[i] = f.Field
	}
	log.
		WithField("query", findReq).
		Tracef("executing Find request")
	countResp, err := handler.svr.Count(ctx, &event.CountRequest{
		Query:          *findReq,
		Fields:         reqFields, // TODO: fix type of this field in zing-proto
		CountInstances: req.CountInstances,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to execute event context query: %q", err)
	}
	response := &eventquery.CountResponse{
		Results: make(map[string]*eventquery.CountResponse_Results),
	}
	for field, countResults := range countResp.Results {
		countResultsProto := make([]*eventquery.CountResult, len(countResults))
		for i, r := range countResults {
			sV, err := structpb.NewValue(r.Value)
			if err != nil {
				log.WithError(err).Error("failed to convert result to struct value")
				return nil, err
			}
			countResultsProto[i] = &eventquery.CountResult{
				Value: sV,
				Count: r.Count,
			}
		}
		response.Results[field] = &eventquery.CountResponse_Results{Values: countResultsProto}

	}
	return response, nil
}

func (handler *EventQueryService) Frequency(ctx context.Context, req *eventquery.FrequencyRequest) (*eventquery.FrequencyResponse, error) {
	ctx, span := trace.StartSpan(ctx, "zenoss.cloud/eventQuery/v2/go/EventQueryService.Frequency")
	defer span.End()
	log := zenkit.ContextLogger(ctx).WithField("req", req)
	ident, err := auth.CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	findReq, err := QueryProtoToEventQuery(ident.TenantName(), req.Query)
	if err != nil {
		log.WithError(err).Error("failed to convert request to domain query")
		return nil, err
	}
	log.
		WithField("query", findReq).
		Tracef("executing Find request")
	fields := make([]string, 0)
	for _, fieldPb := range req.Fields {
		fields = append(fields, fieldPb.Field)
	}
	groupBys := make([]string, 0)
	for _, groupByPb := range req.GroupBy {
		groupBys = append(groupBys, groupByPb.Field)
	}
	resp1, err := handler.svr.Frequency(ctx, &event.FrequencyRequest{
		Query:          *findReq,
		Fields:         fields,
		GroupBy:        groupBys,
		Downsample:     req.Downsample,
		PersistCounts:  req.PersistCounts,
		CountInstances: req.CountInstances,
	})
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to execute frequency request: %q", err)
	}
	responseProto := &eventquery.FrequencyResponse{
		Timestamps: resp1.Timestamps,
	}
	freqResults := make(map[string]*eventquery.FrequencyResponse_Results, 0)
	for _, r := range resp1.Results {
		keys := make([]string, 0)
		for k := range r.Key {
			keys = append(keys, k)
			if len(keys) > 0 {
				break
			}
		}
		sv, err := structpb.NewValue(r.Key[keys[0]])
		if err != nil {
			return nil, err
		}
		if freqResultEntry, ok := freqResults[keys[0]]; !ok {
			freqResults[keys[0]] = &eventquery.FrequencyResponse_Results{
				Values: []*eventquery.FrequencyResult{
					{
						Value:  sv,
						Counts: r.Values,
					},
				},
			}
		} else {
			freqResultEntry.Values = append(freqResultEntry.Values, &eventquery.FrequencyResult{
				Value:  sv,
				Counts: r.Values,
			})
		}
	}
	responseProto.Results = freqResults
	return responseProto, nil
}
