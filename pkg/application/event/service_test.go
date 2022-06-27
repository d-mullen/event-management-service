package event_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/zenoss/event-management-service/pkg/application/event"
	eventContext "github.com/zenoss/event-management-service/pkg/domain/event"
	eventContextMocks "github.com/zenoss/event-management-service/pkg/domain/event/mocks"
	"github.com/zenoss/event-management-service/pkg/domain/eventts"
	eventTSMocks "github.com/zenoss/event-management-service/pkg/domain/eventts/mocks"
)

var _ = Describe("eventquery.Service", func() {
	var (
		svc         event.Service
		mockCtx     = mock.Anything
		mockQuery   = mock.AnythingOfType("*event.Query")
		eventsRepo  *eventContextMocks.Repository
		eventTSRepo *eventTSMocks.Repository
	)
	BeforeEach(func() {
		eventsRepo = eventContextMocks.NewRepository(suiteTestingT)
		eventTSRepo = eventTSMocks.NewRepository(suiteTestingT)
		svc = event.NewService(eventsRepo, eventTSRepo, nil)
	})

	Context("Service.Find", func() {
		var (
			ctx    context.Context
			cancel context.CancelFunc
		)
		BeforeEach(func() {
			ctx, cancel = context.WithCancel(context.Background())
		})
		AfterEach(func() {
			cancel()
		})
		When("when the event.Repository.Find passes an error", func() {
			It("should fail", func() {
				testErr := errors.New("test error")
				eventsRepo.On("Find", mockCtx, mockQuery).Return(nil, testErr).Once()
				_, err := svc.Find(ctx, &eventContext.Query{})
				Ω(err).Should(HaveOccurred())
			})
		})
		When("when eventts.Repository.Get passes an error", func() {
			It("should fail", func() {
				eventsRepo.On("Find", mockCtx, mockQuery).Return(&eventContext.Page{
					Results: []*eventContext.Event{
						{
							ID: "id1",
							Occurrences: []*eventContext.Occurrence{
								{
									ID:        "id1:1",
									StartTime: 0,
									EndTime:   10000,
									Status:    eventContext.StatusClosed,
								},
							},
						},
					},
				}, nil).Once()
				eventTSRepo.On("Get", mockCtx, mock.Anything).
					Return(nil, errors.New("eventTS error")).Once()

				_, err := svc.Find(ctx, &eventContext.Query{
					Tenant: "acme",
					TimeRange: eventContext.TimeRange{
						Start: 0,
						End:   10000,
					},
					Fields: []string{"metadata"},
				})
				Ω(err).Should(SatisfyAll(
					HaveOccurred(),
					MatchError("eventTS error"),
				))
			})
		})
		When("results are found in both repositories", func() {
			It("should find events", func() {
				eventsRepo.On("Find", mockCtx, mockQuery).Return(&eventContext.Page{
					Results: []*eventContext.Event{
						{
							ID: "event1",
							Occurrences: []*eventContext.Occurrence{
								{
									ID:        "event1:1",
									EventID:   "event1",
									Tenant:    "acme",
									StartTime: 0,
									EndTime:   10000,
									Status:    eventContext.StatusClosed,
								},
							},
						},
					},
				}, nil).Once()

				eventTSRepo.On("Get", mockCtx, mock.Anything).
					Return([]*eventts.Occurrence{{
						ID:      "event1:1",
						EventID: "event1",
						Metadata: map[string][]interface{}{
							"k1": {"v1"},
						},
					}}, nil).Once()
				resp, err := svc.Find(ctx, &eventContext.Query{
					Tenant: "acme",
					TimeRange: eventContext.TimeRange{
						Start: 0,
						End:   10000,
					},
					Fields: []string{"metadata"},
				})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(resp).ShouldNot(BeNil())
			})
		})
	})
	Context("Service.Get", func() {
		var (
			ctx    context.Context
			cancel context.CancelFunc
		)
		BeforeEach(func() {
			ctx, cancel = context.WithCancel(context.Background())
		})
		AfterEach(func() {
			cancel()
		})
		When("when the event.Repository.Get passes an error", func() {
			It("should fail", func() {
				testErr := errors.New("test error")
				eventsRepo.On("Get", mockCtx, mock.Anything).Return(nil, testErr).Once()
				_, err := svc.Get(ctx, &eventContext.GetRequest{})
				Ω(err).Should(HaveOccurred())
			})
		})
		When("when eventts.Repository.Get passes an error", func() {
			It("should fail", func() {
				eventsRepo.On("Get", mockCtx, mock.Anything).Return(
					[]*eventContext.Event{{
						ID: "id1",
						Occurrences: []*eventContext.Occurrence{
							{
								ID:        "id1:1",
								StartTime: 0,
								EndTime:   10000,
								Status:    eventContext.StatusClosed,
							},
						},
					}}, nil).Once()
				eventTSRepo.On("Get", mockCtx, mock.Anything).
					Return(nil, errors.New("eventTS error")).Once()

				_, err := svc.Get(ctx, &eventContext.GetRequest{
					Tenant: "acme",
					ByOccurrenceIDs: struct{ IDs []string }{
						IDs: []string{"id1:1"},
					},
				})
				Ω(err).Should(SatisfyAll(
					HaveOccurred(),
					MatchError("eventTS error"),
				))
			})
		})
		When("results are found in both repositories", func() {
			It("should find events", func() {
				eventsRepo.On("Get", mockCtx, mock.Anything).Return(
					[]*eventContext.Event{{
						ID: "event1",
						Occurrences: []*eventContext.Occurrence{
							{
								ID:        "event1:1",
								StartTime: 0,
								EndTime:   10000,
								Status:    eventContext.StatusClosed,
								Metadata: map[string][]interface{}{
									"k0": {"v01", "v02"},
								},
							},
						},
					}}, nil).Once()

				eventTSRepo.On("Get", mockCtx, mock.Anything).
					Return([]*eventts.Occurrence{{
						ID:      "event1:1",
						EventID: "event1",
						Metadata: map[string][]interface{}{
							"k1": {"v1"},
						},
					}}, nil).Once()
				resp, err := svc.Get(ctx, &eventContext.GetRequest{
					Tenant: "acme",
					ByOccurrenceIDs: struct{ IDs []string }{
						IDs: []string{"id1:1"},
					},
				})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(resp).ShouldNot(BeNil())
			})
		})
	})
})
