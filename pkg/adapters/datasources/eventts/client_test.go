package eventts_test

import (
	"context"
	"fmt"
	"io"

	eventts2 "github.com/zenoss/event-management-service/pkg/adapters/datasources/eventts"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"github.com/zenoss/event-management-service/pkg/models/eventts"
	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	eventtsPb "github.com/zenoss/zing-proto/v11/go/cloud/eventts"
	"github.com/zenoss/zingo/v4/protobufutils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = DescribeTable(
	"EventTSRequestToProto Table-driven Tests",
	func(req *eventts.GetRequest, expected *eventtsPb.EventTSRequest, shouldFail bool) {
		actual, err := eventts2.EventTSRequestToProto(req)
		if shouldFail {
			Ω(err).Should(HaveOccurred())
			Ω(actual).Should(BeNil())
		} else {
			Ω(err).ShouldNot(HaveOccurred())
			Ω(actual).ShouldNot(BeNil())
		}
	},
	func(req *eventts.GetRequest, expected *eventtsPb.EventTSRequest, shouldFail bool) string {
		failStr := "should not"
		if shouldFail {
			failStr = "should"
		}
		return fmt.Sprintf("EventTSRequestToProto(%#v) = %#v %s fail\n", req, expected, failStr)
	},
	Entry(nil, nil, nil, true),
	Entry(nil, &eventts.GetRequest{}, nil, true),
	Entry(
		nil,
		&eventts.GetRequest{
			EventTimeseriesInput: eventts.EventTimeseriesInput{
				TimeRange: eventts.TimeRange{Start: 0, End: 10000},
				ByEventIDs: struct {
					IDs []string
				}{IDs: []string{"event1", "event2"}},
				Latest:  0,
				Fields:  []string{},
				Filters: []*eventts.Filter{},
			},
		},
		&eventtsPb.EventTSRequest{
			EventIds: []string{"event1", "event2"},
			TimeRange: &common.TimeRange{
				Start: 0,
				End:   10000,
			},
		},
		false),
	Entry(
		nil,
		&eventts.GetRequest{
			EventTimeseriesInput: eventts.EventTimeseriesInput{
				ByOccurrences: struct {
					ShouldApplyIntervals bool
					OccurrenceMap        map[string][]*eventts.OccurrenceInput
				}{
					ShouldApplyIntervals: false,
					OccurrenceMap: map[string][]*eventts.OccurrenceInput{
						"event1": {{
							ID:      "event1:1",
							EventID: "event1",
							TimeRange: eventts.TimeRange{
								Start: 0,
								End:   5000,
							},
							IsActive: false,
						}},
						"event2": {{
							ID:      "event2:1",
							EventID: "event2",
							TimeRange: eventts.TimeRange{
								Start: 2000,
								End:   9500,
							},
							IsActive: false,
						}},
					},
				},
			}},
		&eventtsPb.EventTSRequest{
			OccurrenceMap: map[string]*eventtsPb.EventTSOccurrenceCollection{
				"event1": {
					Occurrences: []*eventtsPb.EventTSOccurrence{
						{
							Id:       "event1:1",
							IsActive: false,
							TimeRange: &common.TimeRange{
								Start: 0,
								End:   5000,
							},
						},
					},
				},
				"event2": {
					Occurrences: []*eventtsPb.EventTSOccurrence{
						{
							Id:       "event2:1",
							IsActive: false,
							TimeRange: &common.TimeRange{
								Start: 2000,
								End:   9500,
							},
						},
					},
				},
			},
			TimeRange: &common.TimeRange{
				Start: 0,
				End:   9500,
			},
		},
		false),
)

// var _ = DescribeTable("EventTSSeriesToOccurrence Table-driven Tests", func(a, b any) {
// 	Ω(true).Should(BeTrue())
// })

