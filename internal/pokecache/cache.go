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

func (cache *Cache) reapLoop(reapInterval time.Duration) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	now := time.Now()
	for k, v := range cache.cache {
		if now.Sub(v.createdAt) > reapInterval {
			delete(cache.cache, k)
		}
	}
}

func NewCache(reapInterval time.Duration) *Cache {
	cache := &Cache{
		cache: map[string]cacheEntry{},
		mutex: sync.Mutex{},
	}
	go func(tickCh <-chan time.Time, cache *Cache, reapInterval time.Duration) {
		for _ = range tickCh {
			cache.reapLoop(reapInterval)
		}
	}(time.Tick(reapInterval), cache, reapInterval)
	return cache
}
