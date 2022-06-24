package mongodb

import (
	"reflect"
	"time"

	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/pkg/domain/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	defaultTimeRange event.TimeRange
)

func getOccurrenceTemporalFilters(activeEventsOnly bool, tr event.TimeRange) (bson.D, error) {
	if tr.End < tr.Start {
		return nil, errors.New("invalid time range")
	}
	start, end := tr.Start, tr.End
	if tr == defaultTimeRange {
		now := time.Now()
		end = now.UnixMilli()
		start = now.Add(-24 * time.Hour).UnixMilli() // TODO: expose this as config
	}
	if activeEventsOnly {
		return bson.D{{Key: "startTime", Value: bson.D{{Key: "$lte", Value: end}}}}, nil
	}
	return bson.D{
		{Key: "$or", Value: bson.A{
			bson.D{{Key: "$and", Value: bson.A{
				bson.D{{Key: "status", Value: event.StatusClosed}},
				bson.D{{Key: "startTime", Value: bson.D{{Key: "$lte", Value: end}}}},
				bson.D{{Key: "endTime", Value: bson.D{{Key: "$gte", Value: start}}}},
			}}},
			bson.D{{Key: "startTime", Value: bson.D{{Key: "$lte", Value: end}}}},
		}},
	}, nil
}

const (
	MongoOpEqualTo              = "$eq"
	MongoOpLessThan             = "$lt"
	MongoOpGreaterThan          = "$gt"
	MongoOpGreaterThanOrEqualTo = "$gte"
	MongoOpLessThanOrEqualTo    = "$lte"
	MongoOpNotEqualTo           = "$ne"
	MongoOpIn                   = "$in"
	MongoOpNotIn                = "$nin"
	MongoOpOr                   = "$or"
	MongoOpAnd                  = "$and"
)

var domainMongoFilterMap = map[event.FilterOp]string{
	event.FilterOpLessThan:             MongoOpLessThan,
	event.FilterOpLessThanOrEqualTo:    MongoOpLessThanOrEqualTo,
	event.FilterOpGreaterThan:          MongoOpGreaterThan,
	event.FilterOpGreaterThanOrEqualTo: MongoOpGreaterThanOrEqualTo,
	event.FilterOpEqualTo:              MongoOpEqualTo,
	event.FilterOpNotEqualTo:           MongoOpNotEqualTo,
	event.FilterOpIn:                   MongoOpIn,
	event.FilterOpNotIn:                MongoOpNotIn,
	event.FilterOpOr:                   MongoOpOr,
	event.FilterOpAnd:                  MongoOpAnd,
}

func DomainFilterToMongoD(orig *event.Filter) (bson.D, error) {
	op, ok := domainMongoFilterMap[orig.Op]
	if !ok {
		return nil, errors.New("invalid filter operation")
	}
	switch orig.Op {
	case event.FilterOpAnd, event.FilterOpOr:
		if reflect.TypeOf(orig.Value).Kind() != reflect.Slice {
			return nil, errors.New("invalid filter values")
		}
		t := reflect.ValueOf(orig.Value)
		filterValues := make(bson.A, t.Len())
		filterKind := reflect.TypeOf((*event.Filter)(nil)).Kind()
		for i := 0; i < t.Len(); i++ {
			v := t.Index(i)
			if v.Kind() != filterKind {
				return nil, errors.Errorf("invalid filter values: %v", orig.Value)
			}
			otherFilter := v.Interface().(*event.Filter)
			newFilter, err := DomainFilterToMongoD(otherFilter)
			if err != nil {
				return nil, err
			}
			filterValues[i] = newFilter

		}
		return bson.D{{Key: orig.Field, Value: bson.D{{Key: op, Value: filterValues}}}}, nil
	case event.FilterOpIn, event.FilterOpNotIn:
		if reflect.TypeOf(orig.Value).Kind() != reflect.Slice {
			return nil, errors.New("invalid filter values")
		}
		return bson.D{{Key: orig.Field, Value: bson.D{{Key: op, Value: orig.Value}}}}, nil
	default:
		return bson.D{{Key: orig.Field, Value: bson.D{{Key: op, Value: orig.Value}}}}, nil
	}
}

func isActiveEventsOnly(q *event.Query) bool {
	for _, status := range q.Statuses {
		if status == event.StatusClosed {
			return false
		}
	}
	return true
}

func QueryToFindArguments(query *event.Query) (bson.D, []*options.FindOptions, error) {
	filters := bson.D{{Key: "tenantId", Value: query.Tenant}}
	moreFilters, err := getOccurrenceTemporalFilters(isActiveEventsOnly(query), query.TimeRange)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to make query plan")
	}
	filters = append(filters, moreFilters...)
	if len(query.Statuses) > 0 {
		if len(query.Statuses) == 1 {
			filters = append(filters, bson.E{Key: "status", Value: query.Statuses[0]})
		} else {
			filters = append(filters, bson.E{Key: "status", Value: bson.D{{Key: MongoOpIn, Value: query.Statuses}}})
		}
	}
	if len(query.Severities) > 0 {
		if len(query.Severities) == 1 {
			filters = append(filters, bson.E{Key: "severity", Value: query.Severities[0]})
		} else {
			filters = append(filters, bson.E{Key: "severity", Value: bson.D{{Key: MongoOpIn, Value: query.Severities}}})
		}
	}
	if query.Filter != nil {
		anotherFilter, err := DomainFilterToMongoD(query.Filter)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to convert domain filter to mongo filter")
		}
		filters = append(filters, anotherFilter...)
	}
	findOpts := make([]*options.FindOptions, 0)
	if query.PageInput != nil && len(query.PageInput.SortBy) > 0 {
		sortBy := query.PageInput.SortBy[0]
		findOpts = append(findOpts, options.Find().SetSort(bson.D{{Key: sortBy.Field, Value: sortBy.SortOrder}}))
	}
	if query.PageInput != nil && query.PageInput.Limit > 0 {
		findOpts = append(findOpts, options.Find().SetLimit(int64(query.PageInput.Limit+1)))
	}
	return filters, findOpts, nil
}

type decodable interface {
	*bson.M | *bson.D
}

func defaultDecodeFunc[D decodable](doc D, dest interface{}) error {
	docBytes, err := bson.Marshal(doc)
	if err != nil {
		return err
	}
	return bson.Unmarshal(docBytes, dest)
}
