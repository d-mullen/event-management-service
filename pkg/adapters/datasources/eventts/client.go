package eventts

import (
	"context"
	"io"
	"math"
	"time"

	legacyProto "github.com/golang/protobuf/proto"
	"github.com/zenoss/event-management-service/pkg/models/eventts"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	eventtsProto "github.com/zenoss/zing-proto/v11/go/cloud/eventts"
	"github.com/zenoss/zingo/v4/protobufutils"
	"golang.org/x/exp/slices"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type (
	eventTSRepo struct {
		client eventtsProto.EventTSServiceClient
	}
	GetRequest eventts.GetRequest
)

func NewAdapter(client eventtsProto.EventTSServiceClient) *eventTSRepo {
	return &eventTSRepo{
		client: client,
	}
}

var _ eventts.Repository = &eventTSRepo{}

func (repo *eventTSRepo) Get(ctx context.Context, req *eventts.GetRequest) ([]*eventts.Occurrence, error) {
	results := make([]*eventts.Occurrence, 0)
	err := repo.getStream(ctx, req, func(r *eventts.Occurrence) bool {
		results = append(results, r)
		return true
	})
	if err != nil {
		return nil, err
	}
	return results, nil
}

func v1MessageToV2(msg legacyProto.Message) proto.Message {
	return legacyProto.MessageV2(msg)
}

func mustMarshal(msg proto.Message) string {
	b, _ := protojson.Marshal(msg)
	return string(b)
}

func (repo *eventTSRepo) getStream(ctx context.Context, req *eventts.GetRequest, callback func(*eventts.Occurrence) bool) error {
	log := zenkit.ContextLogger(ctx)
	reqProto, err := EventTSRequestToProto(req)
	if err != nil {
		return err
	}
	ewcReq := &eventtsProto.EventsWithCountsRequest{
		EventIds:      reqProto.EventIds,
		TimeRange:     reqProto.TimeRange,
		OccurrenceMap: reqProto.OccurrenceMap,
		EventParams: &eventtsProto.EventsWithCountsRequest_EventsParams{
			ResultFields: reqProto.ResultFields,
		},
		Limit:   1,
		Filters: reqProto.Filters,
	}
	log.WithField("get_eventts_proto", mustMarshal(reqProto)).Trace("making event-ts request")
	// stream, err := repo.client.GetEventsStream(ctx, reqProto)
	stream, err := repo.client.EventsWithCountsStream(ctx, ewcReq)
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			// no-op
		}
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.WithField("getEventsStreamResponse", resp).Trace("got response from event-ts-svc")
		for _, r := range resp.GetResults() {
			for _, series := range r.GetSeries() {
				results, err := EventTSSeriesToOccurrence(series)
				if err != nil {
					return err
				}
				for _, r := range results {
					callback(r)
				}
			}
			for _, series := range r.GetPartialSeries() {
				results, err := EventTSSeriesToOccurrence(series)
				if err != nil {
					return err
				}
				for _, r := range results {
					callback(r)
				}
			}
		}
	}
	return nil
}

func (repo *eventTSRepo) GetStream(ctx context.Context, req *eventts.GetRequest) <-chan *eventts.OccurrenceOptional {
	out := make(chan *eventts.OccurrenceOptional)
	go func(gCtx context.Context) {
		defer close(out)
		err := repo.getStream(gCtx, req, func(o *eventts.Occurrence) bool {
			out <- &eventts.OccurrenceOptional{
				Result: o,
			}
			return true
		})
		if err != nil {
			out <- &eventts.OccurrenceOptional{
				Err: err,
			}
			return
		}
	}(ctx)
	return out
}

