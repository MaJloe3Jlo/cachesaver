package iq

import (
	"time"
	"sync"
)

type Cache struct {
	ttl time.Duration
	Cache map[string]*data
	lock *sync.RWMutex
}

type data struct {
	val interface{}
	exp *time.Time
}


func NewCache(interval time.Duration) *Cache {
	if interval < time.Second {
		interval = time.Second
	}
	cache := &Cache{
		ttl: interval,
		Cache: make(map[string]*data),
		lock: &sync.RWMutex{},
	}

	go func() {
		ticker := time.NewTicker(cache.ttl)
		for {
			now := <- ticker.C

			cache.lock.Lock()
			for id, data := range cache.Cache {
				if data.exp != nil && data.exp.Before(now) {
					delete(cache.Cache, id)
				}
			}
			cache.lock.Unlock()
		}
	}()
	return cache
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	e, ok := c.Cache[key]

	if ok && e.exp != nil && e.exp.After(time.Now()) {
		return e.val, true
	}
	return nil, false
}

func (c *Cache) Set(key string, val interface{}, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	exp := time.Now().Add(ttl)

	c.Cache[key] = &data{
		val: val,
		exp: &exp,
	}
}

func (c *Cache) Remove(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()

	delete(c.Cache, key)
}

func (c *Cache) Keys() []interface{} {
	c.lock.RLock()
	defer  c.lock.RUnlock()

	keys := make([]interface{}, len(c.Cache))
	var i int
	for k:= range c.Cache{
		keys[i] = k
		i++
	}
	return keys
}