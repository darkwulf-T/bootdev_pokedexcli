package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	entries 	map[string]cacheEntry
	period 		time.Duration
	mu 			sync.RWMutex
}

type cacheEntry struct{
	createdAt 	time.Time
	val 		[]byte
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		entries: make(map[string]cacheEntry),
		period: interval,
	}

	go newCache.reapLoop()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, ok := c.entries[key]
	if !ok {
		return []byte{}, false
	}
	return data.val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.period)
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.entries {
			if time.Since(v.createdAt) > c.period {
				delete(c.entries, k)
			}
		}
		c.mu.Unlock()
	}
}