package redis_test

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	redisCursors "github.com/zenoss/event-management-service/pkg/adapters/datasources/cursors/redis"
	"github.com/zenoss/event-management-service/pkg/adapters/datasources/cursors/redis/mocks"
	"github.com/zenoss/event-management-service/pkg/models/event"
)

var _ = Describe("Redis Cursor Adapter Tests", func() {
	var (
		testCtx context.Context
		cancel  context.CancelFunc
		adapter event.CursorRepository
		client  *mocks.RedisCommander
	)
	BeforeEach(func() {
		testCtx, cancel = context.WithCancel(context.Background())
		client = mocks.NewRedisCommander(GinkgoT())
		adapter = redisCursors.NewAdapter(client)
	})
	AfterEach(func() {
		cancel()
	})
	It("create and retrieve a cursor", func() {
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

		cursorStr := redisCursors.CursorToString(req)
		Expect(cursorStr).ToNot(BeEmpty())

		cursorBytes, err := json.Marshal(req)
		Expect(err).ToNot(HaveOccurred())
		Expect(cursorBytes).ToNot(BeEmpty())

		// 	cmd := a.client.Set(ctx, fmt.Sprintf("%s-%s", cursorKeyPrefix, cur), cursorBytes, time.Hour)
		client.On("Set",
			mock.AnythingOfType("*context.cancelCtx"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("[]uint8"),
			mock.AnythingOfType("time.Duration"),
		).
			Return(redis.NewStatusResult("", nil)).
			Run(func(args mock.Arguments) {
				GinkgoWriter.Write([]byte(fmt.Sprintf("args: %v\n", args)))
			})
		cursor, err := adapter.New(testCtx, req)

		Expect(err).ToNot(HaveOccurred())
		Expect(cursor).ToNot(BeEmpty())

		client.On("Get",
			mock.AnythingOfType("*context.cancelCtx"),
			mock.AnythingOfType("string"),
		).
			Return(redis.NewStringResult(string(cursorBytes), nil)).
			Run(func(args mock.Arguments) {
				GinkgoWriter.Write([]byte(fmt.Sprintf("args: %v\n", args)))
			})
		client.On("Expire",
			mock.AnythingOfType("*context.cancelCtx"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("time.Duration"),
		).
			Return(redis.NewBoolResult(true, nil)).
			Run(func(args mock.Arguments) {
				GinkgoWriter.Write([]byte(fmt.Sprintf("args: %v\n", args)))
			})
		actual, err := adapter.Get(testCtx, cursor)
		Expect(err).ToNot(HaveOccurred())
		Expect(actual).To(BeEquivalentTo(req))
	})
})
