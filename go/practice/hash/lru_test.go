package hash

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	lc := newLRUCache(100)
	k, v := "mgxian", "will"
	kvCount := 100

	for i := 0; i < kvCount; i++ {
		idx := strconv.Itoa(i)
		err := lc.set(k+idx, v+idx)
		assert.NoError(t, err)
	}

	err := lc.set(k, v)
	assert.Equal(t, nil, err)

	for i := 1; i < kvCount; i += 3 {
		idx := strconv.Itoa(i)
		iv, ok := lc.get(k + idx)
		assert.Equal(t, v+idx, iv)
		assert.Equal(t, true, ok)
	}

	iv, ok := lc.get("will0")
	assert.Equal(t, "", iv)
	assert.Equal(t, false, ok)

	for i := 1; i < kvCount; i += 5 {
		idx := strconv.Itoa(i)
		iv := lc.delete(k + idx)
		assert.Equal(t, v+idx, iv)
	}

	assert.Equal(t, "", lc.delete("will0"))
}
