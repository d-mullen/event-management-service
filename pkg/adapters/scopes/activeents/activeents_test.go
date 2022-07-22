package activeents_test

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/zenoss/event-management-service/pkg/adapters/scopes/activeents"
	"github.com/zenoss/event-management-service/pkg/models/scopes"
)

var _ = Describe("In-Memory Active Entity Adapter", func() {
	activeents.RunActiveEntitySpecs(0, func() scopes.ActiveEntityRepository {
		return activeents.NewInMemoryActiveEntityAdapter(1024, activeents.DefaultBucketSize)
	})
})
