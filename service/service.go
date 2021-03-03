package service

import (
	"context"

	//"github.com/pkg/errors"
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
		ecConn, err := zenkit.NewClientConnWithRetry(ctx, "event-context-svc", zenkit.DefaultRetryOpts())
		if err != nil {
			log.WithError(err).Error("failed to connect to event-context-svc")
			return nil, err
		}
		svc.eventCtxClient = ecproto.NewEventContextIngestClient(ecConn)

		log.Info("connected to event-context-svc")
	}
	return svc, nil
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
	if request == nil || request.Tenant == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid set status request (need tenant)")
	}

	tenant, err := utils.ValidateIdentity(ctx)
	if err != nil {
		log.WithError(err).Error("SetStatus failed: unauthenticated")
		return nil, err
	}

	response := new(proto.EventStatusResponse)
	response.SuccessList = make(map[string]bool)

	for k, v := range request.StatusList {
		ecRequest := ecproto.UpdateEventRequest{
			Tenant:       tenant,
			Id:           k,
			Status:       getECStatus(v.Status),
			Acknowledged: v.Acknowledge,
		}

		resp, err := svc.eventCtxClient.UpdateEvent(ctx, &ecRequest)
		response.SuccessList[k] = resp.Status
		if err != nil {
			log.Error("Failed setting status", err)
		}
	}
	return response, nil
}

// Annotate adds a annotation to the associated event
func (svc *EventManagementService) Annotate(ctx context.Context, request *proto.EventAnnotationRequest) (*proto.EventAnnotationResponse, error) {
	log := zenkit.ContextLogger(ctx)
	if request == nil || request.Tenant == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid annotate request (need tenant)")
	}

	_, err := utils.ValidateIdentity(ctx)
	if err != nil {
		log.WithError(err).Error("Annotate failed: unauthenticated")
		return nil, err
	}

	response := new(proto.EventAnnotationResponse)
	response.AnnotationList = make(map[string]*proto.AnnotationResponse)

	for k, v := range request.AnnotationList {
		ecRequest := ecproto.UpdateEventRequest{
			Id:     k,
			NoteId: v.Id,
			Note:   v.Annotation,
		}

		resp, err := svc.eventCtxClient.UpdateEvent(ctx, &ecRequest)
		aresp := proto.AnnotationResponse{}
		aresp.Success = resp.Status
		if err == nil {
			aresp.NoteId = resp.NoteId
		} else {
			log.Error("Failed annotating", err)
		}
		response.AnnotationList[k] = &aresp
	}

	return response, nil
}

// asserts
var _ proto.EventManagementServer = &EventManagementService{}
