package grpc_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zenoss/event-management-service/pkg/adapters/server/grpc"
	"github.com/zenoss/zing-proto/v11/go/cloud/eventquery"
	"github.com/zenoss/zing-proto/v11/go/event"

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
})
