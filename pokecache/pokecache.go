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
        mu   sync.Mutex
        data map[string]cacheEntry
}

func NewCache(interval time.Duration) *Cache {
        c := &Cache{
                data: make(map[string]cacheEntry),
        }
        go c.reapLoop(interval)
        return c
}

func (c *Cache) Add(key string, val []byte) {
        c.mu.Lock()
        defer c.mu.Unlock()
        c.data[key] = cacheEntry{
                createdAt: time.Now(),
                val:       val,
        }
}

func (c *Cache) Get(key string) ([]byte, bool) {
        c.mu.Lock()
        defer c.mu.Unlock()
        entry, ok := c.data[key]
        if !ok {
                return nil, false
        }
        return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
        ticker := time.NewTicker(interval)
        for range ticker.C {
                c.mu.Lock()
                for k, v := range c.data {
                        if time.Since(v.createdAt) > interval {
                                delete(c.data, k)
                        }
                }
                c.mu.Unlock()
        }
}