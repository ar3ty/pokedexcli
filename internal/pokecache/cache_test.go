package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "https://example.com",
			value: []byte("testdata"),
		},
		{
			key:   "https://example.com/path",
			value: []byte("moretestdata"),
		},
		{
			key:   "https://example.com/data",
			value: []byte("anotherpileoftestdata"),
		},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			defer func() {
				cache.StopChannel <- 1
			}()

			cache.Add(c.key, c.value)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.value) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReaper(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = 10 * time.Millisecond
	cache := NewCache(interval)
	defer func() {
		cache.StopChannel <- 1
	}()

	cache.Add("https://example.com", []byte("testdata"))
	val, ok := cache.Get("https://example.com")
	fmt.Printf("%v\n", val)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	val, ok = cache.Get("https://example.com")
	fmt.Printf("%v\n", val)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
