package event_test

import (
	"context"

	"github.com/zenoss/event-management-service/pkg/adapters/scopes/activeents"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"github.com/zenoss/event-management-service/pkg/application/event"
	eventContext "github.com/zenoss/event-management-service/pkg/models/event"
	eventContextMocks "github.com/zenoss/event-management-service/pkg/models/event/mocks"
	"github.com/zenoss/event-management-service/pkg/models/eventts"
	eventTSMocks "github.com/zenoss/event-management-service/pkg/models/eventts/mocks"
)

type fakeEntityScopeProvider struct{}

func (fesp *fakeEntityScopeProvider) GetEntityIDs(ctx context.Context, scopeCursor string) ([]string, error) {
	return []string{"1", "2"}, nil
}

var _ = Describe("eventquery.Service", func() {
	var (
		svc         event.Service
		mockCtx     = mock.Anything
		mockQuery   = mock.AnythingOfType("*event.Query")
		eventsRepo  *eventContextMocks.Repository
		eventTSRepo *eventTSMocks.Repository
		entityScope = &fakeEntityScopeProvider{}
	)

	BeforeEach(func() {
		eventsRepo = eventContextMocks.NewRepository(suiteTestingT)
		eventTSRepo = eventTSMocks.NewRepository(suiteTestingT)
		svc = event.NewService(eventsRepo, eventTSRepo, entityScope, activeents.NewInMemoryActiveEntityAdapter(1024, activeents.DefaultBucketSize))
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
				eventsRepo.On("Find", mockCtx, mockQuery).Return(nil, testErr)
				_, err := svc.Find(ctx, &eventContext.Query{
					Tenant: "acme",
					TimeRange: eventContext.TimeRange{
						Start: 0,
						End:   10000,
					},
				})
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
									Severity:  eventContext.SeverityDefault,
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
					Severities: []eventContext.Severity{eventContext.SeverityDefault},
					Statuses:   []eventContext.Status{eventContext.StatusClosed},
					Fields:     []string{"metadata"},
				})
				Ω(err).Should(SatisfyAll(
					HaveOccurred(),
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
									Severity:  eventContext.SeverityDefault,
								},
							},
						},
					},
				}, nil).Once()

				eventTSRepo.On("Get", mockCtx, mock.Anything).
					Return([]*eventts.Occurrence{{
						ID:      "event1:1",
						EventID: "event1",
						Metadata: map[string][]any{
							"k1": {"v1"},
						},
					}}, nil).Once()
				resp, err := svc.Find(ctx, &eventContext.Query{
					Tenant: "acme",
					TimeRange: eventContext.TimeRange{
						Start: 0,
						End:   10000,
					},
					Severities: []eventContext.Severity{eventContext.SeverityDefault},
					Statuses:   []eventContext.Status{eventContext.StatusClosed},
					Fields:     []string{"metadata"},
				})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(resp).ShouldNot(BeNil())
			})
		})
		When("results are found in repositories with filters ", func() {
			It("should find events", func() {
				eventsRepo.On("Find", mockCtx, mockQuery).Return(&eventContext.Page{
					Results: []*eventContext.Event{
						{
							ID: "event2",
							Occurrences: []*eventContext.Occurrence{
								{
									ID:        "event2:1",
									EventID:   "event2",
									Tenant:    "acme",
									StartTime: 0,
									EndTime:   10000,
									Status:    eventContext.StatusClosed,
									Severity:  eventContext.SeverityDefault,
								},
							},
						},
					},
				}, nil).Once()

				eventTSRepo.On("Get", mockCtx, mock.Anything).
					Return([]*eventts.Occurrence{{
						ID:      "event2:1",
						EventID: "event2",
						Metadata: map[string][]any{
							"k1": {"v1"},
						},
					}}, nil).Once()
				resp, err := svc.Find(ctx, &eventContext.Query{
					Tenant: "acme",
					TimeRange: eventContext.TimeRange{
						Start: 0,
						End:   10000,
					},
					Severities: []eventContext.Severity{eventContext.SeverityDefault},
					Statuses:   []eventContext.Status{eventContext.StatusClosed},
					Fields:     []string{"summary", "acknowledged", "body", "UNSupportedField"},
					Filter: &eventContext.Filter{Op: eventContext.FilterOpAnd,
						Field: "",
						Value: []*eventContext.Filter{
							&eventContext.Filter{Op: eventContext.FilterOpEqualTo, Field: "tenant", Value: "Acme"},
							&eventContext.Filter{Op: eventContext.FilterOpNotEqualTo, Field: "status", Value: "1"},
						}},
				})
				Ω(err).ShouldNot(HaveOccurred())
				Ω(resp).ShouldNot(BeNil())
			})
		})

		When("filters are passed to event-ts ", func() {
			It("should get filters in events", func() {
				eventsRepo.On("Find", mockCtx, mockQuery).Return(&eventContext.Page{
					Results: []*eventContext.Event{
						{
							ID: "event3",
							Occurrences: []*eventContext.Occurrence{
								{
									ID:        "event3:1",
									EventID:   "event3",
									Tenant:    "acme",
									StartTime: 0,
									EndTime:   10000,
									Status:    eventContext.StatusClosed,
									Severity:  eventContext.SeverityDefault,
								},
							},
						},
					},
				}, nil).Once()

				eventTSRepo.On("Get", mockCtx, mock.Anything).
					Run(func(args mock.Arguments) {
						reqAny := args.Get(1)
						req, ok := reqAny.(*eventts.GetRequest)
						Expect(ok).To(BeTrue())
						Expect(req.Filters).Should(BeEquivalentTo([](*eventts.Filter){
							{
								Operation: eventts.Operation_OP_IN,
								Field:     "_zv_severity",
								Values:    []interface{}{0},
							},
							{
								Operation: eventts.Operation_OP_IN,
								Field:     "_zv_status",
								Values:    []interface{}{3},
							}}))
						Expect(req.ResultFields).Should(BeEquivalentTo([]string{"summary", "acknowledged", "body", "newField"}))
					}).
					Return([]*eventts.Occurrence{{
						ID:      "event2:1",
						EventID: "event2",
						Metadata: map[string][]any{
							"k1": {"v1"}, // actually here should be fields as above
						},
					}}, nil).Once()
				resp, err := svc.Find(ctx, &eventContext.Query{
					Tenant: "acme",
					TimeRange: eventContext.TimeRange{
						Start: 0,
						End:   10000,
					},
					Severities: []eventContext.Severity{eventContext.SeverityDefault},
					Statuses:   []eventContext.Status{eventContext.StatusClosed},
					Fields:     []string{"summary", "acknowledged", "body", "newField"},
					Filter: &eventContext.Filter{Op: eventContext.FilterOpAnd,
						Field: "",
						Value: []*eventContext.Filter{
							&eventContext.Filter{Op: eventContext.FilterOpEqualTo, Field: "tenant", Value: "Acme"},
							&eventContext.Filter{Op: eventContext.FilterOpNotEqualTo, Field: "status", Value: "1"},
						}},
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
								Metadata: map[string][]any{
									"k0": {"v01", "v02"},
								},
							},
						},
					}}, nil).Once()

				eventTSRepo.On("Get", mockCtx, mock.Anything).
					Return([]*eventts.Occurrence{{
						ID:      "event1:1",
						EventID: "event1",
						Metadata: map[string][]any{
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
