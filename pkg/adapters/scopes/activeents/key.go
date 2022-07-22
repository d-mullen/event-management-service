package activeents

import (
	"fmt"
	"sync"
	"time"

	"github.com/zenoss/event-management-service/config"
)

const (
	DefaultBucketSize = 24 * time.Hour
)

var (
	bucketSizeFloorMut     sync.Mutex
	defaultBucketSizeFloor = config.DefaultActiveEntityStoreMinBucketSize
)

// SetBucketSizeFloor overrides the currently set minimum duration used to bucket active entity record keys.
func SetBucketSizeFloor(dur time.Duration) {
	bucketSizeFloorMut.Lock()
	defaultBucketSizeFloor = dur
	bucketSizeFloorMut.Unlock()
}

// GetKey takes a timestamp, a duration and tenantID and generates a key
// to be used to access a set of active entities for a tenant for a given bucket of time.
func GetKey(ts int64, bucketSize time.Duration, tenantID string) string {
	// use a bucket size floor to avoid generating really small buckets
	_bucketSize := defaultBucketSizeFloor
	if bucketSize > _bucketSize {
		_bucketSize = bucketSize
	}
	bucket := time.UnixMilli(ts).Truncate(_bucketSize).UnixMilli()
	return fmt.Sprintf("active_entities#%s#%d", tenantID, bucket)
}

// GetKeys takes two timestamps representing a time range, a duration and tenantID and generates keys
// to be used to access one or more sets of active entities for a tenant for a given bucket of time.
func GetKeys(start, end int64, bucketSize time.Duration, tenantID string) []string {
	_bucketSize := defaultBucketSizeFloor
	// use a bucket size floor to avoid generating too many keys
	if bucketSize > _bucketSize {
		_bucketSize = bucketSize
	}
	keys := make([]string, 0)
	for i := start; i < end; i += _bucketSize.Milliseconds() {
		ts := time.UnixMilli(i).Truncate(_bucketSize).UnixMilli()
		keys = append(keys, GetKey(ts, _bucketSize, tenantID))
	}
	return keys
}
