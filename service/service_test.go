package service_test

import (
	"context"
	//	"fmt"
	//	"math/rand"
	//	"sync"
	//	"time"

	//	"google.golang.org/grpc/codes"
	//	"google.golang.org/grpc/status"
	"github.com/pkg/errors"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	////wrapperspb "github.com/golang/protobuf/types/known/wrapperspb"
	//	"github.com/onsi/gomega/gstruct"
	. "github.com/zenoss/event-management-service/service"
	zenkitMocks "github.com/zenoss/zenkit/v5/mocks"
	ecproto "github.com/zenoss/zing-proto/v11/go/cloud/event_context"

	//	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
	eproto "github.com/zenoss/zing-proto/v11/go/event"

	//"github.com/zenoss/zingo/v4/protobufutils"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo/v2"

	//	. "github.com/onsi/ginkgo/v2/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/zenoss/zenkit/v5"

	//	"google.golang.org/grpc/metadata"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Management Service", func() {
	var (
		ctx         context.Context
		cancel      context.CancelFunc
		eStatus     proto.EMEventStatus
		eStatus2    proto.EMEventStatus
		eStatus3    proto.EMEventStatus
		eStatus4    proto.EMEventStatus
		eStatus5    proto.EMEventStatus
		annotation1 proto.Annotation
		annotation2 proto.Annotation
		annotation3 proto.Annotation
	)

	BeforeEach(func() {
		log := zenkit.Logger("event-management-service")
		zenkit.InitConfig("event-management-service")
		ctx, cancel = context.WithCancel(ctxlogrus.ToContext(context.Background(), log))
		eStatus = proto.EMEventStatus{
			EventId:       "eventId1",
			OccurrenceId:  "occId1",
			Acknowledged:  &wrappers.BoolValue{Value: true},
			StatusWrapper: &proto.EMEventStatus_Wrapper{Status: eproto.Status_STATUS_OPEN},
		}
		eStatus2 = proto.EMEventStatus{
			EventId:       "eventId1",
			OccurrenceId:  "occId1",
			Acknowledged:  &wrappers.BoolValue{Value: true},
			StatusWrapper: nil,
		}
		eStatus3 = proto.EMEventStatus{
			EventId:       "",
			OccurrenceId:  "",
			Acknowledged:  &wrappers.BoolValue{Value: true},
			StatusWrapper: nil,
		}
		eStatus4 = proto.EMEventStatus{
			EventId:       "eventId1",
			OccurrenceId:  "occId1",
			Acknowledged:  nil,
			StatusWrapper: &proto.EMEventStatus_Wrapper{Status: eproto.Status_STATUS_CLOSED},
		}
		eStatus5 = proto.EMEventStatus{
			EventId:       "eventId1",
			OccurrenceId:  "occId1",
			Acknowledged:  nil,
			StatusWrapper: nil,
		}
		annotation1 = proto.Annotation{
			EventId:      "eventId1",
			OccurrenceId: "occId1",
			AnnotationId: "",
			Annotation:   "this is a new note for testing",
		}
		annotation2 = proto.Annotation{
			EventId:      "eventId1",
			OccurrenceId: "occId1",
			AnnotationId: "noteId1",
			Annotation:   "this is a edited note for testing",
		}
		annotation3 = proto.Annotation{
			EventId:      "",
			OccurrenceId: "occId1",
			AnnotationId: "noteId0",
			Annotation:   "this is a note for testing with no event id",
		}
	})

	AfterEach(func() {
		cancel()
	})

	Context("EventManagementStatus", func() {
		var (
			svc        proto.EventManagementServer
			clientMock *ecproto.MockEventContextIngestClient
		)

		BeforeEach(func() {
			var err error
			ctx = zenkit.WithTenantIdentity(ctx, zenkitMocks.NewMockTenantIdentity("acme", "user@acme.com"))
			clientMock = &ecproto.MockEventContextIngestClient{}
			svc, err = NewEventManagementServiceFromParts(clientMock)
			??(err).ShouldNot(HaveOccurred())
		})

		It("set status - all", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: ""}, nil).Twice()

			resp, err := svc.SetStatus(ctx, &proto.EventStatusRequest{
				Statuses: []*proto.EMEventStatus{
					&eStatus,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.StatusResponses[0].Success).Should(BeTrue())
		})

		It("set status -ack only", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: ""}, nil).Twice()

			resp, err := svc.SetStatus(ctx, &proto.EventStatusRequest{
				Statuses: []*proto.EMEventStatus{
					&eStatus2,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
		})

		It("set status - status only", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: ""}, nil).Twice()

			resp, err := svc.SetStatus(ctx, &proto.EventStatusRequest{
				Statuses: []*proto.EMEventStatus{
					&eStatus4,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.StatusResponses[0].Success).Should(BeTrue())
		})

		It("set status - fail", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: false, NoteId: ""}, errors.New("failing just for test")).Twice()

			resp, err := svc.SetStatus(ctx, &proto.EventStatusRequest{
				Statuses: []*proto.EMEventStatus{
					&eStatus,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.StatusResponses[0].Success).Should(BeFalse())
		})

		It("set status - err", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: false, NoteId: ""}, errors.New("just an error")).Twice()

			resp, err := svc.SetStatus(ctx, &proto.EventStatusRequest{
				Statuses: []*proto.EMEventStatus{
					&eStatus3,
				},
			})
			??(resp).ShouldNot(BeNil())
			Expect(err).ShouldNot(HaveOccurred())
			??(resp.StatusResponses[0].Success).Should(BeFalse())
		})

		It("set status -nil request", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: ""}, nil).Twice()

			resp, err := svc.SetStatus(ctx, nil)
			Expect(err).Should(HaveOccurred())
			??(resp).Should(BeNil())
		})

		It("set status -nothing set", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: ""}, nil).Twice()

			resp, err := svc.SetStatus(ctx, &proto.EventStatusRequest{
				Statuses: []*proto.EMEventStatus{
					&eStatus5,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.StatusResponses[0].Success).Should(BeFalse())
		})

		It("annotate-add", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: "newNoteId1"}, nil).Twice()

			resp, err := svc.Annotate(ctx, &proto.EventAnnotationRequest{
				Annotations: []*proto.Annotation{
					&annotation1,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.AnnotationResponses[0].Success).Should(BeTrue())
		})

		It("annotate-del", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: "noteId1"}, nil).Twice()

			resp, err := svc.DeleteAnnotations(ctx, &proto.EventAnnotationRequest{
				Annotations: []*proto.Annotation{
					&annotation2,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.AnnotationResponses[0].Success).Should(BeTrue())
		})

		It("annotate-add -fail", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: false, NoteId: "newNoteId1"}, errors.New("just a bogus error")).Twice()

			resp, err := svc.Annotate(ctx, &proto.EventAnnotationRequest{
				Annotations: []*proto.Annotation{
					&annotation1,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.AnnotationResponses[0].Success).Should(BeFalse())
		})

		It("annotate-edit", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: "noteId1"}, nil).Twice()

			resp, err := svc.Annotate(ctx, &proto.EventAnnotationRequest{
				Annotations: []*proto.Annotation{
					&annotation2,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.AnnotationResponses[0].Success).Should(BeTrue())
		})

		It("annotate- no eventId", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: false, NoteId: ""}, errors.New("just another error")).Twice()

			resp, err := svc.Annotate(ctx, &proto.EventAnnotationRequest{
				Annotations: []*proto.Annotation{
					&annotation3,
				},
			})
			Expect(err).ShouldNot(HaveOccurred())
			??(resp).ShouldNot(BeNil())
			??(resp.AnnotationResponses[0].Success).Should(BeFalse())
		})

		It("annotate- nil request", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: false, NoteId: ""}, errors.New("just another error")).Twice()

			resp, err := svc.Annotate(ctx, nil)
			Expect(err).Should(HaveOccurred())
			??(resp).Should(BeNil())
		})
	})
})
