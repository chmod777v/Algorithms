package main

import (
	"fmt"
	"hash/fnv"
	"sync"
)

// Написать реализацию InMemory кэша

// key1: "bar", value1 hash=6126512941 %5 ->остаток 0-4
// key2: "bar2", value2 hash=4124213421 %5
// [0-key1] [1] [2] [3-key2] [4] в зависимости какой остаток кладем в опредеделенную мапу
type Cache interface {
	Set(k string, v string)
	Get(k string) (string, bool)
}
type Shard struct {
	data map[string]string
	mu   sync.RWMutex
}

type InMemoryCache struct {
	shards []Shard
}

func (c *InMemoryCache) Set(k string, v string) {
	shardId := hasher(k) % len(c.shards)
	fmt.Println(shardId)
	c.shards[shardId].Set(k, v)
}

func (c *InMemoryCache) Get(k string) (string, bool) {
	shardId := hasher(k) % len(c.shards)
	data, ok := c.shards[shardId].Get(k)
	return data, ok
}

func hasher(k string) int {
	h := fnv.New32a()
	h.Write([]byte(k))
	return int(h.Sum32())
}
func NewInMemoryCache(numShards int) *InMemoryCache {
	shards := make([]Shard, 0, numShards)
	for i := 0; i < numShards; i++ {
		shards = append(shards, Shard{data: make(map[string]string)})
	}
	return &InMemoryCache{
		shards: shards,
	}
}

func (s *Shard) Set(k string, v string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[k] = v
}

func (s *Shard) Get(k string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	data, ok := s.data[k]
	return data, ok
}

func main() {
	cache := NewInMemoryCache(5)
	cache.Set("foo", "bar")
	cache.Set("foo1", "bar1")
	cache.Set("123", "bar2")

	data, ok := cache.Get("foo")
	if !ok {
		fmt.Println("Not found")
		return
	}
	fmt.Println(data)

}

// в production часто используется consistent hashing
