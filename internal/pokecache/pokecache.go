package pokecache

import (
    "time"
    "sync"
)

type Cache struct {
    cache   map[string]cacheEntry
    mux     *sync.Mutex
}

type cacheEntry struct {
    createdAt   time.Time
    val         []byte
}

func NewCache(interval time.Duration) Cache {
    c := Cache{
        cache: make(map[string]cacheEntry),
        mux: &sync.Mutex{},
    }

    go c.reapLoop(interval)

    return c
}

// Add
// Add the value to the cache at the key.
func (c *Cache) Add(key string, value []byte) {
    c.mux.Lock()
    defer c.mux.Unlock()
    c.cache[key] = cacheEntry{
        createdAt: time.Now().UTC(),
        val: value,
    }
}

// Get
// Returns the value from the cache at the key and whether that value exists in the cache.
func (c *Cache) Get(key string) ([]byte, bool) {
    c.mux.Lock()
    defer c.mux.Unlock()
    value, exists := c.cache[key]
    return value.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
        c.reap(time.Now().UTC(), interval)
    }
}

func (c *Cache) reap(now time.Time, last time.Duration) {
    c.mux.Lock()
    defer c.mux.Unlock()
    for k, v := range c.cache {
        if v.createdAt.Before(now.Add(-last)) {
            delete(c.cache, k)
        }
    }
}
