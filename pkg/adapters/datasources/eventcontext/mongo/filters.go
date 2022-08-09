package mongo

import (
	"reflect"
	"time"

	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	OpEqualTo              = "$eq"
	OpLessThan             = "$lt"
	OpGreaterThan          = "$gt"
	OpGreaterThanOrEqualTo = "$gte"
	OpLessThanOrEqualTo    = "$lte"
	OpNotEqualTo           = "$ne"
	OpIn                   = "$in"
	OpNotIn                = "$nin"
	OpOr                   = "$or"
	OpAnd                  = "$and"
	OpNot                  = "$not"
)

var domainMongoFilterMap = map[event.FilterOp]string{
	event.FilterOpLessThan:             OpLessThan,
	event.FilterOpLessThanOrEqualTo:    OpLessThanOrEqualTo,
	event.FilterOpGreaterThan:          OpGreaterThan,
	event.FilterOpGreaterThanOrEqualTo: OpGreaterThanOrEqualTo,
	event.FilterOpEqualTo:              OpEqualTo,
	event.FilterOpNotEqualTo:           OpNotEqualTo,
	event.FilterOpIn:                   OpIn,
	event.FilterOpNotIn:                OpNotIn,
	event.FilterOpOr:                   OpOr,
	event.FilterOpAnd:                  OpAnd,
	event.FilterOpNot:                  OpNot,
}

func ApplyNotFilterTransform(orig *event.Filter) (bson.D, error) {
	filterKind := reflect.TypeOf((*event.Filter)(nil)).Kind()
	v := reflect.ValueOf(orig.Value)
	if v.Kind() != filterKind {
		return bson.D{{Key: orig.Field, Value: bson.E{Key: OpNotEqualTo, Value: orig.Value}}}, nil
	}
	otherFilter := v.Interface().(*event.Filter)
	switch otherFilter.Op {
	case event.FilterOpEqualTo:
		return bson.D{{Key: otherFilter.Field, Value: bson.E{Key: OpNotEqualTo, Value: otherFilter.Value}}}, nil
	case event.FilterOpIn:
		return bson.D{{Key: otherFilter.Field, Value: bson.E{Key: OpNotIn, Value: otherFilter.Value}}}, nil
	case event.FilterOpAnd, event.FilterOpOr, event.FilterOpNot:
		newFilter, err := DomainFilterToMongoD(otherFilter)
		if err != nil {
			return nil, err
		}
		return bson.D{{Key: OpNot, Value: newFilter}}, nil
	default:
		otherOperator, ok := domainMongoFilterMap[otherFilter.Op]
		if !ok {
			return nil, errors.New("invalid filter operation")
		}
		return bson.D{
			{Key: otherFilter.Field, Value: bson.E{Key: OpNot, Value: bson.E{Key: otherOperator, Value: otherFilter.Value}}}}, nil

	}

}

