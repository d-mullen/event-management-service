package service

import (
	"context"

	//"github.com/pkg/errors"
	"github.com/zenoss/event-context-svc/utils"
	"github.com/zenoss/zenkit/v5"
	ecproto "github.com/zenoss/zing-proto/v11/go/cloud/event_context"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EventManagementService implements the interface proto.EventManagementServer
type EventManagementService struct {
	eventCtxClient ecproto.EventContextQueryClient
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
		svc.eventCtxClient = ecproto.NewEventContextQueryClient(ecConn)

		log.Info("connected to event-context-svc")
	}
	return svc, nil
}

// SetStatus sets the staus of the event(s) passed in
func (svc *EventManagementService) SetStatus(ctx context.Context, request *proto.EventStatusRequest) (*proto.EventStatusResponse, error) {

	log := zenkit.ContextLogger(ctx)
	if request == nil || request.Tenant == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid set status request (need tenant")
	}

	_, err := utils.ValidateIdentity(ctx)
	if err != nil {
		log.WithError(err).Error("SetStatus failed: unauthenticated")
		return nil, err
	}

	response := new(proto.EventStatusResponse)

	return response, nil
}

// Annotate adds a annotation to the associated event
func (svc *EventManagementService) Annotate(ctx context.Context, request *proto.EventAnnotationRequest) (*proto.EventAnnotationResponse, error) {
	response := new(proto.EventAnnotationResponse)

	return response, nil
}

// asserts
var _ proto.EventManagementServer = &EventManagementService{}
