package grpc_test

import (
	"context"
	"sync"
	"testing"

	"github.com/zenoss/event-management-service/pkg/adapters/server/grpc"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	mocks2 "github.com/zenoss/event-management-service/internal/mocks"
	"github.com/zenoss/event-management-service/pkg/application/event/mocks"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/event-management-service/pkg/models/eventts"
	"github.com/zenoss/zenkit/v5"
	"github.com/zenoss/zing-proto/v11/go/cloud/eventquery"
	eventPb "github.com/zenoss/zing-proto/v11/go/event"
)

var _ = Describe("Query gRPC Service Tests", func() {
	var (
		testingT     = &testing.T{}
		ctx          context.Context
		cancel       context.CancelFunc
		svc          *grpc.EventQueryService
		appMock      *mocks.Service
		testIdentity = &mocks2.TenantIdentity{}
		testOnce     = sync.Once{}
	)
	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		testOnce.Do(func() {
			testIdentity.On("Tenant").Return("acme")
			// testIdentity.On("TenantID").Return("tenant1")
			testIdentity.On("TenantName").Return("tenant1")
		})
	})
	AfterEach(func() {
		cancel()
	})

	Context("Search", func() {
		BeforeEach(func() {
			appMock = mocks.NewService(testingT)
			svc = grpc.NewEventQueryService(appMock)
			ctx = zenkit.WithTenantIdentity(ctx, testIdentity)
		})
		It("should Find events", func() {
			appMock.On("Find", mock.Anything, mock.Anything).
				Return(&event.Page{
					Results: []*event.Event{
						{
							ID:     "event1",
							Tenant: "acme",
							Entity: "entity1",
							Name:   "event1",
							Dimensions: map[string]any{
								"dim1": "dimV1",
							},
							Occurrences: []*event.Occurrence{
								{
									ID:            "event1:1",
									EventID:       "event1",
									Tenant:        "acme",
									Summary:       "event1 summary",
									Body:          "event1 body",
									Type:          "testEvent",
									Status:        event.StatusOpen,
									Severity:      event.SeverityDefault,
									Acknowledged:  nil,
									StartTime:     1,
									EndTime:       0,
									CurrentTime:   5,
									InstanceCount: 2,
									Entity:        "entity1",
									Metadata: map[string][]any{
										"m1": {"v1"},
									},
								},
							},
							OccurrenceCount: 0,
						},
						{
							ID:     "event2",
							Tenant: "acme",
							Entity: "entity1",
							Name:   "event2",
							Dimensions: map[string]any{
								"dim1": "dimV1",
							},
							Occurrences: []*event.Occurrence{
								{
									ID:            "event2:1",
									EventID:       "event2",
									Tenant:        "acme",
									Summary:       "event2 summary",
									Body:          "event2 body",
									Type:          "testEvent",
									Status:        event.StatusOpen,
									Severity:      event.SeverityDefault,
									Acknowledged:  nil,
									StartTime:     1,
									EndTime:       0,
									CurrentTime:   5,
									InstanceCount: 2,
									Entity:        "entity1",
									Metadata: map[string][]any{
										"m1": {"v1"},
									},
								},
							},
							OccurrenceCount: 0,
						},
					},
				}, nil)
			resp, err := svc.Search(ctx, &eventquery.SearchRequest{
				Query: &eventquery.Query{
					TimeRange: &eventquery.TimeRange{
						Start: 0,
						End:   10,
					},
					Severities: []eventPb.Severity{eventPb.Severity_SEVERITY_DEFAULT},
					Statuses:   []eventPb.Status{eventPb.Status_STATUS_OPEN},
					PageInput: &eventquery.PageInput{
						Limit: 100,
					},
					Fields: nil,
				},
			})
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})
	})

	Context("Frequency", func() {
		BeforeEach(func() {
			appMock = mocks.NewService(testingT)
			svc = grpc.NewEventQueryService(appMock)
			ctx = zenkit.WithTenantIdentity(ctx, testIdentity)
		})
		It("should get Frequency", func() {
			appMock.On("Frequency", mock.Anything, mock.AnythingOfType("*event.FrequencyRequest")).
				Return(&eventts.FrequencyResponse{
					Timestamps: []int64{5, 10},
					Results: []*eventts.FrequencyResult{
						{
							Key: map[string]any{
								"severity": int(event.SeverityDefault),
							},
							Values: []int64{1, 1},
						},
					},
				}, nil)
			resp, err := svc.Frequency(ctx, &eventquery.FrequencyRequest{
				Query: &eventquery.Query{TimeRange: &eventquery.TimeRange{Start: 0, End: 10}, Severities: []eventPb.Severity{eventPb.Severity_SEVERITY_DEFAULT}, Statuses: []eventPb.Status{eventPb.Status_STATUS_OPEN}, PageInput: &eventquery.PageInput{Limit: 100}, Fields: nil},
				Fields: []*eventquery.Field{{
					Field:      "severity",
					Aggregator: eventquery.Aggregator_AGGREGATOR_FIRST,
					Label:      "",
				}},
				GroupBy: []*eventquery.Field{{
					Field:      "severity",
					Aggregator: eventquery.Aggregator_AGGREGATOR_FIRST,
					Label:      "",
				}},
				Downsample:     0,
				PersistCounts:  true,
				CountInstances: false,
			})
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})
	})
})
