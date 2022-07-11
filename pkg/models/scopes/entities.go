package scopes

import "context"

type (
	EntityScopeProvider interface {
		GetEntityIDs(ctx context.Context, cursor string) ([]string, error)
	}
)
