package grpc

import (
	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/zing-proto/v11/go/cloud/eventquery"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
)

var (
	errInvalidClause = status.Error(codes.InvalidArgument, "invalid clause")
	errInvalidFilter = status.Error(codes.InvalidArgument, "invalid filter")
)

var filterOpProtoEventFilterMap = map[string]event.FilterOp{
	eventquery.Filter_OPERATOR_EQUALS.String():                 event.FilterOpEqualTo,
	eventquery.Filter_OPERATOR_LESS_THAN.String():              event.FilterOpLessThan,
	eventquery.Filter_OPERATOR_LESS_THEN_OR_EQUALS.String():    event.FilterOpLessThanOrEqualTo,
	eventquery.Filter_OPERATOR_GREATER_THAN.String():           event.FilterOpGreaterThan,
	eventquery.Filter_OPERATOR_GREATER_THAN_OR_EQUALS.String(): event.FilterOpGreaterThanOrEqualTo,
	eventquery.Filter_OPERATOR_REGEX.String():                  event.FilterOpRegex,
	eventquery.Filter_OPERATOR_PREFIX.String():                 event.FilterOpPrefix,
	eventquery.Filter_OPERATOR_SUFFIX.String():                 event.FilterOpSuffix,
	eventquery.Filter_OPERATOR_CONTAINS.String():               event.FilterOpContains,
	eventquery.Filter_OPERATOR_DOES_NOT_CONTAIN.String():       event.FilterOpDoesNotContain,
	eventquery.Filter_OPERATOR_EXISTS.String():                 event.FilterOpExists,
	eventquery.Filter_OPERATOR_NOT_EQUALS.String():             event.FilterOpNotEqualTo,
}

func FilterProtoToEventFilter(filterPb *eventquery.Filter) (*event.Filter, error) {
	op, ok := filterOpProtoEventFilterMap[filterPb.Operator.String()]
	if !ok {
		return nil, errInvalidFilter
	}
	return &event.Filter{
		Field: filterPb.Field,
		Op:    op,
		Value: filterPb.GetValue().AsInterface(),
	}, nil
}

func WithScopeProtoToDomainFilter(pb *eventquery.WithScope) (*event.Filter, error) {
	switch v := pb.GetScope().(type) {
	case *eventquery.WithScope_EntityIds:
		if v.EntityIds == nil || len(v.EntityIds.Ids) == 0 {
			return nil, status.Error(codes.InvalidArgument, "invalid entity id scope")
		}
		return &event.Filter{
			Field: "entity",
			Op:    event.FilterOpIn,
			Value: v.EntityIds.Ids,
		}, nil
	case *eventquery.WithScope_EntityScopeCursor:
		return &event.Filter{
			Field: "scope",
			Op:    event.FilterOpScope,
			Value: &event.Scope{
				ScopeType: event.ScopeEntity,
				Cursor:    v.EntityScopeCursor.GetCursor(),
			},
		}, nil
	}
	return nil, errInvalidFilter
}

func valuePbSliceToAnySlice(pbSlice []*structpb.Value) []any {
	results := make([]any, len(pbSlice))
	for i, valuePb := range pbSlice {
		results[i] = valuePb.AsInterface()
	}
	return results
}

func ClauseProtoToEventFilter(clause *eventquery.Clause) (*event.Filter, error) {
	if clause == nil {
		return nil, errors.New("nil clause")
	}
	result := &event.Filter{}
	// Types that are assignable to Clause:
	//    *Clause_And_
	//    *Clause_Or_
	//    *Clause_Not_
	//    *Clause_Filter
	//    *Clause_In
	//    *Clause_WithScope
	switch v := clause.GetClause().(type) {
	case *eventquery.Clause_Not_:
		cl, err := ClauseProtoToEventFilter(v.Not.GetClause())
		if err != nil {
			return nil, err
		}
		return &event.Filter{
			Op:    event.FilterOpNot,
			Value: cl,
		}, nil
	case *eventquery.Clause_And_:
		result.Op = event.FilterOpAnd
		others := make([]*event.Filter, 0)
		for _, otherPb := range v.And.GetClauses() {
			other, err := ClauseProtoToEventFilter(otherPb)
			if err != nil {
				return nil, err
			}
			others = append(others, other)
		}
		result.Value = others
		return result, nil
	case *eventquery.Clause_Or_:
		result.Op = event.FilterOpOr
		others := make([]*event.Filter, 0)
		for _, otherPb := range v.Or.GetClauses() {
			other, err := ClauseProtoToEventFilter(otherPb)
			if err != nil {
				return nil, err
			}
			others = append(others, other)
		}
		result.Value = others
		return result, nil
	case *eventquery.Clause_Filter:
		return FilterProtoToEventFilter(v.Filter)
	case *eventquery.Clause_In:
		return &event.Filter{
			Op:    event.FilterOpIn,
			Field: v.In.GetField(),
			Value: valuePbSliceToAnySlice(v.In.GetValues()),
		}, nil
	case *eventquery.Clause_WithScope:
		return WithScopeProtoToDomainFilter(v.WithScope)
	}
	return nil, errInvalidClause
}
