package pokecache

import (
	"time"
	"sync"
)


type cacheEntry struct {
    createdAt		time.Time
    val			[]byte
}

type Cache struct {
    entries		map[string]cacheEntry
    mu			sync.Mutex
}


func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
	    entries:	make(map[string]cacheEntry),
	    mu:		sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
}


func (c *Cache) Add(key string, val []byte) {
	entry := cacheEntry{
	    createdAt:	time.Now(),
	    val:	val,
	}
	c.mu.Lock()
	c.entries[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.entries[key]
	c.mu.Unlock()
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.entries {
		    if time.Now().Sub(entry.createdAt) > interval {
			delete(c.entries, key)
		    }
		}
		c.mu.Unlock()
	}
}
