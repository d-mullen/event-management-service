//go:build integration
// +build integration

package mongodb_test

import (
	"context"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
	"github.com/zenoss/event-management-service/pkg/domain/event"
)

var _ = Describe("MongoDB Integration Test", func() {
	var (
		testCtx    context.Context
		testCancel context.CancelFunc
	)
	BeforeEach(func() {
		testCtx, testCancel = context.WithCancel(context.Background())
		log := logrus.NewEntry(logrus.StandardLogger())
		log.Logger.SetFormatter(&logrus.JSONFormatter{})
		log.Logger.SetLevel(logrus.DebugLevel)
		testCtx = ctxlogrus.ToContext(testCtx, log)
	})
	AfterEach(func() {
		testCancel()
	})
	It("should connect to the configured MongoDB deployment", func() {
		cfg := mongodb.Config{
			Address:    "mongodb",
			Username:   "zing-user",
			Password:   ".H1ddenPassw0rd.",
			DefaultTTL: 90 * 24 * time.Hour,
			DBName:     "event-context-svc",
		}

		db, err := mongodb.NewAdapter(testCtx, cfg)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(db).ShouldNot(BeNil())
	})
	Context("mongodb.Adapter.Find", func() {
		var (
			db  *mongodb.Adapter
			cfg mongodb.Config
		)
		BeforeEach(func() {
			var err error
			cfg = mongodb.Config{
				Address:    "mongodb",
				Username:   "zing-user",
				Password:   ".H1ddenPassw0rd.",
				DefaultTTL: 90 * 24 * time.Hour,
				DBName:     "event-context-svc",
			}

			db, err = mongodb.NewAdapter(testCtx, cfg)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(db).ShouldNot(BeNil())
		})
		It("find documents", func() {
			input := &event.Query{
				Tenant: "acme",
				TimeRange: event.TimeRange{
					Start: time.Now().Add(-1 * time.Hour).UnixMilli(),
					End:   time.Now().UnixMilli(),
				},
				Statuses:   []event.Status{event.StatusOpen},
				Severities: []event.Severity{event.SeverityError, event.SeverityCritical},
			}
			resp, err := db.Find(testCtx, input)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})
	})
})
