//go:build integration

package redis_test

import (
	"context"
	"fmt"
	"sync"
	"time"

	redis2 "github.com/go-redis/redis/v8"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/spf13/viper"
	"github.com/zenoss/event-management-service/pkg/adapters/scopes/activeents"
	"github.com/zenoss/event-management-service/pkg/adapters/scopes/activeents/redis"
	"github.com/zenoss/event-management-service/pkg/models/scopes"
	"github.com/zenoss/zenkit/v5"
)

var _ = Describe("Redis ActiveEntity Adapter Integration Test", func() {
	var (
		once    sync.Once
		adapter scopes.ActiveEntityRepository
		client  *redis2.Ring

		adapterFactory = func() scopes.ActiveEntityRepository {
			once.Do(func() {
				adapter = redis.NewAdapter(24*time.Hour, 24*time.Hour, client)
			})
			return adapter
		}
	)
	Context("...", Ordered, func() {
		BeforeAll(func() {
			client = initRedis()
			resp := client.FlushAll(context.Background())
			gomega.Expect(resp.Err()).ToNot(gomega.HaveOccurred())
		})
		AfterAll(func() {
			_ = client.Close()
		})
		activeents.RunActiveEntitySpecs(1, adapterFactory)
	})
})

func initRedis() *redis2.Ring {
	zenkit.InitConfigForTests(GinkgoT(), "event-management-service")
	addresses := make(map[string]string)
	for i, addr := range viper.GetStringSlice(zenkit.GCMemstoreAddressConfig) {
		addresses[fmt.Sprintf("%d", i+1)] = addr
	}
	return redis2.NewRing(&redis2.RingOptions{
		Addrs:       addresses,
		DialTimeout: 2 * time.Second,
	})
}
