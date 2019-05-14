package hash

import (
	"strconv"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lc := newLRUCache(100)
	k, v := "mgxian", "will"
	kvCount := 100

	for i := 0; i < kvCount; i++ {
		idx := strconv.Itoa(i)
		lc.set(k+idx, v+idx)
	}

	t.Log(lc.String())

	// for i := 0; i < kvCount; i++ {
	// 	idx := strconv.Itoa(i)
	// 	v, _ := lc.get(k + idx)
	// 	t.Log(v)
	// }
}
