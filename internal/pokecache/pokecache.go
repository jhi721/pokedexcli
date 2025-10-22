package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheInterface interface {
	Add(key string, val []byte)
	Get(key string) ([]byte, bool)
	reapLoop()
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, exists := c.data[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			c.mu.Lock()

			for key, entry := range c.data {
				if entry.createdAt.UnixMilli() < time.Now().UnixMilli() {
					delete(c.data, key)
				}
			}

			c.mu.Unlock()
		}
	}()
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}

	cache.reapLoop(interval)

	return cache
}
