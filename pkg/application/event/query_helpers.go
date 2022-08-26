package event

import (
	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"reflect"
)

func SplitOutQueries(batchSize int, targetField string, origQuery *event.Query) []*event.Query {
	var (
		results           = make([]*event.Query, 0)
		defaultStatuses   = []event.Status{event.StatusOpen, event.StatusClosed, event.StatusDefault, event.StatusSuppressed}
		defaultSeverities = []event.Severity{event.SeverityCritical, event.SeverityError, event.SeverityWarning, event.SeverityDebug, event.SeverityInfo}
		statuses          []event.Status
		severities        []event.Severity
	)

	if len(origQuery.Statuses) == 0 {
		statuses = defaultStatuses
	} else {
		statuses = origQuery.Statuses
	}
	if len(origQuery.Severities) == 0 {
		severities = defaultSeverities
	} else {
		severities = origQuery.Severities
	}

	for _, currStatus := range statuses {
		for _, currSeverity := range severities {
			newFilters := SplitOutFilter(batchSize, targetField, origQuery.Filter)
			if len(newFilters) > 0 {
				for _, filter := range newFilters {
					results = append(results, &event.Query{
						Tenant:     origQuery.Tenant,
						TimeRange:  origQuery.TimeRange,
						Severities: []event.Severity{currSeverity},
						Statuses:   []event.Status{currStatus},
						Fields:     origQuery.Fields,
						PageInput:  origQuery.PageInput,
						Filter:     filter,
					})
				}
			} else {
				results = append(results, &event.Query{
					Tenant:     origQuery.Tenant,
					TimeRange:  origQuery.TimeRange,
					Severities: []event.Severity{currSeverity},
					Statuses:   []event.Status{currStatus},
					Fields:     origQuery.Fields,
					PageInput:  origQuery.PageInput,
				})
			}
		}
	}
	return results
}

func batchUpFilter[T any](batchSize int, fieldName string, op event.FilterOp, values []T) []*event.Filter {
	results := make([]*event.Filter, 0)
	for i := 0; i < len(values); i += batchSize {
		j := i + batchSize
		if j > len(values) {
			j = len(values)
		}
		results = append(results, &event.Filter{
			Field: fieldName,
			Op:    op,
			Value: values[i:j],
		})
	}
	return results
}

func anyToSlice(v any) []any {
	var (
		out []any
	)
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i++ {
			out = append(out, rv.Index(i).Interface())
		}
	}
	return out
}

func SplitOutFilter(batchSize int, targetField string, filter *event.Filter) []*event.Filter {
	if filter == nil {
		return nil
	}
	switch filter.Op {
	case event.FilterOpScope, event.FilterOpIn, event.FilterOpNotIn:
		if k := reflect.TypeOf(filter.Value); filter.Field == targetField && k.Kind() == reflect.Slice {
			return batchUpFilter(batchSize, filter.Field, filter.Op, anyToSlice(filter.Value))
		}
	case event.FilterOpAnd:
		if others, ok := filter.Value.([]*event.Filter); ok {
			index := -1
			for i, other := range others {
				if other.Field == targetField {
					index = i
					break
				}
			}
			if index > -1 {
				newFilters := SplitOutFilter(batchSize, targetField, others[index])
				if len(newFilters) > 0 {
					out := make([]*event.Filter, 0)
					for _, newFilter := range newFilters {
						_newOthers := make([]*event.Filter, 0)
						_newOthers = append(_newOthers, others[:index]...)
						_newOthers = append(_newOthers, newFilter)
						_newOthers = append(_newOthers, others[index:]...)
						out = append(out, &event.Filter{
							Field: filter.Field,
							Op:    filter.Op,
							Value: _newOthers,
						})
					}
					return out
				}
			}
		}
	}
	return nil
}

func transformFilter(orig *event.Filter, cb func(*event.Filter) (*event.Filter, bool, error)) (*event.Filter, error) {
	if orig == nil {
		return nil, nil
	}
	switch orig.Op {
	case event.FilterOpNot:
		other, isFilter := orig.Value.(*event.Filter)
		if isFilter {
			newFilter, ok2, err := cb(other)
			if err != nil {
				return nil, errors.Wrap(err, "failed to transform filter")
			}
			if ok2 {
				return &event.Filter{
					Field: orig.Field,
					Op:    orig.Op,
					Value: newFilter,
				}, nil
			}
		}
	case event.FilterOpAnd, event.FilterOpOr:
		others, isFilterSlice := orig.Value.([]*event.Filter)
		if isFilterSlice {
			newSubclauses := make([]*event.Filter, 0)
			for _, other := range others {
				newFilter, ok2, err := cb(other)
				if err != nil {
					return nil, errors.Wrap(err, "failed to transform filter")
				}
				if ok2 {
					newSubclauses = append(newSubclauses, newFilter)
				}
			}
			return &event.Filter{
				Field: orig.Field,
				Op:    orig.Op,
				Value: newSubclauses,
			}, nil
		}
	default:
		newFilter, ok, err := cb(orig)
		if err != nil {
			return nil, errors.Wrap(err, "failed to transform filter")
		}
		if ok {
			return newFilter, nil
		}
		return nil, nil
	}
	return orig, nil
}
