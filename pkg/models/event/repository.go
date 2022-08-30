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
	FilterOpScope                FilterOp = "_in_scope"
	FilterOpContains             FilterOp = "contains"
	FilterOpExists               FilterOp = "exists"
	FilterOpDoesNotContain       FilterOp = "not_contain"
	FilterOpPrefix               FilterOp = "prefix"
	FilterOpSuffix               FilterOp = "suffix"
	FilterOpRegex                FilterOp = "regex"
)

type SortOrder int

const (
	SortOrderDescending = -1
	SortOrderAscending  = 1
)

type ScopeEnum int

const (
	ScopeEntity ScopeEnum = iota
)

type PageDirection int

const (
	PageDirectionForward = iota
	PageDirectionBackward
)

type (
	TimeRange struct {
		Start int64 `json:"start"`
		End   int64 `json:"end"`
	}

	Scope struct {
		ScopeType ScopeEnum
		Cursor    string
	}
	Query struct {
		Tenant     string     `json:"tenant,omitempty"`
		TimeRange  TimeRange  `json:"timeRange,omitempty"`
		Severities []Severity `json:"severities,omitempty"`
		Statuses   []Status   `json:"statuses,omitempty"`
		Fields     []string   `json:"fields,omitempty"`
		Filter     *Filter    `json:"filter,omitempty"`
		PageInput  *PageInput `json:"pageInput,omitempty"`
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
		Op    FilterOp `json:"op"`
		Field string   `json:"field"`
		Value any      `json:"value"`
	}
	SortOpt struct {
		Field string
		SortOrder
	}
	PageInput struct {
		Direction PageDirection
		Limit     uint64
		Cursor    string
		SortBy    []SortOpt
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
	//	"name":         true,
	"summary":       true,
	"type":          true,
	"status":        true,
	"severity":      true,
	"entity":        true,
	"acknowledged":  true,
	"dimensions":    true,
	"notes":         true,
	"eventId":       true,
	"createdAt":     true,
	"updatedAt":     true,
	"expireAt":      true,
	"tenantId":      true,
	"currentTime":   true,
	"instanceCount": true,
	"body":          true,
	"startTime":     true,
	"endTime":       true,
}

func IsSupportedField(field string) bool {
	_, ok := supportedFields[field]
	return ok
}
