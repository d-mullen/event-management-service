package zenkit

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CacheExpiration() time.Duration {
	raw := viper.GetString(GCMemstoreTTLConfig)
	dur, err := time.ParseDuration(raw)
	if err != nil {
		logrus.WithError(err).WithField("cache_duration", raw).Error("Could not parse cache duration")
	}
	return dur
}

func NewCache() *cache.Cache {
	return NewCacheId(0)
}

func NewCacheId(dbid int) *cache.Cache {
	opts := &cache.Options{
		Marshal:   json.Marshal,
		Unmarshal: json.Unmarshal,
	}
	if rds := NewRedisRingId(dbid); rds != nil {
		opts.Redis = rds
	} else {
		maxLen := viper.GetInt(GCMemstoreLocalMaxLen)
		opts.LocalCache = cache.NewTinyLFU(maxLen, CacheExpiration())
	}

	return cache.New(opts)
}

func NewRedisRing() *redis.Ring {
	return NewRedisRingId(0)
}

func NewRedisRingId(dbid int) *redis.Ring {
	redisAddrs := viper.GetStringSlice(GCMemstoreAddressConfig)
	if len(redisAddrs) == 0 {
		return nil
	}
	addrRing := make(map[string]string)
	for i, addr := range redisAddrs {
		addrRing[fmt.Sprintf("%d", i+1)] = addr
	}
	return redis.NewRing(&redis.RingOptions{
		Addrs: addrRing,
		DB:    dbid,
	})
}
