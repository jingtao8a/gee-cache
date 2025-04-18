package main

import (
	"org/jingtao8a/gee-cache/lru"
	"sync"
)

type Cache struct {
	lru        *lru.Cache
	mu         sync.Mutex
	cacheBytes int64
}

func (c *Cache) Add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.NewCache(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *Cache) Get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if value, ok := c.lru.Get(key); ok {
		return value.(ByteView), true
	}
	return
}
