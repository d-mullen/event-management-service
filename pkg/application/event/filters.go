package event

import (
	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/pkg/domain/event"
	"github.com/zenoss/event-management-service/pkg/domain/eventts"
)

var (
	eventToEventTSFilterOpMap = map[event.FilterOp]eventts.Operation{
		event.FilterOpLessThan:             eventts.Operation_OP_LESS,
		event.FilterOpLessThanOrEqualTo:    eventts.Operation_OP_LESS_OR_EQ,
		event.FilterOpGreaterThan:          eventts.Operation_OP_GREATER,
		event.FilterOpGreaterThanOrEqualTo: eventts.Operation_OP_GREATER_OR_EQ,
		event.FilterOpEqualTo:              eventts.Operation_OP_EQUALS,
		event.FilterOpNotEqualTo:           eventts.Operation_OP_NOT_EQUALS,
		event.FilterOpIn:                   eventts.Operation_OP_IN,
		event.FilterOpNotIn:                eventts.Operation_OP_NOT_IN,
		// event.FilterOpOr
		// event.FilterOpAnd
		// event.FilterOpNot
	}
)

func EventFilterToEventTSFilter(orig *event.Filter) ([]*eventts.Filter, error) {
	if orig == nil {
		return nil, errors.New("invalid filter")
	}
	results := make([]*eventts.Filter, 0)
	if orig.Op == event.FilterOpAnd {
		if others, ok := orig.Value.([]*event.Filter); ok {
			for _, filter := range others {
				newFilter, err := EventFilterToEventTSFilter(filter)
				if err != nil {
					results = append(results, newFilter...)
				}
			}
		} else {
			return nil, errors.Errorf("invalid filter: %#v", orig)
		}
		return results, nil
	}
	if op, ok := eventToEventTSFilterOpMap[orig.Op]; ok {
		results = append(results, &eventts.Filter{
			Field:     orig.Field,
			Operation: op,
			Values:    []any{orig.Value},
		})
		return results, nil
	}
	return nil, errors.Errorf("invalid filter: %#v", orig)
}
