package service

import (
	"context"
	"time"

	"github.com/zenoss/event-context-svc/utils"
	"github.com/zenoss/event-management-service/metrics"
	"github.com/zenoss/zenkit/v5"
	ecproto "github.com/zenoss/zing-proto/v11/go/cloud/event_context"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
	"go.opencensus.io/stats"
	"go.opencensus.io/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EventManagementService implements the interface proto.EventManagementServer
type EventManagementService struct {
	eventCtxClient ecproto.EventContextIngestClient
}

// NewEventManagementService returns implementation of proto.EventManagementServer
func NewEventManagementService(ctx context.Context) (proto.EventManagementServer, error) {
	log := zenkit.ContextLogger(ctx)
	svc := &EventManagementService{}
	if svc.eventCtxClient == nil {
		ecConn, err := zenkit.NewClientConnWithRetry(ctx, "event-context-ingest", zenkit.DefaultRetryOpts())
		if err != nil {
			log.WithError(err).Error("failed to connect to event-context-ingest-svc")
			return nil, err
		}
		svc.eventCtxClient = ecproto.NewEventContextIngestClient(ecConn)

		log.Info("connected to event-context-ingest-svc")
	}
	return svc, nil
}

// NewEventManagementServiceFromParts returns implementation of proto.EventManagementServer with the passed in client for testing
func NewEventManagementServiceFromParts(client ecproto.EventContextIngestClient) (proto.EventManagementServer, error) {
	return &EventManagementService{
		eventCtxClient: client,
	}, nil
}

func sinceInMilliseconds(startTime time.Time) float64 {
	return float64(time.Since(startTime).Nanoseconds()) / 1e6
}

func addStatusResponse(response *proto.EventStatusResponse, item *proto.EMEventStatus, success bool, error string) {
	resp := proto.EMEventStatusResponse{
		EventId:      item.EventId,
		OccurrenceId: item.OccurrenceId,
		Success:      success,
		Error:        error,
	}
	response.StatusResponses = append(response.StatusResponses, &resp)
}

// SetStatus sets the staus of the event(s) passed in
func (svc *EventManagementService) SetStatus(ctx context.Context, request *proto.EventStatusRequest) (*proto.EventStatusResponse, error) {
	ctx, span := trace.StartSpan(ctx, "EventMangement.SetStatus")
	log := zenkit.ContextLogger(ctx)
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid set status nil request")
	}

	// setup metrics
	mTime := time.Now()
	defer func() {
		span.End()
		stats.Record(ctx, metrics.MSetStatusTimeMs.M(sinceInMilliseconds(mTime)),
			metrics.MSetStatusCount.M(int64(len(request.Statuses))))
	}()

	// validate
	_, err := utils.ValidateIdentity(ctx)
	if err != nil {
		log.WithError(err).Error("SetStatus failed: unauthenticated")
		return nil, err
	}

	// iterate thru the statuses sent in
	response := new(proto.EventStatusResponse)
	for _, item := range request.Statuses {
		if item.EventId == "" {
			addStatusResponse(response, item, false, "Event id cannot be empty")
		} else if item.OccurrenceId == "" {
			addStatusResponse(response, item, false, "Occurrence id cannot be empty")
		} else if item.Acknowledged == nil && item.StatusWrapper == nil {
			addStatusResponse(response, item, false, "Need status or acknowledged to be set")
		} else {
			// process
			ecRequest := ecproto.UpdateEventRequest{
				OccurrenceId: item.OccurrenceId,
				Acknowledged: item.Acknowledged,
				EventId:      item.EventId,
			}

			if item.StatusWrapper != nil {
				sw := ecproto.UpdateEventRequest_Wrapper{
					Status: item.StatusWrapper.Status,
				}
				ecRequest.StatusWrapper = &sw
			}

			_, err := svc.eventCtxClient.UpdateEvent(ctx, &ecRequest) // ignore response as we dont expect note id
			if err != nil {
				log.Error("Failed setting status", err)
				addStatusResponse(response, item, false, err.Error())
			} else {
				addStatusResponse(response, item, true, "")
			}
		}
	}
	return response, nil
}

func addAnotationResponse(response *proto.EventAnnotationResponse, item *proto.Annotation, success bool, aid string, error string) {
	resp := proto.AnnotationResponse{
		EventId:      item.EventId,
		OccurrenceId: item.OccurrenceId,
		AnnotationId: aid,
		Success:      success,
		Error:        error,
	}
	response.AnnotationResponses = append(response.AnnotationResponses, &resp)
}

// Annotate adds a annotation to the associated event
func (svc *EventManagementService) Annotate(ctx context.Context, request *proto.EventAnnotationRequest) (*proto.EventAnnotationResponse, error) {
	ctx, span := trace.StartSpan(ctx, "EventMangement.Annotate")
	log := zenkit.ContextLogger(ctx)
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid annotate nil request")
	}

	// setup metrics
	mTime := time.Now()
	defer func() {
		span.End()
		stats.Record(ctx, metrics.MAnnotateTimeMs.M(sinceInMilliseconds(mTime)),
			metrics.MAnnotateCount.M(int64(len(request.Annotations))))
	}()

	// validate
	_, err := utils.ValidateIdentity(ctx)
	if err != nil {
		log.WithError(err).Error("Annotate failed: unauthenticated")
		return nil, err
	}

	// iterate thru the annotations sent in
	response := new(proto.EventAnnotationResponse)
	for _, item := range request.Annotations {
		if item.EventId == "" || item.OccurrenceId == "" {
			addAnotationResponse(response, item, false, "", "Event id, Occurrence id cannot be empty")
		} else if item.Annotation == "" {
			addAnotationResponse(response, item, false, "", "Annotation cannot be empty")
		} else {
			ecRequest := ecproto.UpdateEventRequest{
				OccurrenceId: item.OccurrenceId,
				NoteId:       item.AnnotationId,
				Note:         item.Annotation,
				EventId:      item.EventId,
			}

			resp, err := svc.eventCtxClient.UpdateEvent(ctx, &ecRequest)
			if err == nil {
				addAnotationResponse(response, item, true, resp.NoteId, "")
			} else {
				log.Error("Failed annotating", err)
				addAnotationResponse(response, item, false, "", err.Error())
			}
		}
	}
	return response, nil
}

// asserts
var _ proto.EventManagementServer = &EventManagementService{}
