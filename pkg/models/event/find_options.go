package event

import (
	"context"

	"github.com/pkg/errors"
)

type (
	FindOption struct {
		OccurrenceProcessors []OccurrenceProcessor
	}

	OccurrenceProcessor func(context.Context, []*Occurrence) ([]*Occurrence, error)
)

func MergeFindOptions(opts ...*FindOption) *FindOption {
	result := &FindOption{}
	for _, opt := range opts {
		if len(opt.OccurrenceProcessors) > 0 {
			if result.OccurrenceProcessors == nil {
				result.OccurrenceProcessors = make([]OccurrenceProcessor, 0)
			}
			result.OccurrenceProcessors = append(result.OccurrenceProcessors, opt.OccurrenceProcessors...)
		}
	}
	return result
}

func NewFindOption() *FindOption { return &FindOption{} }

func (opt *FindOption) SetOccurrenceProcessors(proc []OccurrenceProcessor) {
	if opt == nil {
		opt = &FindOption{}
	}
	opt.OccurrenceProcessors = proc
}

func (opt *FindOption) ApplyOccurrenceProcessors(ctx context.Context, orig []*Occurrence) ([]*Occurrence, error) {
	var err error
	results := make([]*Occurrence, len(orig))
	if opt == nil || len(opt.OccurrenceProcessors) == 0 {
		return orig, nil
	}
	copy(results, orig)
	for _, proc := range opt.OccurrenceProcessors {
		results, err = proc(ctx, results)
		if err != nil {
			return nil, errors.Wrap(err, "failed to execute occurrence processor")
		}
		if len(results) == 0 {
			break
		}
	}
	return results, nil
}
