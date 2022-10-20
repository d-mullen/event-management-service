package scopes

import (
	"context"

	"github.com/zenoss/event-management-service/pkg/models/event"
)

type (
	EntityScopeProvider interface {
		GetEntityIDs(ctx context.Context, cursor string, timeRange event.TimeRange) ([]string, error)
	}
	ActiveEntityInput struct {
		Timestamp int64
		TenantID  string
		EntityID  string
		EventID   string
	}
	PutActiveEntityRequest struct {
		Inputs []*ActiveEntityInput
	}
	ActiveEntityRepository interface {
		// Put takes an array of inputs, each representing an instance in time for which an event was recorded
		// for an entity, and adds them to appropriate set of entities with active event occurrences.
		Put(ctx context.Context, request *PutActiveEntityRequest) error
		// Get takes an array of entities, a duration, and tenant ID and returns an array of IDs of entities known to have
		// events recorded in the given time range. If no active entity sets are stored for the requested time, then
		// original entityIDs are returned to the caller.
		Get(ctx context.Context, timeRange event.TimeRange, tenantID string, entityIDs []string) ([]string, error)
	}
)
