//go:build integration

package mongo_test

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	redisCursors "github.com/zenoss/event-management-service/pkg/adapters/datasources/cursors/redis"
	"github.com/zenoss/event-management-service/pkg/adapters/datasources/eventcontext/mongo"
	"github.com/zenoss/zenkit/v5"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"
	"github.com/zenoss/event-management-service/pkg/adapters/framework/mongodb"
	"github.com/zenoss/event-management-service/pkg/models/event"
)

var _ = Describe("MongoDB Integration Test", func() {
	var (
		testCtx    context.Context
		testCancel context.CancelFunc
	)
	BeforeEach(func() {
		zenkit.InitConfigForTests(GinkgoT(), "event-management-service")
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
		addresses := make(map[string]string)
		for i, addr := range viper.GetStringSlice(zenkit.GCMemstoreAddressConfig) {
			addresses[fmt.Sprintf("%d", i+1)] = addr
		}
		Expect(addresses).ToNot(BeEmpty())
		redisClient := redis.NewRing(&redis.RingOptions{
			Addrs:       addresses,
			DialTimeout: 2 * time.Second,
		})
		cursorAdapter := redisCursors.NewAdapter(redisClient)
		db, err := mongodb.NewMongoDatabaseConnection(testCtx, cfg)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(db).ShouldNot(BeNil())
		adapter, err := mongo.NewAdapter(testCtx, cfg, db, cursorAdapter)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(adapter).ShouldNot(BeNil())
	})
	Context("mongodb.Adapter.Find", func() {
		var (
			cursorAdapter event.CursorRepository
			adapter       *mongo.Adapter
			cfg           mongodb.Config
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

			addresses := make(map[string]string)
			for i, addr := range viper.GetStringSlice(zenkit.GCMemstoreAddressConfig) {
				addresses[fmt.Sprintf("%d", i+1)] = addr
			}
			Expect(addresses).ToNot(BeEmpty())
			redisClient := redis.NewRing(&redis.RingOptions{
				Addrs:       addresses,
				DialTimeout: 2 * time.Second,
			})
			cursorAdapter = redisCursors.NewAdapter(redisClient)

			db, err := mongodb.NewMongoDatabaseConnection(testCtx, cfg)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(db).ShouldNot(BeNil())
			adapter, err = mongo.NewAdapter(testCtx, cfg, db, cursorAdapter)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(adapter).ShouldNot(BeNil())
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
			resp, err := adapter.Find(testCtx, input)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})
	})
})
