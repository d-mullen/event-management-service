package grpc_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/zenoss/event-management-service/pkg/adapters/server/grpc"
	eventModel "github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/zing-proto/v11/go/cloud/eventquery"
	"github.com/zenoss/zing-proto/v11/go/event"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/structpb"
)

func mustToList(values ...any) *structpb.ListValue {
	lv, _ := structpb.NewList(values)
	return lv
}

func mustToValue(v any) *structpb.Value {
	vv, _ := structpb.NewValue(v)
	return vv
}

func mustToStruct(v map[string]any) *structpb.Struct {
	s, _ := structpb.NewStruct(v)
	return s
}

var _ = Describe("DTO Tests", func() {
	Context("SearchRequestToQuery", func() {
		It("should convert query protobufs to event queries", func() {
			// `{
			// 	"query": {
			// 	  "tenants": [
			// 		"qapreview"
			// 	  ],
			// 	  "timeRange": {
			// 		"start": "1652125380000",
			// 		"end": "1652136180000"
			// 	  },
			// 	  "clause": {
			// 		"filter": {
			// 		  "field": "entity",
			// 		  "operator": "EQUALS",
			// 		  "value": {
			// 			"stringVal": "AAAAA3A_TM2plJfjbV6V6Cn0ZQA="
			// 		  }
			// 		}
			// 	  },
			// 	  "severities": [
			// 		"SEVERITY_CRITICAL",
			// 		"SEVERITY_ERROR",
			// 		"SEVERITY_WARNING"
			// 	  ],
			// 	  "statuses": [
			// 		"STATUS_OPEN"
			// 	  ],
			// 	  "sortBy": [],
			// 	  "pageInput": {
			// 		"limit": "50"
			// 	  }
			// 	}
			//   }`
			req := &eventquery.SearchRequest{}
			q1, err := grpc.QueryProtoToEventQuery("acme", nil)
			Ω(err).Should(HaveOccurred())
			Ω(q1).Should(BeNil())

			req.Query = &eventquery.Query{
				TimeRange: &eventquery.TimeRange{Start: 0, End: 1000},
				Clause: &eventquery.Clause{
					Clause: &eventquery.Clause_Filter{
						Filter: &eventquery.Filter{
							Field:    "f1",
							Operator: eventquery.Filter_OPERATOR_EQUALS,
							Value:    mustToValue(1),
						},
					},
				},
			}

			r2, err := grpc.QueryProtoToEventQuery("acme", req.Query)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(r2).ShouldNot(BeNil())

			req.Query = &eventquery.Query{
				TimeRange: &eventquery.TimeRange{Start: 0, End: 1000},
				Clause: &eventquery.Clause{
					Clause: &eventquery.Clause_In{
						In: &eventquery.In{
							Field:  "f2",
							Values: mustToList(1, 2),
						},
					},
				},
			}
			r3, err := grpc.QueryProtoToEventQuery("acme", req.Query)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(r3).ShouldNot(BeNil())

			req.Query = &eventquery.Query{
				TimeRange: &eventquery.TimeRange{Start: 0, End: 1000},
				Clause: &eventquery.Clause{
					Clause: &eventquery.Clause_WithScope{
						WithScope: &eventquery.WithScope{
							Scope: &eventquery.WithScope_EntityIds{
								EntityIds: &eventquery.WithScope_WithIds{
									Ids: []string{"entity1", "entity2"},
								},
							},
						},
					},
				},
			}
			r4, err := grpc.QueryProtoToEventQuery("acme", req.Query)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(r4).ShouldNot(BeNil())
			GinkgoWriter.Printf("%#v", r4)

			r5, err := grpc.QueryProtoToEventQuery("acme", &eventquery.Query{
				TimeRange:  &eventquery.TimeRange{Start: 0, End: 1000},
				Statuses:   []event.Status{event.Status_STATUS_OPEN},
				Severities: []event.Severity{event.Severity_SEVERITY_CRITICAL},
				SortBy: []*eventquery.SortBy{
					{
						SortType: &eventquery.SortBy_ByField{
							ByField: &eventquery.SortByField{
								SortField: &eventquery.SortByField_Property{
									Property: "_id",
								},
								Order: eventquery.SortOrder_SORT_ORDER_ASC,
							},
						},
					},
					{
						SortType: &eventquery.SortBy_ByField{
							ByField: &eventquery.SortByField{
								SortField: &eventquery.SortByField_Property{
									Property: "lastSeen",
								},
								Order: eventquery.SortOrder_SORT_ORDER_DESC,
							},
						},
					},
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(r5).ShouldNot(BeNil())
			GinkgoWriter.Printf("%+v %+v", r5, r5.PageInput)
		})
	})
	Context("ClauseToProtoEventFilter", func() {
		DescribeTable("ClauseProtoToEventFilter Table-driven Tests",
			func(
				input *eventquery.Clause,
				expected *eventModel.Filter,
				expectedErr error,
			) {
				actualFilter, actualErr := grpc.ClauseProtoToEventFilter(input)
				if expectedErr != nil {
					Expect(actualErr).To(MatchError(expectedErr.Error()))
				} else {
					Expect(actualErr).ToNot(HaveOccurred())
				}
				if expected != nil {
					Expect(actualFilter).To(BeEquivalentTo(expected))
				}
			},
			Entry("nil clause",
				nil,
				nil,
				errors.New("nil clause"),
			),
			Entry("with-scope filter",
				&eventquery.Clause{
					Clause: &eventquery.Clause_WithScope{
						WithScope: &eventquery.WithScope{
							Scope: &eventquery.WithScope_EntityScopeCursor{
								EntityScopeCursor: &eventquery.WithScope_WithEntityScopeCursor{Cursor: "cursor_1"},
							},
						},
					},
				},
				&eventModel.Filter{
					Op:    eventModel.FilterOpScope,
					Field: "scope",
					Value: &eventModel.Scope{
						ScopeType: eventModel.ScopeEntity,
						Cursor:    "cursor_1",
					},
				},
				nil,
			),
			Entry("regex filter",
				&eventquery.Clause{
					Clause: &eventquery.Clause_Filter{
						Filter: &eventquery.Filter{
							Field:    "summary",
							Operator: eventquery.Filter_OPERATOR_REGEX,
							Value:    mustToValue("alarm"),
						},
					},
				},
				&eventModel.Filter{
					Op:    eventModel.FilterOpRegex,
					Field: "summary",
					Value: "alarm",
				},
				nil,
			),
			Entry("invalid filter",
				&eventquery.Clause{
					Clause: &eventquery.Clause_Filter{
						Filter: &eventquery.Filter{
							Field:    "summary",
							Operator: eventquery.Filter_Operator(9999),
							Value:    mustToValue("alarm"),
						},
					},
				},
				nil,
				status.Error(codes.InvalidArgument, "invalid filter"),
			),
			Entry("not clause",
				&eventquery.Clause{
					Clause: &eventquery.Clause_Not_{
						Not: &eventquery.Clause_Not{
							Clause: &eventquery.Clause{
								Clause: &eventquery.Clause_Filter{Filter: &eventquery.Filter{
									Field:    "f1",
									Operator: eventquery.Filter_OPERATOR_EQUALS,
									Value:    mustToValue(1),
								}},
							},
						},
					},
				},
				&eventModel.Filter{
					Op: eventModel.FilterOpNot,
					Value: &eventModel.Filter{
						Op:    eventModel.FilterOpEqualTo,
						Field: "f1",
						Value: float64(1),
					},
				},
				nil,
			),
			Entry("and clause",
				&eventquery.Clause{
					Clause: &eventquery.Clause_And_{
						And: &eventquery.Clause_And{
							Clauses: []*eventquery.Clause{
								{
									Clause: &eventquery.Clause_Filter{Filter: &eventquery.Filter{
										Field:    "f1",
										Operator: eventquery.Filter_OPERATOR_EQUALS,
										Value:    mustToValue(1),
									}},
								},
								{
									Clause: &eventquery.Clause_Filter{Filter: &eventquery.Filter{
										Field:    "f2",
										Operator: eventquery.Filter_OPERATOR_NOT_EQUALS,
										Value:    mustToValue("foo"),
									}},
								},
							},
						},
					},
				},
				&eventModel.Filter{
					Op: eventModel.FilterOpAnd,
					Value: []*eventModel.Filter{
						{
							Op:    eventModel.FilterOpEqualTo,
							Field: "f1",
							Value: float64(1),
						},
						{
							Op:    eventModel.FilterOpNotEqualTo,
							Field: "f2",
							Value: "foo",
						},
					},
				},
				nil,
			),
			Entry("or clause",
				&eventquery.Clause{
					Clause: &eventquery.Clause_Or_{
						Or: &eventquery.Clause_Or{
							Clauses: []*eventquery.Clause{
								{
									Clause: &eventquery.Clause_Filter{Filter: &eventquery.Filter{
										Field:    "f1",
										Operator: eventquery.Filter_OPERATOR_EQUALS,
										Value:    mustToValue(1),
									}},
								},
								{
									Clause: &eventquery.Clause_Filter{Filter: &eventquery.Filter{
										Field:    "f2",
										Operator: eventquery.Filter_OPERATOR_NOT_EQUALS,
										Value:    mustToValue("foo"),
									}},
								},
							},
						},
					},
				},
				&eventModel.Filter{
					Op: eventModel.FilterOpOr,
					Value: []*eventModel.Filter{
						{
							Op:    eventModel.FilterOpEqualTo,
							Field: "f1",
							Value: float64(1),
						},
						{
							Op:    eventModel.FilterOpNotEqualTo,
							Field: "f2",
							Value: "foo",
						},
					},
				},
				nil,
			),
		)
	})
})
