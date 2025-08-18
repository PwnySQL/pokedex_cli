package pokecache

import (
	"sync"
	"time"
)

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.cache[key] = cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	val, isKeyInCache := cache.cache[key]
	if !isKeyInCache {
		return []byte{}, isKeyInCache
	}
	return val.val, isKeyInCache
}

func NewCache() Cache {
	return Cache{
		cache: map[string]cacheEntry{},
		mutex: sync.Mutex{},
	}
}
