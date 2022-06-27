package grpc_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/grpc"
	"github.com/zenoss/zing-proto/v11/go/cloud/eventquery"

	"google.golang.org/protobuf/types/known/structpb"
)

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
			q1, err := grpc.QueryProtoToEventQuery("acme", nil, nil)
			Ω(err).Should(HaveOccurred())
			Ω(q1).Should(BeNil())

			req.Query = &eventquery.Query{
				TimeRange: &eventquery.TimeRange{Start: 0, End: 1000},
				Clause: &eventquery.Clause{
					Clause: &eventquery.Clause_Filter{
						Filter: &eventquery.Filter{
							Field:    "f1",
							Operator: eventquery.Filter_EQUALS,
							Value:    mustToStruct(map[string]any{"f1": 1}),
						},
					},
				},
			}

			r2, err := grpc.QueryProtoToEventQuery("acme", req.Query, nil)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(r2).ShouldNot(BeNil())

			req.Query = &eventquery.Query{
				TimeRange: &eventquery.TimeRange{Start: 0, End: 1000},
				Clause: &eventquery.Clause{
					Clause: &eventquery.Clause_In{
						In: &eventquery.In{
							Field:  "f2",
							Values: mustToStruct(map[string]any{"f2": []any{1, 2}}),
						},
					},
				},
			}
			r3, err := grpc.QueryProtoToEventQuery("acme", req.Query, nil)
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
			r4, err := grpc.QueryProtoToEventQuery("acme", req.Query, nil)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(r4).ShouldNot(BeNil())
			GinkgoWriter.Printf("%#v", r4)
		})
	})
})
