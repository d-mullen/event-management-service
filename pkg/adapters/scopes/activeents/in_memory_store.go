package activeents

import (
	"context"
	"errors"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	lru "github.com/hashicorp/golang-lru"
	"github.com/zenoss/event-management-service/pkg/models/event"
	"github.com/zenoss/event-management-service/pkg/models/scopes"
)

type (
	inMemoryActiveEntityAdapter struct {
		mu         *sync.RWMutex
		bucketSize time.Duration
		data       *lru.Cache
	}
)

// NewInMemoryActiveEntityAdapter returns an implementation of the `ActiveEntityRepository` interface that stores
// active-entity sets in memory.
func NewInMemoryActiveEntityAdapter(size int, bucketSize time.Duration) *inMemoryActiveEntityAdapter {
	c, err := lru.New(size)
	if err != nil {
		panic(err)
	}
	return &inMemoryActiveEntityAdapter{
		data:       c,
		bucketSize: bucketSize,
		mu:         &sync.RWMutex{},
	}
}

// Put takes an array of inputs, each representing an instance in time for which an event was recorded
// for an entity, and adds them to appropriate set of entities with active event occurrences.
func (adapter inMemoryActiveEntityAdapter) Put(_ context.Context, request *scopes.PutActiveEntityRequest) error {
	// For each input, retrieve or create the set of active entities for the configured time bucket ...
	for _, input := range request.Inputs {
		key := GetKey(input.Timestamp, adapter.bucketSize, input.TenantID)
		adapter.mu.Lock()
		var set mapset.Set[string]
		setAny, ok := adapter.data.Get(key)
		if !ok {
			set = mapset.NewSet[string]()
		} else {
			var ok2 bool
			set, ok2 = setAny.(mapset.Set[string])
			if !ok2 {
				return errors.New("got invalid set")
			}
		}
		// ... and add that entity ID to the corresponding set
		set.Add(input.EntityID)
		adapter.data.Add(key, set)
		adapter.mu.Unlock()
	}
	return nil
}

// Get takes an array of entity IDs, a duration, and tenant ID and returns an array of IDs of entities known to have
// events recorded in the given time range. If no active entity sets are stored for the requested time, then
// original entityIDs are returned to the caller.
func (adapter inMemoryActiveEntityAdapter) Get(_ context.Context, timeRange event.TimeRange, tenantID string, entityIDs []string) ([]string, error) {
	activeEntities := mapset.NewSet[string]()
	foundAnySets := false
	adapter.mu.RLock()
	// For a given time range and tenantID, retrieve the corresponding sets of active entities ...
	for _, key := range GetKeys(timeRange.Start, timeRange.End, adapter.bucketSize, tenantID) {
		if setAny, ok := adapter.data.Get(key); ok {
			set, ok2 := setAny.(mapset.Set[string])
			if !ok2 {
				continue
			}
			// ... and combine them into a single set.
			activeEntities = activeEntities.Union(set)
			foundAnySets = true
		}
	}
	adapter.mu.RUnlock()
	// if no active entity sets were found, then pass all entities through
	if !foundAnySets {
		return entityIDs, nil
	}
	// only return entity IDs found in the intersection of the combined active entity set.
	otherSet := mapset.NewSet(entityIDs...)
	intersection := activeEntities.Intersect(otherSet)
	results := intersection.ToSlice()
	return results, nil
}
