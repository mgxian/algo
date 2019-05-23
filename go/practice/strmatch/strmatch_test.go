package strmatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContain(t *testing.T) {
	mains := []string{"abcdefg", "abcde", "", "bac", "cab", "cab"}
	patterns := []string{"efg", "ba", "", "ba", "c", "b"}
	indexs := []int{4, -1, -1, 0, 0, 2}
	for i := 0; i < len(mains); i++ {
		main := mains[i]
		pattern := patterns[i]
		assert.Equal(t, indexs[i], contain(main, pattern))
	}
}
