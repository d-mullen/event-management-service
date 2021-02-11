package service

import (
	"context"

	"github.com/zenoss/event-context-svc/utils"

	"github.com/pkg/errors"
	"github.com/zenoss/event-context-svc/store"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// EventManagementService implements the interface proto.EventManagementServer
type EventManagementService struct {
	qryEventStore store.EventContextStore
}

// NewEventService returns implementation of proto.EventContextContextServer
func NewEventService(ctx context.Context) (proto.EventManagementServer, error) {
	svc := &EventManagementService{}
	if svc.qryEventStore == nil {
		eventStore, err := store.DefaultEventsStore(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get event context store for event management service")
		}
		svc.qryEventStore = eventStore
	}
	return svc, nil
}

// SetStatus sets the staus of the event(s) passed in
func (svc *EventManagementService) SetStatus(ctx context.Context, request *proto.EventStatusRequest) (*proto.EventManagementResponse, error) {
	var (
		results *store.ActiveEvents
		pi      *store.PageInput
	)
	log := zenkit.ContextLogger(ctx)
	if request == nil || request.Tenant == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid set status request (need tenant")
	}

	tenant, err := utils.ValidateIdentity(ctx)
	if err != nil {
		log.WithError(err).Error("SetStatus failed: unauthenticated")
		return nil, err
	}

	response := new(proto.EventManagementResponse)

	return response, nil
}

// SetAcknowledge sets acknowledge boolean value for events
func (svc *EventManagementService) SetAcknowledge(ctx context.Context, request *proto.EventAcknowledgeRequest) (*proto.EventManagementResponse, error) {
	response := new(proto.EventManagementResponse)

	return response, nil
}

// Annotate adds a annotation to the associated event
func (svc *EventManagementService) Annotate(ctx context.Context, request *proto.EventAnnotationRequest) (*proto.EventAnnotationResponse, error) {
	response := new(proto.EventAnnotationResponse)

	return response, nil
}

// GetAllowedStates returns a list of allowed stated for the tenant
func (svc *EventManagementService) GetAllowedStates(ctx context.Context, request *proto.EventAllowedStatesRequest) (*proto.EventAllowedStatesResponse, error) {
	response := new(proto.EventAllowedStatesResponse)

	return response, nil
}

// asserts
var _ proto.EventManagementServer = &EventManagementService{}
