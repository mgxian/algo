package main

import (
	"fmt"
	"strconv"

	"github.com/mgxian/algo/go/linkedlist/singly"
)

// LRUCache lrc cache
type LRUCache struct {
	list *singly.LinkedListWithDummyHead
	size int
}

// NewLRUCache create a lru cache
func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		list: singly.NewLinkedListWithDummyHead(),
		size: size,
	}
}

// Get get a cache
func (l LRUCache) Get(url string) {
	if l.list.Find(url) {
		l.list.Delete(url)
		l.list.InsertToHead(url)
		return
	}

	if l.list.Size() < l.size {
		l.list.InsertToHead(url)
		return
	}

	l.list.DeleteTail()
	l.list.InsertToHead(url)
}

// Print all lru cache element
func (l LRUCache) Print() {
	fmt.Println(l.list.Tranverse())
}

func main() {
	aLRUCache := NewLRUCache(10)
	for i := 0; i < 10; i++ {
		aLRUCache.Get(strconv.Itoa(i))
	}

	aLRUCache.Print()
	for i := 10; i > 0; i-- {
		aLRUCache.Get(strconv.Itoa(i))
	}
	aLRUCache.Print()
}
