package lru

import (
	"fmt"
	"testing"
)

type CustomValue struct {
	name string
}

func (v *CustomValue) Len() int {
	return len(v.name)
}

func TestGet(t *testing.T) {
	lruCache := NewCache(int64(1024), nil)
	lruCache.Add("key1", &CustomValue{name: "1234"})
	if v, ok := lruCache.Get("key1"); !ok || v.(*CustomValue).name != "1234" {
		t.Fatalf("cache hit key1=1234 failed")
	}
	if _, ok := lruCache.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}

func TestAdd(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := &CustomValue{name: "v1"}, &CustomValue{name: "v2"}, &CustomValue{name: "v3"}
	cap := len(k1) + len(k2) + v1.Len() + v2.Len()
	lruCache := NewCache(int64(cap), nil)
	lruCache.Add(k1, v1)
	lruCache.Add(k2, v2)
	lruCache.Add(k3, v3)
	if _, ok := lruCache.Get("key1"); ok || lruCache.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lruCache := NewCache(int64(12), callback)
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := &CustomValue{name: "v1"}, &CustomValue{name: "v2"}, &CustomValue{name: "v3"}
	lruCache.Add(k1, v1)
	lruCache.Add(k2, v2)
	lruCache.Add(k3, v3)

	for _, key := range keys {
		fmt.Println(key)
	}
}