var _ = Describe("EventTSService Adapter Unit-tests", func() {
	var (
		ctx               context.Context
		cancel            context.CancelFunc
		repo              eventts.Repository
		eventTSMockClient *eventtsPb.MockEventTSServiceClient
		// getEventStreamMock *eventtsPb.MockEventTSService_GetEventsStreamClient
		ewcStreamMock *eventtsPb.MockEventTSService_EventsWithCountsStreamClient
		mockCtx       = mock.AnythingOfType("*context.cancelCtx")
		// mockEventTSReq     = mock.AnythingOfType("*eventts.EventTSRequest")
		mockEWCReq = mock.AnythingOfType("*eventts.EventsWithCountsRequest")
	)
	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())
		eventTSMockClient = &eventtsPb.MockEventTSServiceClient{}
		// getEventStreamMock = &eventtsPb.MockEventTSService_GetEventsStreamClient{}
		ewcStreamMock = &eventtsPb.MockEventTSService_EventsWithCountsStreamClient{}
		repo = eventts2.NewAdapter(eventTSMockClient)
	})
	AfterEach(func() {
		cancel()
	})
	Context("Get", func() {
		It("should get event timeseries results result", func() {

			eventTSMockClient.On("EventsWithCountsStream", mockCtx, mockEWCReq).
				Return(ewcStreamMock, nil).Once()
			ewcStreamMock.On("Recv").
				Return(&eventtsPb.EventsWithCountsResponse{
					Results: []*eventtsPb.EventTSResult{
						{
							Series: []*eventtsPb.EventTSSeries{{
								EventId:      "event1",
								OccurrenceId: "event1:1",
								Values: []*eventtsPb.EventTSField{
									{
										Timestamp: 0,
										Data: protobufutils.MustToScalarArrayMap(map[string][]any{
											"k1": {"v1", "v2"},
											"k2": {1},
										}),
									},
									{
										Timestamp: 5,
										Data: protobufutils.MustToScalarArrayMap(map[string][]any{
											"k1": {"v3", "v4"},
											"k2": {10},
										}),
									},
								},
							}},
						},
					},
				}, nil).Once()
			ewcStreamMock.On("Recv").Return(nil, io.EOF).Once()
			results, err := repo.Get(ctx, &eventts.GetRequest{
				EventTimeseriesInput: eventts.EventTimeseriesInput{
					TimeRange: eventts.TimeRange{
						Start: 0,
						End:   10,
					},
					ByEventIDs: struct {
						IDs []string
					}{
						IDs: []string{"event1", "event2"},
					},
				},
			})
			Ω(err).ShouldNot(HaveOccurred())
			Ω(results).ShouldNot(BeNil())
		})
		When("event-ts-svc returns an error", func() {
			It("it should fail", func() {

				By("passing up an error when event-ts-svc fails to return stream")
				eventTSMockClient.On("EventsWithCountsStream", mockCtx, mockEWCReq).
					Return(nil, status.Error(codes.Unknown, "test error")).Once()

				result, err := repo.Get(ctx, &eventts.GetRequest{})
				Ω(err).Should(HaveOccurred())
				Ω(result).Should(BeNil())

				By("by passing up an error when an error occurs in the stream")
				eventTSMockClient.On("EventsWithCountsStream", mockCtx, mockEWCReq).
					Return(ewcStreamMock, nil).Once()
				ewcStreamMock.On("Recv").
					Return(&eventtsPb.EventsWithCountsResponse{
						Results: []*eventtsPb.EventTSResult{
							{
								Series: []*eventtsPb.EventTSSeries{{
									EventId:      "event1",
									OccurrenceId: "event1:1",
									Values: []*eventtsPb.EventTSField{
										{
											Timestamp: 6,
											Data: protobufutils.MustToScalarArrayMap(map[string][]any{
												"k1": {"v1", "v2"},
												"k2": {1},
											}),
										},
										{
											Timestamp: 5,
											Data: protobufutils.MustToScalarArrayMap(map[string][]any{
												"k1": {"v3", "v4"},
												"k2": {10},
											}),
										},
									},
								}},
							},
						},
					}, nil).Once()
				ewcStreamMock.On("Recv").Return(nil, status.Error(codes.Unknown, "stream error")).Once()
				result, err = repo.Get(ctx, &eventts.GetRequest{
					EventTimeseriesInput: eventts.EventTimeseriesInput{
						ByOccurrences: struct {
							ShouldApplyIntervals bool
							OccurrenceMap        map[string][]*eventts.OccurrenceInput
						}{
							OccurrenceMap: map[string][]*eventts.OccurrenceInput{
								"event1": {
									{
										ID:        "event1:1",
										EventID:   "event1",
										TimeRange: eventts.TimeRange{Start: 0, End: 10},
										IsActive:  false,
									},
								},
							},
						},
					},
				})
				Ω(err).Should(HaveOccurred())
				Ω(result).Should(BeNil())
			})
		})
	})
	Context("GetStream", func() {
		When("when the calls to event-ts-svc succeed", func() {
			It("should succeed", func() {
				eventTSMockClient.On("EventsWithCountsStream", mockCtx, mockEWCReq).
					Return(ewcStreamMock, nil).Once()
				ewcStreamMock.On("Recv").
					Return(&eventtsPb.EventsWithCountsResponse{
						Results: []*eventtsPb.EventTSResult{
							{
								Series: []*eventtsPb.EventTSSeries{{
									EventId:      "event1",
									OccurrenceId: "event1:1",
									Values: []*eventtsPb.EventTSField{
										{
											Timestamp: 6,
											Data: protobufutils.MustToScalarArrayMap(map[string][]any{
												"k1": {"v1", "v2"},
												"k2": {1},
											}),
										},
										{
											Timestamp: 5,
											Data: protobufutils.MustToScalarArrayMap(map[string][]any{
												"k1": {"v3", "v4"},
												"k2": {10},
											}),
										},
									},
								}},
							},
						},
					}, nil).Once()
				ewcStreamMock.On("Recv").Return(nil, io.EOF).Once()
				ch := repo.GetStream(ctx, &eventts.GetRequest{
					EventTimeseriesInput: eventts.EventTimeseriesInput{
						ByOccurrences: struct {
							ShouldApplyIntervals bool
							OccurrenceMap        map[string][]*eventts.OccurrenceInput
						}{
							OccurrenceMap: map[string][]*eventts.OccurrenceInput{
								"event1": {
									{
										ID:        "event1:1",
										EventID:   "event1",
										TimeRange: eventts.TimeRange{Start: 0, End: 10},
										IsActive:  false,
									},
								},
							},
						},
					},
				})
				resp := <-ch
				Ω(resp).ShouldNot(BeNil())
				Ω(resp.Result).ShouldNot(BeNil())
				Eventually(ch).Should(BeClosed())
			})
		})
		When("when the calls to event-ts-svc fail", func() {
			It("should fail", func() {
				By("passing up an error when event-ts-svc fails to return stream")
				eventTSMockClient.On("EventsWithCountsStream", mockCtx, mockEWCReq).
					Return(nil, status.Error(codes.Unknown, "test error")).Once()

				ch := repo.GetStream(ctx, &eventts.GetRequest{})
				resp := <-ch
				Ω(resp).ShouldNot(BeNil())
				Ω(resp.Result).Should(BeNil())
				Ω(resp.Err).Should(HaveOccurred())
				Ω(ch).Should(BeClosed())

			})
			It("should fail", func() {
				By("by passing up an error when an error occurs in the stream")
				eventTSMockClient.On("EventsWithCountsStream", mockCtx, mockEWCReq).
					Return(ewcStreamMock, nil).Once()
				ewcStreamMock.On("Recv").
					Return(&eventtsPb.EventsWithCountsResponse{
						Results: []*eventtsPb.EventTSResult{
							{
								Series: []*eventtsPb.EventTSSeries{{
									EventId:      "event1",
									OccurrenceId: "event1:1",
									Values: []*eventtsPb.EventTSField{
										{
											Timestamp: 6,
											Data: protobufutils.MustToScalarArrayMap(map[string][]any{
												"k1": {"v1", "v2"},
												"k2": {1},
											}),
										},
										{
											Timestamp: 5,
											Data: protobufutils.MustToScalarArrayMap(map[string][]any{
												"k1": {"v3", "v4"},
												"k2": {10},
											}),
										},
									},
								}},
							},
						},
					}, nil).Once()
				ewcStreamMock.On("Recv").Return(nil, status.Error(codes.Unknown, "stream error")).Once()
				ch2 := repo.GetStream(ctx, &eventts.GetRequest{
					EventTimeseriesInput: eventts.EventTimeseriesInput{
						ByOccurrences: struct {
							ShouldApplyIntervals bool
							OccurrenceMap        map[string][]*eventts.OccurrenceInput
						}{
							OccurrenceMap: map[string][]*eventts.OccurrenceInput{
								"event1": {
									{
										ID:        "event1:1",
										EventID:   "event1",
										TimeRange: eventts.TimeRange{Start: 0, End: 10},
										IsActive:  false,
									},
								},
							},
						},
					},
				})
				resp2 := <-ch2
				Ω(resp2).ShouldNot(BeNil())
				Ω(resp2.Result).ShouldNot(BeNil())
				Ω(resp2.Err).ShouldNot(HaveOccurred())
				resp2 = <-ch2
				Ω(resp2).ShouldNot(BeNil())
				Ω(resp2.Result).ShouldNot(BeNil())
				Ω(resp2.Err).ShouldNot(HaveOccurred())
				resp2 = <-ch2
				Ω(resp2.Result).Should(BeNil())
				Ω(resp2.Err).Should(HaveOccurred())
				Eventually(ch2).Should(BeClosed())
			})
		})
	})
})
