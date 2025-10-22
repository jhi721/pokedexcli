package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "test",
			val: []byte("hello"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test vase %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, exists := cache.Get(c.key)
			if !exists {
				t.Errorf("expected to find key")
				return
			}

			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Second
	const waitTime = baseTime + 1*time.Second

	cache := NewCache(baseTime)
	cache.Add("test", []byte("test"))

	_, ok := cache.Get("test")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("test")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
