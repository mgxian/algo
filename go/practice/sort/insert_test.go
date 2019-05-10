package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSort(t *testing.T) {
	nums := []int{9, 6, 7, 4, 5, 1, 2, 3, 8}
	insertSort(MyInts(nums))

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	assert.Equal(t, expected, nums)
}
