package store

import (
	"crypto/sha1"
	"fmt"
	"sync"
)

type ShardMap map[string]*Shard

type Shard struct {
	lock *sync.RWMutex
}

func (s Shard) Lock() {
	s.lock.Lock()
}

func (s Shard) Unlock() {
	s.lock.Unlock()
}

func (s Shard) RLock() {
	s.lock.RLock()
}

func (s Shard) RUnlock() {
	s.lock.RUnlock()
}

var _ sync.Locker = &Shard{}

func NewShardMap() ShardMap {
	s := make(ShardMap, 256)
	for i := 0; i < 256; i++ {
		s[fmt.Sprintf("%02x", i)] = &Shard{lock: new(sync.RWMutex)}
	}
	return s
}

func (s ShardMap) GetShard(key string) *Shard {
	hasher := sha1.New()
	hasher.Write([]byte(key))
	shardKey := fmt.Sprintf("%x", hasher.Sum(nil))[0:2]
	return s[shardKey]
}
