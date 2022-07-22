package activeents

import (
	"context"
	"fmt"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/event-management-service/pkg/models/scopes"
	"github.com/zenoss/zenkit/v5"
	"time"
)

func RunActiveEntitySpecs(offset int, factory func() scopes.ActiveEntityRepository) {
	var (
		ctx          context.Context
		cancel       context.CancelFunc
		adapter      scopes.ActiveEntityRepository
		testTenantId = zenkit.RandString(16)
	)
	ginkgo.BeforeEach(func() {
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		adapter = factory()
	})
	ginkgo.AfterEach(func() {
		cancel()
	})
	ginkgo.Context("when repo is empty", ginkgo.Offset(offset+1), func() {
		ginkgo.It("should pass the requested IDs back", func() {
			input := []string{"e1", "e2"}
			now := time.Now()
			actual, err := adapter.Get(ctx, event.TimeRange{
				Start: now.Add(-24 * time.Hour).UnixMilli(),
				End:   now.UnixMilli(),
			}, testTenantId, input)
			gomega.Expect(err).ToNot(gomega.HaveOccurred())
			gomega.Expect(actual).To(gomega.ConsistOf(input))
		})
	})
	ginkgo.Context("when repo has entries", ginkgo.Offset(1), func() {
		var (
			dataEnd   = time.Now().UnixMilli()
			dataStart = time.Now().Add(-48 * time.Hour).UnixMilli()
			entityIDs = []string{"e1", "e2"}
		)
		ginkgo.JustBeforeEach(func() {
			adapter.Put(ctx,
				&scopes.PutActiveEntityRequest{
					Inputs: []*scopes.ActiveEntityInput{
						{
							Timestamp: dataStart,
							TenantID:  testTenantId,
							EntityID:  entityIDs[0],
						},
						{
							Timestamp: dataStart + (24 * time.Hour).Milliseconds(),
							TenantID:  testTenantId,
							EntityID:  entityIDs[1],
						},
						{
							Timestamp: dataEnd,
							TenantID:  testTenantId,
							EntityID:  entityIDs[0],
						},
						{
							Timestamp: dataEnd,
							TenantID:  testTenantId,
							EntityID:  entityIDs[1],
						},
					}})
		})
		ginkgo.It("should return the expected entity IDs", func() {
			for _, testCase := range []struct {
				start, end          int64
				entityIDs, expected []string
			}{
				{start: dataStart - (24 * time.Hour).Milliseconds(), end: dataStart, entityIDs: []string{entityIDs[0]}, expected: []string{entityIDs[0]}},
				{start: dataStart - (24 * time.Hour).Milliseconds(), end: dataStart + (24 * time.Hour).Milliseconds() + 1, entityIDs: []string{entityIDs[1]}, expected: []string{entityIDs[1]}},
				{start: dataStart - (24 * time.Hour).Milliseconds(), end: dataEnd, entityIDs: entityIDs, expected: entityIDs},
				{start: dataStart - (24 * time.Hour).Milliseconds(), end: dataEnd, entityIDs: []string{"e3", "e4"}, expected: nil},
			} {
				actual, _ := adapter.Get(
					ctx,
					event.TimeRange{Start: testCase.start, End: testCase.end},
					testTenantId,
					testCase.entityIDs)
				gomega.Expect(actual).
					To(gomega.ConsistOf(testCase.expected), fmt.Sprintf("adapter.Get(..., %#v): %#v should equal %#v", testCase, actual, testCase.expected))
			}
		})
	})

}