func DomainFilterToMongoD(orig *event.Filter) (bson.D, error) {
	if orig == nil {
		return nil, errors.New("invalid filter: got nil value")
	}
	op, ok := domainMongoFilterMap[orig.Op]
	if !ok {
		return nil, errors.Errorf("invalid filter operation: %v", orig.Op)
	}
	switch orig.Op {
	case event.FilterOpNot:
		return ApplyNotFilterTransform(orig)
	case event.FilterOpAnd, event.FilterOpOr:
		if reflect.TypeOf(orig.Value).Kind() != reflect.Slice {
			return nil, errors.New("invalid filter values")
		}
		t := reflect.ValueOf(orig.Value)
		filterValues := make(bson.A, t.Len())
		filterKind := reflect.TypeOf((*event.Filter)(nil)).Kind()
		for i := 0; i < t.Len(); i++ {
			v := t.Index(i)
			otherFilter, ok := v.Interface().(*event.Filter)
			if !ok {
				return nil, errors.Errorf("invalid filter values: %v %v != %T", orig.Value, v.Kind(), filterKind)
			}
			newFilter, err := DomainFilterToMongoD(otherFilter)
			if err != nil {
				return nil, err
			}
			filterValues[i] = newFilter

		}
		return bson.D{{Key: op, Value: filterValues}}, nil
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

func QueryToFindArguments(query *event.Query) (bson.D, *options.FindOptions, error) {
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
			filters = append(filters, bson.E{Key: "status", Value: bson.D{{Key: OpIn, Value: query.Statuses}}})
		}
	}
	if len(query.Severities) > 0 {
		if len(query.Severities) == 1 {
			filters = append(filters, bson.E{Key: "severity", Value: query.Severities[0]})
		} else {
			filters = append(filters, bson.E{Key: "severity", Value: bson.D{{Key: OpIn, Value: query.Severities}}})
		}
	}
	if query.Filter != nil {
		anotherFilter, err := DomainFilterToMongoD(query.Filter)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to convert domain filter to mongo filter")
		}
		filters = append(filters, anotherFilter...)
	}
	findOpts := options.Find()
	if query.PageInput != nil && len(query.PageInput.SortBy) > 0 {
		sortBy := query.PageInput.SortBy[0]
		findOpts.SetSort(&bson.D{{Key: sortBy.Field, Value: sortBy.SortOrder}})
	}
	if query.PageInput != nil && query.PageInput.Limit > 0 {
		findOpts.SetLimit(int64(query.PageInput.Limit + 1))
	}
	return filters, findOpts, nil
}

type decodable interface {
	*bson.M | *bson.D
}

func defaultDecodeFunc[D decodable](doc D, dest any) error {
	docBytes, err := bson.Marshal(doc)
	if err != nil {
		return err
	}
	return bson.Unmarshal(docBytes, dest)
}

type (
	MongoQuery struct {
		Filter   primitive.D
		FindOpts *options.FindOptions
	}
)

func NextKey(sortField string, items []*bson.M) *bson.D {
	if len(items) == 0 {
		return nil
	}
	item := items[len(items)-1]
	itemAsMap := make(map[string]any)
	b, err := bson.Marshal(item)
	if err != nil {
		return nil
	}
	bson.Unmarshal(b, itemAsMap)
	if len(sortField) == 0 {
		return &bson.D{{Key: "_id", Value: itemAsMap["_id"]}}
	}
	return &bson.D{
		{Key: "_id", Value: itemAsMap["_id"]},
		{Key: sortField, Value: itemAsMap[sortField]},
	}
}

func GeneratePaginationQuery(filter, sort, nextKey *bson.D) *bson.D {
	var (
		paginatedQuery bson.D
		sortField      string
		sortOperator   string
	)

	// if next key is nil, return the query unmodified
	if nextKey == nil {
		return filter
	}

	// copy the original query filter
	if filter != nil {
		for k, v := range filter.Map() {
			paginatedQuery = append(paginatedQuery, bson.E{k, v})
		}
	}

	if sort == nil {
		paginatedQuery = append(paginatedQuery,
			bson.E{Key: "_id", Value: bson.D{{Key: OpGreaterThan, Value: nextKey.Map()["_id"]}}})
		return &paginatedQuery
	}

	i := 0
	for k, v := range sort.Map() {
		if i > 0 {
			break
		}
		sortField = k
		sortOperator = OpGreaterThan
		if iv, ok := v.(event.SortOrder); ok && iv != event.SortOrderAscending {
			sortOperator = OpLessThan
		}
		i++
	}
	pqM := bson.D{
		{Key: sortField, Value: bson.D{{Key: sortOperator, Value: nextKey.Map()[sortField]}}},
		{Key: OpAnd, Value: bson.D{
			{Key: sortField, Value: nextKey.Map()[sortField]},
			{Key: "_id", Value: bson.D{{Key: sortOperator, Value: nextKey.Map()[sortField]}}},
		}},
	}

	if _, ok2 := paginatedQuery.Map()[OpOr]; !ok2 {
		paginatedQuery = append(paginatedQuery, bson.E{Key: OpOr, Value: pqM})
	} else {
		paginatedQuery = bson.D{
			{Key: OpAnd, Value: bson.A{
				filter,
				bson.D{{Key: OpOr, Value: pqM}},
			}},
		}
	}
	return &paginatedQuery

}