func (repo *eventTSRepo) Frequency(_ context.Context, _ *eventts.FrequencyRequest) (*eventts.FrequencyResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (repo *eventTSRepo) FrequencyStream(_ context.Context, _ *eventts.FrequencyRequest) chan *eventts.FrequencyOptional {
	panic("not implemented") // TODO: Implement
}

func EventTSRequestToProto(req *eventts.GetRequest) (*eventtsProto.EventTSRequest, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "nil request")
	}
	output := &eventtsProto.EventTSRequest{
		// Count:        0,   // TODO
	}

	// temporary workaround until clients are updated to pass explicit fields
	_, usedCatchall := slices.BinarySearch(req.ResultFields, "metadata")
	if len(req.ResultFields) == 0 || usedCatchall {
		output.ResultFields = []string{"_zv_status",
			"_zv_severity",
			"_zv_summary",
			"contextTitle",
			"parentContextTitle",
			"_zv_name",
			"lastSeen",
			"eventClass",
			"source",
			"CZ_EVENT_DETAIL-zenoss.device.production_state",
			"CZ_EVENT_DETAIL-zenoss.device.location",
			"CZ_EVENT_DETAIL-zenoss.device.groups",
			"CZ_EVENT_DETAIL-zenoss.device.systems",
			"CZ_EVENT_DETAIL-zenoss.device.priority",
			"CZ_EVENT_DETAIL-zenoss.device.device_class",
			"CZ_EVENT_DETAIL-zenoss.device.IncidentManagement.number",
			"_zen_entityIds",
			"_zen_parentEntityIds",
			"source-type"}
	}
	if len(req.Filters) > 0 {
		output.Filters = eventTsFilter2eventProtoFilter(req.Filters)
	}

	if len(req.ByOccurrences.OccurrenceMap) > 0 {
		eventIDs := make([]string, 0)
		var minStart, maxEnd int64 = math.MaxInt64, math.MinInt64
		occurrenceMap := make(map[string]*eventtsProto.EventTSOccurrenceCollection)
		for eventID, occurrenceInputs := range req.ByOccurrences.OccurrenceMap {
			_, ok := occurrenceMap[eventID]
			if !ok {
				occurrenceMap[eventID] = &eventtsProto.EventTSOccurrenceCollection{}
				eventIDs = append(eventIDs, eventID)
			}
			currOccSlice := make([]*eventtsProto.EventTSOccurrence, 0)
			for _, occurrenceInput := range occurrenceInputs {
				if occurrenceInput.TimeRange.Start < minStart {
					minStart = occurrenceInput.TimeRange.Start
				}
				if occurrenceInput.TimeRange.End > maxEnd {
					maxEnd = occurrenceInput.TimeRange.End
				}
				currOccSlice = append(currOccSlice, &eventtsProto.EventTSOccurrence{
					Id: occurrenceInput.ID,
					TimeRange: &common.TimeRange{
						Start: occurrenceInput.TimeRange.Start,
						End:   occurrenceInput.TimeRange.End,
					},
					IsActive: occurrenceInput.IsActive,
				})
			}
			occurrenceMap[eventID].Occurrences = currOccSlice
		}
		if req.ByOccurrences.ShouldApplyIntervals {
			if maxEnd <= 0 {
				maxEnd = time.Now().UnixMilli() // TODO: HACK!!!
			}
			output.TimeRange = &common.TimeRange{
				Start: minStart,
				End:   maxEnd,
			}
		} else {
			output.TimeRange = &common.TimeRange{
				Start: req.TimeRange.Start,
				End:   req.TimeRange.End,
			}
		}
		output.EventIds = eventIDs
		output.OccurrenceMap = occurrenceMap
		return output, nil
	}
	if len(req.ByEventIDs.IDs) > 0 {
		output.EventIds = req.ByEventIDs.IDs
		output.TimeRange = &common.TimeRange{
			Start: req.TimeRange.Start,
			End:   req.TimeRange.End,
		}
		return output, nil
	}
	return nil, status.Error(codes.InvalidArgument, "must specify get request by either non-empty IDs or occurrences")
}

func EventTSSeriesToOccurrence(msg *eventtsProto.EventTSSeries) ([]*eventts.Occurrence, error) {
	if msg == nil {
		return nil, status.Error(codes.Aborted, "got nil results")
	}
	results := make([]*eventts.Occurrence, 0)
	for _, eventTSFields := range msg.Values {
		md, err := protobufutils.FromScalarArrayMap(eventTSFields.GetData())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to decode scalar array map: %q", err)
		}
		results = append(results, &eventts.Occurrence{
			ID:        msg.OccurrenceId,
			EventID:   msg.EventId,
			Timestamp: eventTSFields.Timestamp,
			Metadata:  md,
		})
	}
	return results, nil
}

// main reason to dial with all this transformations - possibility to completely
// remove ts-svc later.
func eventTsFilter2eventProtoFilter(in []*eventts.Filter) []*eventtsProto.EventTSFilter {
	result := make([]*eventtsProto.EventTSFilter, 0)
	for _, filter := range in {
		item := eventtsProto.EventTSFilter{}
		item.FieldName = filter.Field
		item.Op = common.Operation(filter.Operation)
		if len(filter.Values) > 0 {
			item.Values = protobufutils.MustToScalarArray(filter.Values) // force panic on err
			result = append(result, &item)
		}
	}
	return result
}
