package dynamic

import "testing"

func TestZeroOneBag(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	limitedWeight := 16
	maxWeight := zeroOneBag(items, limitedWeight)
	t.Log(maxWeight)
}

func TestZeroOneBag2(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	limitedWeight := 16
	maxWeight := zeroOneBag2(items, limitedWeight)
	t.Log(maxWeight)
}
