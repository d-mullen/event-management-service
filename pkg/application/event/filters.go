package event

import (
	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/internal"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/event-management-service/pkg/models/eventts"
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
		event.FilterOpExists:               eventts.Operation_OP_EXISTS,
		event.FilterOpContains:             eventts.Operation_OP_CONTAINS,
		event.FilterOpDoesNotContain:       eventts.Operation_OP_NOT_CONTAINS,
		event.FilterOpPrefix:               eventts.Operation_OP_STARTS_WITH,
		event.FilterOpSuffix:               eventts.Operation_OP_END_WITH,
		//event.FilterOpRegex:           // not implemented in eventTS
		// event.FilterOpOr
		// event.FilterOpAnd
		// event.FilterOpNot
	}
)

func eventFilterToEventTSFilter(orig *event.Filter) ([]*eventts.Filter, error) {
	if orig == nil {
		return nil, errors.New("invalid filter")
	}
	results := make([]*eventts.Filter, 0)
	if orig.Op == event.FilterOpAnd {
		if others, ok := orig.Value.([]*event.Filter); ok {
			for _, filter := range others {
				newFilter, err := eventFilterToEventTSFilter(filter)
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

		if orig.Field != "entity" && orig.Field != "" { // trow already applied during eventContextRepo.Find

			if val, e := internal.AnyToSlice(orig.Value); e == nil {
				results = append(results, &eventts.Filter{
					Field:     orig.Field,
					Operation: op,
					Values:    val,
				})
			} else {
				return nil, errors.Errorf("invalid filter: %#v", orig)
			}
		}
		return results, nil
	}
	return nil, errors.Errorf("invalid filter: %#v", orig)
}
