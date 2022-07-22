package event

import (
	"context"
)

type (
	Cursor struct {
		Query  `json:"query,omitempty"`
		ID     string         `json:"id,omitempty"`
		Config map[string]any `json:"config,omitempty"`
	}

	CursorRepository interface {
		New(context.Context, *Cursor) (string, error)
		Get(context.Context, string) (*Cursor, error)
	}
)
