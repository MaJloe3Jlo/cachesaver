package iq

import (
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cache := NewCache(time.Second)
	if cache == nil {
		t.Errorf("Cache not created")
	}
}

func TestCache_Get(t *testing.T) {

	cache := NewCache(3 * time.Second)
	key := "one"
	value := "two"
	cache.Set(key, value, 1 * time.Second)
	val, ok := cache.Get(key)
	if val != value {
		t.Errorf("Value isn't correct")
	}
	if !ok {
		t.Errorf("There is no cache")
	}

	time.Sleep(time.Second)

	val, ok = cache.Get(key)
	if ok {
		t.Errorf("Time-to-live is ended \"ok\"")
	}
	if val == value {
		t.Errorf("Time-to-live is ended, value exist")
	}
}

func TestCache_Set(t *testing.T) {
	cache := NewCache(2 * time.Second)
	key := "one"
	value := "two"
	cache.Set(key, value, time.Second)
	if _, ok := cache.Get(key); !ok {
		t.Errorf("Value doesn't exist")
	}
}

func TestCache_Keys(t *testing.T) {
	cache := NewCache(time.Second)
	key := "one"
	value := "two"
	cache.Set(key, value, time.Second)
	{
		keys := cache.Keys()

		if len(keys) != 1 {
			t.Errorf("Isn't correct len of keys, %d", len(keys))
		}
	}
	time.Sleep(time.Second)
	{
		keys := cache.Keys()
		if len(keys) != 0 {
			t.Errorf("Isn't correct len of keys after end time-to-live, %d", len(keys))
		}
	}
}

func TestCache_Remove(t *testing.T) {
	cache := NewCache(3*time.Second)
	key := "one"
	value := "two"
	cache.Set(key, value, 2*time.Second)
	cache.Remove(key)
	val, ok := cache.Get(key)
	if ok {
		t.Errorf("Cache doesn't removed")
	}
	if val == value {
		t.Errorf("Value exist after removing")
	}

}