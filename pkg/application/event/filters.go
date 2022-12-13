package event

import (
	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/internal"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/event-management-service/pkg/models/eventts"
)

var eventToEventTSFilterOpMap = map[event.FilterOp]eventts.Operation{
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
	event.FilterOpRegex:                eventts.Operation_OP_CONTAINS,
	// not implemented in eventTS
	// event.FilterOpOr
	// event.FilterOpAnd
	// event.FilterOpNot
}

func eventFilterToEventTSFilter(orig *event.Filter) ([]*eventts.Filter, error) {
	if orig == nil {
		return nil, nil
	}
	results := make([]*eventts.Filter, 0)
	if orig.Op == event.FilterOpAnd {
		if others, ok := orig.Value.([]*event.Filter); ok {
			for _, filter := range others {
				// Do not apply filters that will be applied in event-context store
				if event.IsSupportedField(filter.Field) {
					continue
				}
				newFilter, err := eventFilterToEventTSFilter(filter)
				if err != nil {
					return nil, errors.Wrapf(err, "failed to convert to eventts filter: %v", filter)
				} else {
					results = append(results, newFilter...)
				}
			}
		} else {
			return nil, errors.Errorf("invalid filter: %#v", orig)
		}
		return results, nil
	}
	if op, ok := eventToEventTSFilterOpMap[orig.Op]; ok {

		// Do not apply filters that will be applied in event-context store
		if orig.Field != "entity" && orig.Field != "" && !event.IsSupportedField(orig.Field) { // trow already applied during eventContextRepo.Find
			results = append(results, &eventts.Filter{
				Field:     orig.Field,
				Operation: op,
				Values:    internal.AnyToSlice(orig.Value),
			})
		}
		return results, nil
	}
	return nil, errors.Errorf("invalid filter: %#v", orig)
}
