package testdata

import (
	"container/ring"
	"testing"
)

type LRUCache struct {
	capacity int
	cache    map[string]*ring.Ring
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*ring.Ring),
	}
}

func (lru *LRUCache) Get(key string) string {
	if node, ok := lru.cache[key]; ok {
		return node.Value.(string)
	}
	return ""
}

func (lru *LRUCache) Put(key, value string) {
	if node, ok := lru.cache[key]; ok {
		node.Value = value
		// 将最近使用的节点移动到头部
		lru.cache[key] = node.Move(0)
	} else {
		// 创建新节点并添加到头部
		newNode := ring.New(1)
		newNode.Value = value
		lru.cache[key] = newNode
		// 如果超出容量，删除最久未使用的节点
		if len(lru.cache) > lru.capacity {
			delete(lru.cache, lru.cache[key].Prev().Value.(string))
		}
	}
}

func TestRingLru(t *testing.T) {

	cache := NewLRUCache(3)

	cache.Put("banana", "yellow")
	cache.Put("cherry", "red")
	cache.Put("apple", "red")

	//fmt.Println(cache.Get("apple"))
	//fmt.Println(cache.Get("banana"))
	//fmt.Println(cache.Get("cherry"))
	//fmt.Println(cache.Get("date"))
}
