package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache   map[string]cacheEntry
	cacheMu sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	cache := make(map[string]cacheEntry)
	c := Cache{
		cache: cache,
	}

	go c.readLoop(interval)

	return &c
}

func (c *Cache) readLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		c.cacheMu.Lock()

		for key, value := range c.cache {
			if now.Sub(value.createdAt) > interval {
				delete(c.cache, key)
			}
		}

		c.cacheMu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cacheMu.Lock()
	defer c.cacheMu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) (cacheEntry, bool) {
	c.cacheMu.RLock()
	defer c.cacheMu.RUnlock()

	entry, exists := c.cache[key]
	return entry, exists
}
