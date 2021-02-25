package service_test

import (
	"context"
	//	"fmt"
	//	"math/rand"
	//	"sync"
	//	"time"

	//	"google.golang.org/grpc/codes"
	//	"google.golang.org/grpc/status"

	//	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	//	"github.com/onsi/gomega/gstruct"
	. "github.com/zenoss/event-management-service/service"
	zenkitMocks "github.com/zenoss/zenkit/v5/mocks"
	//	"github.com/zenoss/zing-proto/v11/go/cloud/common"
	proto "github.com/zenoss/zing-proto/v11/go/cloud/event_management"
	//"github.com/zenoss/zingo/v4/protobufutils"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo"
	//	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/zenoss/zenkit/v5"
	//	"google.golang.org/grpc/metadata"
)

var _ = Describe("Management Service", func() {
	var (
		ctx     context.Context
		cancel  context.CancelFunc
		eStatus proto.EMEventStatus
	)

	BeforeEach(func() {
		log := zenkit.Logger("event-management-service")
		zenkit.InitConfig("event-management-service")
		ctx, cancel = context.WithCancel(ctxlogrus.ToContext(context.Background(), log))
		eStatus = proto.EMEventStatus{
			Acknowledge: true,
			Status:      proto.EMStatus_EM_STATUS_OPEN,
		}
	})

	AfterEach(func() {
		cancel()
	})

	Context("EventManagementStatus", func() {
		var (
			svc proto.EventManagementServer
		)

		BeforeEach(func() {
			var err error
			ctx = zenkit.WithTenantIdentity(ctx, zenkitMocks.NewMockTenantIdentity("acme", "user@acme.com"))
			svc, err = NewEventManagementService(ctx)
			Ω(err).ShouldNot(HaveOccurred())
		})

		It("set status", func() {
			_, err := svc.SetStatus(ctx, &proto.EventStatusRequest{
				Tenant: "acme",
				StatusList: map[string]*proto.EMEventStatus{
					"eventId1": &eStatus,
				},
			})
			Ω(err).ShouldNot(HaveOccurred())
		})

	})
})
