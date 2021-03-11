package service_test

import (
	"context"
	//	"fmt"
	//	"math/rand"
	//	"sync"
	//	"time"

	//	"google.golang.org/grpc/codes"
	//	"google.golang.org/grpc/status"

	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	////wrapperspb "github.com/golang/protobuf/types/known/wrapperspb"
	//	"github.com/onsi/gomega/gstruct"
	. "github.com/zenoss/event-management-service/service"
	zenkitMocks "github.com/zenoss/zenkit/v5/mocks"
	ecproto "github.com/zenoss/zing-proto/v11/go/cloud/event_context"
	//	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
	//"github.com/zenoss/zingo/v4/protobufutils"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo"
	//	. "github.com/onsi/ginkgo/extensions/table"
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
		annotation1 proto.Annotation
		annotation2 proto.Annotation
	)

	BeforeEach(func() {
		log := zenkit.Logger("event-management-service")
		zenkit.InitConfig("event-management-service")
		ctx, cancel = context.WithCancel(ctxlogrus.ToContext(context.Background(), log))
		eStatus = proto.EMEventStatus{
			Acknowledge:   &wrappers.BoolValue{Value: true},
			StatusWrapper: &proto.EMEventStatus_Wrapper{Status: proto.EMStatus_EM_STATUS_OPEN},
		}
		annotation1 = proto.Annotation{
			AnnotationId: "",
			Annotation:   "this is a new note for testing",
		}
		annotation2 = proto.Annotation{
			AnnotationId: "noteId1",
			Annotation:   "this is a new note for testing",
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
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("set status", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: ""}, nil).Once()

			_, err := svc.SetStatus(ctx, &proto.EventStatusRequest{
				StatusList: map[string]*proto.EMEventStatus{
					"eventId1": &eStatus,
				},
			})
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("annotate-add", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: "newNoteId1"}, nil).Once()

			_, err := svc.Annotate(ctx, &proto.EventAnnotationRequest{
				AnnotationList: map[string]*proto.Annotation{
					"eventId1": &annotation1,
				},
			})
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("annotate-edit", func() {
			clientMock.On("UpdateEvent", mock.Anything, mock.AnythingOfType("*event_context.UpdateEventRequest")).Return(
				&ecproto.UpdateEventResponse{Status: true, NoteId: "noteId1"}, nil).Once()

			_, err := svc.Annotate(ctx, &proto.EventAnnotationRequest{
				AnnotationList: map[string]*proto.Annotation{
					"eventId1": &annotation2,
				},
			})
			Ω(err).ShouldNot(HaveOccurred())
		})

	})
})
