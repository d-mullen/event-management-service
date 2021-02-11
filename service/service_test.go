package service_test

import (
	"context"

	. "github.com/zenoss/event-management-service/service"
	// proto "github.com/zenoss/zing-proto/v11/go/event-management-service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Service", func() {

	var (
		ctx context.Context
		// svc proto.Event-Management-ServiceService
	)

	BeforeEach(func() {
		ctx = context.Background()
		// svc = NewEvent-Management-ServiceServiceServer()
	})

	It("should do something event-management-serviceish", func() {
		Ω(true).Should(BeTrue())
	})

})
