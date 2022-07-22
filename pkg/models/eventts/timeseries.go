package eventts

import (
	"context"
)

type Operation int32

const (
	Operation_OP_EXISTS        Operation = 0
	Operation_OP_EQUALS        Operation = 1
	Operation_OP_NOT_EQUALS    Operation = 2
	Operation_OP_CONTAINS      Operation = 3
	Operation_OP_LESS          Operation = 4
	Operation_OP_GREATER       Operation = 5
	Operation_OP_LESS_OR_EQ    Operation = 6
	Operation_OP_GREATER_OR_EQ Operation = 7
	Operation_OP_STARTS_WITH   Operation = 8
	Operation_OP_IN            Operation = 9
	Operation_OP_NOT_IN        Operation = 10
	Operation_OP_END_WITH      Operation = 11
	Operation_OP_NOT_CONTAINS  Operation = 12
)

type (
	TimeRange struct {
		Start, End int64
	}
	Occurrence struct {
		Timestamp int64
		ID        string
		EventID   string
		Metadata  map[string][]any
	}
	OccurrenceOptional struct {
		Result *Occurrence
		Err    error
	}
	OccurrenceInput struct {
		ID        string
		EventID   string
		TimeRange TimeRange
		IsActive  bool
	}
	Filter struct {
		Operation Operation
		Field     string
		Values    []any
	}
	EventTimeseriesInput struct {
		ByEventIDs struct {
			IDs       []string
			TimeRange TimeRange
		}
		ByOccurrences struct {
			OccurrenceMap map[string][]*OccurrenceInput
		}
		Latest  uint64
		Filters []*Filter
	}
	GetRequest struct {
		EventTimeseriesInput
	}
	CountResult struct {
		Value any
		Count uint64
	}
	CountResponse struct {
		Results map[string][]*CountResult
	}
	FrequencyRequest struct {
		EventTimeseriesInput
		Fields        []string
		GroupBy       []string
		Downsample    int64
		PersistCounts bool
	}
	FrequencyResult struct {
		Key    map[string]any
		Values []int64
	}
	FrequencyResponse struct {
		Timestamps []int64
		Results    []*FrequencyResult
	}
	FrequencyOptional struct {
		Result *FrequencyResponse
		Err    error
	}
	Repository interface {
		Get(context.Context, *GetRequest) ([]*Occurrence, error)
		GetStream(context.Context, *GetRequest) <-chan *OccurrenceOptional
		Frequency(context.Context, *FrequencyRequest) (*FrequencyResponse, error)
		FrequencyStream(context.Context, *FrequencyRequest) chan *FrequencyOptional
	}
)
