package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zenoss/event-context-svc/utils"
	"github.com/zenoss/zenkit/v5"
	ecproto "github.com/zenoss/zing-proto/v11/go/cloud/event_context"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
	eproto "github.com/zenoss/zing-proto/v11/go/event"
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
			log.WithError(err).Error("failed to connect to event-context-svc")
			return nil, err
		}
		svc.eventCtxClient = ecproto.NewEventContextIngestClient(ecConn)

		log.Info("connected to event-context-svc")
	}
	return svc, nil
}

// NewEventManagementServiceFromParts returns implementation of proto.EventManagementServer with the passed in client for testing
func NewEventManagementServiceFromParts(client ecproto.EventContextIngestClient) (proto.EventManagementServer, error) {
	return &EventManagementService{
		eventCtxClient: client,
	}, nil
}

func getECStatus(status proto.EMStatus) eproto.Status {
	switch status {
	case proto.EMStatus_EM_STATUS_DEFAULT:
		return eproto.Status_STATUS_DEFAULT
	case proto.EMStatus_EM_STATUS_OPEN:
		return eproto.Status_STATUS_OPEN
	case proto.EMStatus_EM_STATUS_SUPPRESSED:
		return eproto.Status_STATUS_SUPPRESSED
	case proto.EMStatus_EM_STATUS_CLOSED:
		return eproto.Status_STATUS_CLOSED
	default:
		// Should never happen ?
		return eproto.Status_STATUS_DEFAULT
	}
}

// SetStatus sets the staus of the event(s) passed in
func (svc *EventManagementService) SetStatus(ctx context.Context, request *proto.EventStatusRequest) (*proto.EventStatusResponse, error) {

	log := zenkit.ContextLogger(ctx)
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid set status nil request")
	}

	_, err := utils.ValidateIdentity(ctx)
	if err != nil {
		log.WithError(err).Error("SetStatus failed: unauthenticated")
		return nil, err
	}

	response := new(proto.EventStatusResponse)
	response.SuccessList = make(map[string]bool)

	for k, v := range request.StatusList {
		if v.EventId == "" {
			return nil, errors.Wrap(err, "Event id cannot be empty")
		}
		ecRequest := ecproto.UpdateEventRequest{
			OccurrenceId: k,
			Acknowledged: v.Acknowledge,
			EventId:      v.EventId,
		}

		if v.StatusWrapper != nil {
			sw := ecproto.UpdateEventRequest_Wrapper{
				Status: getECStatus(v.StatusWrapper.Status),
			}
			ecRequest.StatusWrapper = &sw
		}

		resp, err := svc.eventCtxClient.UpdateEvent(ctx, &ecRequest)
		if err != nil {
			log.Error("Failed setting status", err)
			response.SuccessList[k] = false
		} else {
			response.SuccessList[k] = resp.Status
		}
	}
	return response, nil
}

// Annotate adds a annotation to the associated event
func (svc *EventManagementService) Annotate(ctx context.Context, request *proto.EventAnnotationRequest) (*proto.EventAnnotationResponse, error) {
	log := zenkit.ContextLogger(ctx)
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid annotate nil request")
	}

	_, err := utils.ValidateIdentity(ctx)
	if err != nil {
		log.WithError(err).Error("Annotate failed: unauthenticated")
		return nil, err
	}

	response := new(proto.EventAnnotationResponse)
	response.AnnotationResponseList = make(map[string]*proto.AnnotationResponse)

	for k, v := range request.AnnotationList {
		if v.Annotation == "" || v.EventId == "" {
			return nil, errors.Wrap(err, "Event id, Annotation cannot be empty")
		}
		ecRequest := ecproto.UpdateEventRequest{
			OccurrenceId: k,
			NoteId:       v.AnnotationId,
			Note:         v.Annotation,
			EventId:      v.EventId,
		}

		resp, err := svc.eventCtxClient.UpdateEvent(ctx, &ecRequest)
		aresp := proto.AnnotationResponse{}
		if err == nil {
			aresp.Success = resp.Status
			aresp.AnnotationId = resp.NoteId
		} else {
			aresp.Success = false
			log.Error("Failed annotating", err)
		}
		response.AnnotationResponseList[k] = &aresp
	}

	return response, nil
}

// asserts
var _ proto.EventManagementServer = &EventManagementService{}
