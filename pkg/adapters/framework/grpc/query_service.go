package grpc

import (
	"context"
	"errors"

	appEvent "github.com/zenoss/event-management-service/pkg/application/event"
	"github.com/zenoss/event-management-service/pkg/domain/event"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	"github.com/zenoss/zing-proto/v11/go/cloud/eventquery"
	eventPb "github.com/zenoss/zing-proto/v11/go/event"
	"github.com/zenoss/zingo/v4/protobufutils"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	errGrpcNotImplemented = status.Error(codes.Unimplemented, "unimplemented")
)

type (
	EventQueryService struct {
		svr appEvent.Service
	}
)

func NewEventQueryService(eventQuerySvr appEvent.Service) *EventQueryService {
	return &EventQueryService{
		svr: eventQuerySvr,
	}
}

var _ eventquery.EventQueryServer = &EventQueryService{}

func CheckAuth(ctx context.Context) (zenkit.TenantIdentity, error) {
	ident := zenkit.ContextTenantIdentity(ctx)
	if ident == nil {
		return nil, status.Error(codes.Unauthenticated, "no identity found on context")
	}
	if len(ident.TenantName()) == 0 && len(ident.Tenant()) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "not tenant claim found identity: %#v", ident)
	}
	return ident, nil
}

func sortOrderFromProto(order common.SortOrder) event.SortOrder {
	if order == common.SortOrder_DESC {
		return event.SortOrderDescending
	}
	return event.SortOrderAscending
}
func fieldNameFromProto(byField *common.SortBy_ByField) string {
	switch v := byField.ByField.SortField.(type) {
	case *common.SortByField_Property:
		return v.Property
	case *common.SortByField_Dimension:
		return v.Dimension
	case *common.SortByField_Metadata:
		return v.Metadata
	default:
		return ""
	}
}

func ProtoSortByToDomain(orig *common.SortBy) (*event.SortOpt, error) {
	switch v := orig.SortType.(type) {
	case *common.SortBy_ByField:
		return &event.SortOpt{
			Field:     fieldNameFromProto(v),
			SortOrder: sortOrderFromProto(v.ByField.GetOrder()),
		}, nil
	case *common.SortBy_BySeries:
		return &event.SortOpt{
			Field:     v.BySeries.SeriesName,
			SortOrder: sortOrderFromProto(v.BySeries.Order),
		}, nil
	default:
		return nil, status.Error(codes.InvalidArgument, "unsupported sort by type")
	}
}

func QueryProtoToEventQuery(tenantID string, query *eventquery.Query, pageInput *eventquery.PageInput) (*event.Query, error) {
	if query == nil {
		return nil, status.Error(codes.InvalidArgument, "nil query found on request")
	}
	if tr := query.TimeRange; tr == nil || tr.End < tr.Start {
		return nil, status.Error(codes.InvalidArgument, "invalid timerange")
	}
	result := &event.Query{
		Tenant: tenantID,
		TimeRange: event.TimeRange{
			Start: query.TimeRange.Start,
			End:   query.TimeRange.End,
		},
	}
	severities := make([]event.Severity, len(query.Severities))
	for i, sev := range query.Severities {
		severities[i] = event.Severity(sev)
	}
	result.Severities = severities
	statuses := make([]event.Status, len(query.Statuses))
	for i, status := range query.Statuses {
		statuses[i] = event.Status(status)
	}
	result.Statuses = statuses
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
	if pageInput != nil {
		if result.PageInput == nil {
			result.PageInput = &event.PageInput{}
		}
		result.PageInput.Cursor = pageInput.Cursor
		result.PageInput.Limit = uint64(pageInput.Limit)
	}
	return result, nil
}

