package main

import (
	"errors"
	"log"
	"sync"
)

type Group struct {
	name      string
	getter    Getter
	mainCache *Cache
}

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (g GetterFunc) Get(key string) ([]byte, error) {
	return g(key)
}

var (
	mu     sync.Mutex
	groups = make(map[string]*Group)
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("getter is nil")
	}
	mu.Lock()
	defer mu.Unlock()
	group, ok := groups[name]
	if ok {
		return group
	}
	g := &Group{
		name:   name,
		getter: getter,
		mainCache: &Cache{
			cacheBytes: cacheBytes,
		},
	}
	groups[name] = g
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, errors.New("empty key")
	}
	v, ok := g.mainCache.Get(key)
	if ok {
		log.Println("[GeeCache] hit")
		return v, nil
	}
	return g.load(key)
}

func (g *Group) load(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	copyBytes := make([]byte, len(bytes))
	copy(copyBytes, bytes)
	value := ByteView{b: copyBytes}
	g.mainCache.Add(key, value)
	return value, nil
}
