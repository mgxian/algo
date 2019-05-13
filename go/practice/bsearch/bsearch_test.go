package bsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{7, 9, 10, 11, 12, 15, 20}

	i := binarySearch(nums, 7)
	assert.Equal(t, 0, i)

	i = binarySearch(nums, 9)
	assert.Equal(t, 1, i)

	i = binarySearch(nums, 20)
	assert.Equal(t, 6, i)

	i = binarySearch(nums, 30)
	assert.Equal(t, -1, i)
}
