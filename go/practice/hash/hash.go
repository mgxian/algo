package hash

import (
	"errors"

	"github.com/mgxian/algo/go/practice/list"
)

type kv struct {
	key   string
	value interface{}
}

type hashMap struct {
	data        []list.List
	bucketCount int
}

func newHashMap(bucketCount int) *hashMap {
	return &hashMap{
		data:        make([]list.List, bucketCount),
		bucketCount: bucketCount,
	}
}

func (hm *hashMap) save(key string, value interface{}) error {
	code := hm.hashCode(key)
	l := hm.data[code]
	if l == nil {
		l = list.NewSinglyLinkedList()
		hm.data[code] = l
	}
	if e := l.PushBack(kv{key: key, value: value}); e == nil {
		return errors.New("save error")
	}
	return nil
}

func (hm *hashMap) delete(key string) {
	code := hm.hashCode(key)
	l := hm.data[code]
	if l == nil {
		return
	}
	p := l.Front()
	for p != nil {
		if p.Value.(kv).key == key {
			l.Remove(p)
			return
		}
		p = p.Next()
	}
}

func (hm *hashMap) get(key string) (interface{}, bool) {
	code := hm.hashCode(key)
	l := hm.data[code]
	if l == nil {
		return nil, false
	}
	p := l.Front()
	for p != nil {
		if p.Value.(kv).key == key {
			return p.Value.(kv).value, true
		}
		p = p.Next()
	}
	return nil, false
}

func (hm *hashMap) hashCode(key string) int {
	code := 0
	for _, c := range []byte(key) {
		code += int(c)
	}
	return code % hm.bucketCount
}
