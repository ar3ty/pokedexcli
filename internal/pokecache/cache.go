package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	value     []byte
}

type Cache struct {
	data        map[string]cacheEntry
	mu          *sync.Mutex
	StopChannel chan int
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		data:        map[string]cacheEntry{},
		mu:          &sync.Mutex{},
		StopChannel: make(chan int),
	}
	cache.reapLoop(duration)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		value:     val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.data[key]
	return entry.value, ok
}

func (c *Cache) reap(duration time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, entry := range c.data {
		if time.Since(entry.createdAt) > duration {
			delete(c.data, key)
			fmt.Printf("DELETE\n")
		}
	}
}

func (c *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)

	go func() {
		for {
			select {
			case <-ticker.C:
				c.reap(duration)
			case <-c.StopChannel:
				ticker.Stop()
				return
			}
		}
	}()
}
