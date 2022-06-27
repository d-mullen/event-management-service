package event

import (
	"context"

	"github.com/zenoss/event-management-service/pkg/domain/event"
	"github.com/zenoss/event-management-service/pkg/domain/scopes"
)

func CursorToEntityFilter(ctx context.Context, cursor string, provider scopes.EntityScopeProvider) (*event.Filter, error) {
	panic("not implemented")
}
