package cache

import "sync"

type KeyValue struct {
	Key   string
	Value string
}

type KeyValueCache interface {
	Add(key string, val string)
	Get(key string) string
	Len() int
	GetAll() []KeyValue
}

type mapCache struct {
	m    map[string]string
	lock sync.Mutex
}

func NewMapCache() KeyValueCache {
	mc := &mapCache{
		m: make(map[string]string),
	}
	return mc
}

func (c *mapCache) Add(key string, val string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.m[key] = val
}

func (c *mapCache) Get(key string) string {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.m[key]
}

func (c *mapCache) Len() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return len(c.m)
}

func (c *mapCache) GetAll() []KeyValue {
	c.lock.Lock()
	defer c.lock.Unlock()

	res := make([]KeyValue, 0, len(c.m))
	for k, v := range c.m {
		res = append(res, KeyValue{Key: k, Value: v})
	}

	return res
}
