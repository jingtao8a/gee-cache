package lru

import "container/list"

type Cache struct {
	maxBytes  int64 // 允许使用的最大内存
	nbytes    int64 // 当前已使用的内存
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value) // 某条记录被驱逐的回调函数
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

func NewCache(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (Value, bool) {
	elem, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	c.ll.MoveToFront(elem)
	return elem.Value.(*entry).value, true
}

func (c *Cache) RemoveOldest() {
	elem := c.ll.Back()
	if elem == nil {
		return
	}
	c.ll.Remove(elem)
	delete(c.cache, elem.Value.(*entry).key)
	c.nbytes -= int64(len(elem.Value.(*entry).key)) + int64(elem.Value.(*entry).value.Len())
	if c.OnEvicted != nil {
		c.OnEvicted(elem.Value.(*entry).key, elem.Value.(*entry).value)
	}
}

func (c *Cache) Add(key string, value Value) {
	elem, ok := c.cache[key]
	if ok {
		c.ll.MoveToFront(elem)
		kv := elem.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
		return
	}
	elem = c.ll.PushFront(&entry{key, value})
	c.cache[key] = elem
	c.nbytes += int64(len(key)) + int64(value.Len())
	for c.maxBytes > 0 && c.maxBytes < c.nbytes { // 保证cache内存占用不超过maxBytes
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
