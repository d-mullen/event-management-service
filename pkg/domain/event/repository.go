package event

import (
	"context"

	"github.com/pkg/errors"
)

type FilterOp string

// Filter operations
const (
	FilterOpLessThan             FilterOp = "<"
	FilterOpLessThanOrEqualTo    FilterOp = "<="
	FilterOpGreaterThan          FilterOp = ">"
	FilterOpGreaterThanOrEqualTo FilterOp = ">="
	FilterOpEqualTo              FilterOp = "=="
	FilterOpNotEqualTo           FilterOp = "!="
	FilterOpIn                   FilterOp = "in"
	FilterOpNotIn                FilterOp = "not-in"
	FilterOpOr                   FilterOp = "or"
	FilterOpAnd                  FilterOp = "and"
	FilterOpNot                  FilterOp = "not"
)

type SortOrder int

const (
	SortOrderDescending = -1
	SortOrderAscending  = 1
)

type (
	TimeRange struct {
		Start int64
		End   int64
	}
	Query struct {
		Tenant     string
		TimeRange  TimeRange
		Severities []Severity
		Statuses   []Status
		Fields     []string
		Filter     *Filter
		PageInput  *PageInput
	}
	GetRequest struct {
		Tenant          string
		ByOccurrenceIDs struct {
			IDs []string
		}
		ByEventIDs struct {
			EventIDs  []string
			TimeRange TimeRange
		}
		Statuses []Status
		Filter   *Filter
	}
	CountRequest struct {
		Query
		Fields []string
	}
	FrequencyRequest struct {
		Query
		Fields         []string
		GroupBy        []string
		Downsample     int64
		PersistCounts  bool
		CountInstances bool
	}

	Filter struct {
		Op    FilterOp
		Field string
		Value interface{}
	}
	SortOpt struct {
		Field string
		SortOrder
	}
	PageInput struct {
		Limit  uint64
		Cursor string
		SortBy []SortOpt
	}
	Page struct {
		Results []*Event
		HasNext bool
		Cursor  string
	}
	Repository interface {
		Create(context.Context, *Event) (*Event, error)
		Get(context.Context, *GetRequest) ([]*Event, error)
		Find(context.Context, *Query) (*Page, error)
		Update(context.Context, *Event) (*Event, error)
	}
)

func (tr TimeRange) IsValid() bool {
	return tr.End >= tr.Start
}

func (q *Query) Validate() error {
	if q == nil {
		return errors.New("nil query")
	}
	if len(q.Tenant) == 0 {
		return errors.New("missing tenant ID")
	}
	if !q.TimeRange.IsValid() {
		return errors.New("invalid timerange")
	}
	return nil
}

var supportedFields = map[string]bool{
	"name":         true,
	"summary":      true,
	"type":         true,
	"status":       true,
	"severity":     true,
	"entity":       true,
	"acknowledged": true,
	"dimensions":   true,
	"notes":        true,
}

func IsSupportedField(field string) bool {
	_, ok := supportedFields[field]
	return ok
}
