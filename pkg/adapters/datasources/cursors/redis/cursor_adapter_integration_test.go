//go:build integration

package redis_test

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
	redisCursors "github.com/zenoss/event-management-service/pkg/adapters/datasources/cursors/redis"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/zenkit/v5"
)

func initRedis() *redis.Ring {
	zenkit.InitConfigForTests(GinkgoT(), "event-management-service")
	addresses := make(map[string]string)
	for i, addr := range viper.GetStringSlice(zenkit.GCMemstoreAddressConfig) {
		addresses[fmt.Sprintf("%d", i+1)] = addr
	}
	return redis.NewRing(&redis.RingOptions{
		Addrs:       addresses,
		DialTimeout: 2 * time.Second,
	})
}

var _ = Describe("Redis Cursor Adapter Integration Tests", Ordered, func() {
	var (
		testCtx context.Context
		cancel  context.CancelFunc
		adapter event.CursorRepository
		client  *redis.Ring
	)
	BeforeAll(func() {
		client = initRedis()
		resp := client.FlushAll(context.Background())
		Expect(resp.Err()).ToNot(HaveOccurred())
	})
	AfterAll(func() {
		_ = client.Close()
	})
	BeforeEach(func() {
		testCtx, cancel = context.WithCancel(context.Background())
		adapter = redisCursors.NewAdapter(client)
	})
	AfterEach(func() {
		cancel()
	})
	It("should create and retrieve query cursors", func() {

		req := &event.Cursor{
			ID: "c1",
			Query: event.Query{
				Tenant: "acme",
				TimeRange: event.TimeRange{
					Start: 0,
					End:   10000,
				},
				Severities: []event.Severity{event.SeverityCritical},
				Statuses:   []event.Status{event.StatusOpen},
			},
		}

		cursor, err := adapter.New(testCtx, req)
		Expect(err).ToNot(HaveOccurred())
		Expect(cursor).ToNot(BeEmpty())

		actual, err := adapter.Get(testCtx, cursor)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(BeEquivalentTo(req))

	})
})
