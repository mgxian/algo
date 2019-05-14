package hash

import (
	"fmt"
	"strings"
)

type element struct {
	prev  *element
	next  *element
	hnext *element
	key   string
	value string
}

type separateChainList struct {
	head element
	len  int
}

func newSeparateChainList() *separateChainList {
	l := new(separateChainList)
	l.head.prev = &l.head
	l.head.hnext = nil
	l.len = 0
	return l
}

func (scl *separateChainList) findBykey(key string) (prev *element, curr *element, value string) {
	prev = &scl.head
	curr = scl.head.hnext
	for curr != nil {
		if curr.key == key {
			break
		}
		prev = curr
		curr = curr.hnext
	}
	if curr != nil {
		value = curr.value
	}
	return
}

func (scl *separateChainList) pushBack(key, value string) *element {
	_, curr, _ := scl.findBykey(key)
	if curr != nil {
		return nil
	}

	e := new(element)
	e.key = key
	e.value = value

	scl.head.prev.hnext = e
	scl.head.prev = e

	scl.len++
	return e
}

func (scl *separateChainList) moveToBack(key string) *element {
	prev, curr, _ := scl.findBykey(key)
	if curr == nil || curr == scl.head.prev {
		return curr
	}

	prev.hnext = curr.hnext
	curr.hnext = nil
	scl.head.prev.hnext = curr

	scl.head.prev = curr

	return curr
}

func (scl *separateChainList) remove(key string) *element {
	prev, curr, _ := scl.findBykey(key)
	if curr == nil {
		return curr
	}

	prev.hnext = curr.hnext
	if curr.hnext == nil {
		scl.head.prev = prev
	}

	scl.len--
	return curr
}

func (scl *separateChainList) String() string {
	result := new(strings.Builder)
	title := fmt.Sprintf("length: %d ", scl.len)
	result.WriteString(title)

	p := scl.head.hnext
	for p != nil && p != scl.head.prev {
		eString := fmt.Sprintf("(%s, %s)--->", p.key, p.value)
		result.WriteString(eString)
		p = p.hnext
	}

	if p != nil {
		eString := fmt.Sprintf("(%s, %s)", p.key, p.value)
		result.WriteString(eString)
	}

	return result.String()
}

type lruDoublyList struct {
	head element
	len  int
}

func newLRUDoublyList() *lruDoublyList {
	l := new(lruDoublyList)
	l.head.prev = &l.head
	l.head.next = nil
	l.len = 0
	return l
}

func (ldl *lruDoublyList) remove(e *element) *element {
	n := e.next
	e.prev.next = n

	if n != nil {
		n.prev = e.prev
	} else {
		ldl.head.prev = e.prev
	}

	e.prev = nil
	e.next = nil

	ldl.len--
	return e
}

func (ldl *lruDoublyList) insert(e, at *element) *element {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n

	if n != nil {
		n.prev = e
	} else {
		ldl.head.prev = e
	}

	ldl.len++
	return e
}

func (ldl *lruDoublyList) move(e, at *element) *element {
	if e == at {
		return e
	}

	ldl.remove(e)
	ldl.insert(e, at)

	return e
}

func (ldl *lruDoublyList) moveToBack(e *element) *element {
	if ldl.head.prev == e {
		return e
	}
	ldl.move(e, ldl.head.prev)
	return e
}

func (ldl *lruDoublyList) moveToFront(e *element) *element {
	if ldl.head.next == e {
		return e
	}
	ldl.move(e, &ldl.head)
	return e
}

func (ldl *lruDoublyList) pushBack(e *element) *element {
	ldl.insert(e, ldl.head.prev)
	return e
}

func (ldl *lruDoublyList) String() string {
	result := new(strings.Builder)
	title := fmt.Sprintf("length: %d ", ldl.len)
	result.WriteString(title)

	p := ldl.head.next
	for p != nil && p != ldl.head.prev {
		eString := fmt.Sprintf("(%s, %s)--->", p.key, p.value)
		result.WriteString(eString)
		p = p.next
	}

	if p != nil {
		eString := fmt.Sprintf("(%s, %s)", p.key, p.value)
		result.WriteString(eString)
	}

	return result.String()
}

type lruCache struct {
	data        []*separateChainList
	ldl         *lruDoublyList
	capacity    int
	bucketCount int
}

func newLRUCache(capacity int) *lruCache {
	bucketSize := 10
	bucketCount := capacity / bucketSize
	data := make([]*separateChainList, bucketCount)
	for i := 0; i < bucketCount; i++ {
		data[i] = newSeparateChainList()
	}
	return &lruCache{
		data:        data,
		ldl:         newLRUDoublyList(),
		capacity:    capacity,
		bucketCount: bucketCount,
	}
}

func (lru *lruCache) exist(key string) bool {
	code := lru.hashCode(key)
	l := lru.data[code]
	_, curr, _ := l.findBykey(key)
	if curr == nil {
		return false
	}
	return true
}

func (lru *lruCache) set(key string, value string) {
	code := lru.hashCode(key)
	l := lru.data[code]

	_, curr, _ := l.findBykey(key)
	if curr != nil {
		fmt.Println("exist", key, curr.value)
		curr.value = value
		lru.ldl.moveToBack(curr)
		return
	}

	// fmt.Println("not exist", key)
	if e := l.pushBack(key, value); e != nil {
		lru.ldl.pushBack(e)
	}
}

func (lru *lruCache) delete(key string) {
	code := lru.hashCode(key)
	l := lru.data[code]

	if e := l.remove(key); e != nil {
		lru.ldl.remove(e)
	}
}

func (lru *lruCache) get(key string) (string, bool) {
	code := lru.hashCode(key)
	l := lru.data[code]
	_, curr, v := l.findBykey(key)
	if curr != nil {
		return v, true
	}
	return "", false
}

func (lru *lruCache) hashCode(key string) int {
	code := 0
	for _, c := range []byte(key) {
		code += int(c)
	}
	return code % lru.bucketCount
}

func (lru *lruCache) String() string {
	result := ""
	for i, l := range lru.data {
		result += fmt.Sprintf("bucket %d : %s\n", i, l.String())
	}

	result += "all : " + lru.ldl.String()
	return result
}
