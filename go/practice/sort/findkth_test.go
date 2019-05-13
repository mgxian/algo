package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKth(t *testing.T) {
	nums := []int{9, 6, 7, 4, 5, 1, 2, 3, 8}

	n := findKth(MyInts(nums), 3)
	expected := 7
	assert.Equal(t, expected, nums[n])

	nums = []int{12, 11, 15, 7, 9, 20, 10}

	n = findKth(MyInts(nums), 2)
	expected = 15
	assert.Equal(t, expected, nums[n])

	n = findKth(MyInts(nums), 4)
	expected = 11
	assert.Equal(t, expected, nums[n])
}