func (handler *EventQueryService) Search(ctx context.Context, req *eventquery.SearchRequest) (*eventquery.SearchResponse, error) {
	ctx, span := trace.StartSpan(ctx, "zenoss.cloud/eventQuery/v2/go/EventQueryService.Search")
	defer span.End()
	log := zenkit.ContextLogger(ctx).WithField("req", req)
	ident, err := CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	findReq, err := QueryProtoToEventQuery(ident.TenantName(), req.Query, req.PageInput)
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
	resp := &eventquery.SearchResponse{
		Results: make([]*eventquery.EventResult, 0),
		PageInfo: &common.PageInfo{
			EndCursor: page.Cursor,
			HasNext:   page.HasNext,
		},
	}
	for _, result := range page.Results {
		currID := result.ID
		currEntity := result.Entity
		currDim := make(map[string]*common.Scalar)
		currName := result.Name
		for k, v := range result.Dimensions {
			scalar, err := protobufutils.ToScalar(v)
			if err != nil {
				return nil, status.Error(codes.Internal, "failed to marshal dimension")
			}
			currDim[k] = scalar
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
					TimeRange:     &common.TimeRange{Start: occ.StartTime, End: occ.EndTime},
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
					currOcc.Metadata = make(map[string]*common.ScalarArray)
					for k, anyValues := range occ.Metadata {
						scalarArray, err := protobufutils.ToScalarArray(anyValues)
						if err != nil {
							return nil, status.Errorf(codes.Internal, "failed to convert metatadata to protobuf: %q", err)
						}
						currOcc.Metadata[k] = scalarArray
					}
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
	ident, err := CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	findReq, err := QueryProtoToEventQuery(ident.TenantName(), req.Query, nil)
	if err != nil {
		log.WithError(err).Error("failed to convert request to domain query")
		return nil, err
	}
	log.
		WithField("query", findReq).
		Tracef("executing Find request")
	countResp, err := handler.svr.Count(ctx, &event.CountRequest{
		Query:  *findReq,
		Fields: []string{req.Fields.Field}, // TODO: fix type of this field in zing-proto
	})
	if err != nil {
		return nil, status.Errorf(codes.Unknown, "failed to execute event context query: %q", err)
	}
	response := &eventquery.CountResponse{
		Field:  req.Fields.Field,
		Counts: make([]*eventquery.CountResult, 0),
	}
	counts := response.Counts
	if countResults, ok := countResp.Results[response.Field]; ok {
		for _, countResult := range countResults {
			counts = append(counts, &eventquery.CountResult{
				Value: protobufutils.MustToScalar(countResult.Value),
				Count: countResult.Count,
			})
		}
		response.Counts = counts
	}
	return response, nil
}

func (handler *EventQueryService) Frequency(ctx context.Context, req *eventquery.FrequencyRequest) (*eventquery.FrequencyResponse, error) {
	ctx, span := trace.StartSpan(ctx, "zenoss.cloud/eventQuery/v2/go/EventQueryService.Frequency")
	defer span.End()
	log := zenkit.ContextLogger(ctx).WithField("req", req)
	ident, err := CheckAuth(ctx)
	if err != nil {
		return nil, err
	}
	findReq, err := QueryProtoToEventQuery(ident.TenantName(), req.Query, nil)
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
	resp := &eventquery.FrequencyResponse{
		Timestamps: resp1.Timestamps,
	}
	freqResults := make([]*eventquery.FrequencyResult, 0)
	for _, r := range resp1.Results {
		key, err := protobufutils.ToScalarMap(r.Key)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to convert frequency result key to scalar map")
		}
		freqResults = append(freqResults, &eventquery.FrequencyResult{
			Key:    key,
			Values: r.Values,
		})
	}
	resp.Results = freqResults
	return resp, nil
}

func (handler *EventQueryService) SearchStream(_ *eventquery.SearchStreamRequest, _ eventquery.EventQuery_SearchStreamServer) error {
	return errGrpcNotImplemented
}

func (handler *EventQueryService) EventsWithCountsStream(_ *eventquery.EventsWithCountsRequest, _ eventquery.EventQuery_EventsWithCountsStreamServer) error {
	return errGrpcNotImplemented
}
