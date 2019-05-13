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

func TestBinarySearchFirstEqual(t *testing.T) {
	nums := []int{7, 9, 10, 11, 11, 11, 12, 15, 20}

	i := binarySearchFirstEqual(nums, 7)
	assert.Equal(t, 0, i)

	i = binarySearchFirstEqual(nums, 11)
	assert.Equal(t, 3, i)

	i = binarySearchFirstEqual(nums, 20)
	assert.Equal(t, 8, i)

	i = binarySearchFirstEqual(nums, 30)
	assert.Equal(t, -1, i)
}

func TestBinarySearchLastEqual(t *testing.T) {
	nums := []int{7, 9, 10, 11, 11, 11, 12, 15, 20}

	i := binarySearchLastEqual(nums, 7)
	assert.Equal(t, 0, i)

	i = binarySearchLastEqual(nums, 11)
	assert.Equal(t, 5, i)

	i = binarySearchLastEqual(nums, 20)
	assert.Equal(t, 8, i)

	i = binarySearchLastEqual(nums, 30)
	assert.Equal(t, -1, i)
}
