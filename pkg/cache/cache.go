package cache

import (
	"strconv"
	"sync"
)

var caches = map[string]*Cache{}
var lock = sync.Mutex{}

func New(name string) *Cache {
	lock.Lock()
	defer lock.Unlock()
	if c, ok := caches[name]; ok {
		return c
	}
	c := &Cache{items: make(map[string]interface{})}
	caches[name] = c

	return c
}

type Cache struct {
	lock  sync.RWMutex
	items map[string]interface{}
}

func (c *Cache) Set(key string, val interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[key] = val
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, ok := c.items[key]
	return item, ok
}

func (c *Cache) GetByInt(key int) (interface{}, bool) {
	return c.Get(strconv.Itoa(key))
}

func (c *Cache) GetAll() map[string]interface{} {
	c.lock.RLock()
	c.lock.RUnlock()
	m := make(map[string]interface{}, len(c.items))
	for k, v := range c.items {
		m[k] = v
	}

	return m
}

func (c *Cache) Replace(items map[string]interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items = items
}

func (c *Cache) Clear() {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items = make(map[string]interface{})
}
