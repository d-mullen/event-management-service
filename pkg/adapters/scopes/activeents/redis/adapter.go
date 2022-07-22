package redis

import (
	"context"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/go-redis/redis/v8"
	"github.com/zenoss/event-management-service/pkg/adapters/scopes/activeents"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/event-management-service/pkg/models/scopes"
)

type (
	redisActiveEntityAdapter struct {
		bucketSize time.Duration
		defaultTTL time.Duration
		client     *redis.Ring
	}
)

// NewInMemoryActiveEntityAdapter return an implementation of the `ActiveEntityRepository` interface that stores
// and retrieves sets of IDs for active entities in Redis. This implementation uses the Set data structure for fast
// computation of large collection of entity IDs.
func NewAdapter(bucketSize, defaultTTL time.Duration, client *redis.Ring) *redisActiveEntityAdapter {
	return &redisActiveEntityAdapter{
		client:     client,
		defaultTTL: defaultTTL,
		bucketSize: bucketSize,
	}
}

// Put takes an array of inputs, each representing an instance in time for which an event was recorded
// for an entity, and adds them to appropriate set of entities with active event occurrences.
func (adapter *redisActiveEntityAdapter) Put(ctx context.Context, request *scopes.PutActiveEntityRequest) error {
	entries := make(map[string][]any)
	// Iterate over the request input and group the given entity IDs by the respective key.
	for _, input := range request.Inputs {
		key := activeents.GetKey(input.Timestamp, adapter.bucketSize, input.TenantID)
		if ids, ok := entries[key]; !ok {
			entries[key] = []any{input.EntityID}
		} else {
			entries[key] = append(ids, input.EntityID)
		}
	}
	// For each bucket found in the input, add the corresponding entity Ids to that set.
	for key, entityIds := range entries {
		intCmd := adapter.client.SAdd(ctx, key, entityIds...)
		if err := intCmd.Err(); err != nil {
			return err
		}
		if cmd := adapter.client.Expire(ctx, key, adapter.defaultTTL); cmd.Err() != nil {
			return cmd.Err()
		}
	}
	return nil
}

// Get takes an array of entity IDs, a duration, and tenant ID and returns an array of IDs of entities known to have
// events recorded in the given time range. If no active entity sets are stored for the requested time, then
// the original entityIDs are returned to the caller.
func (adapter *redisActiveEntityAdapter) Get(ctx context.Context, timeRange event.TimeRange, tenantID string, entityIDs []string) ([]string, error) {
	keys := activeents.GetKeys(timeRange.Start, timeRange.End, adapter.bucketSize, tenantID)
	sliceCmd := adapter.client.SUnion(ctx, keys...)
	if err := sliceCmd.Err(); err != nil {
		return nil, err
	}
	idSlice := sliceCmd.Val()
	// if no active entity sets were found, then pass all entities through
	if len(idSlice) == 0 {
		return entityIDs, nil
	}
	activeEntities := mapset.NewSet(idSlice...)
	otherSet := mapset.NewSet(entityIDs...)
	intersection := activeEntities.Intersect(otherSet)
	return intersection.ToSlice(), nil
}

var _ scopes.ActiveEntityRepository = &redisActiveEntityAdapter{}
