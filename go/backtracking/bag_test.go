package backtracking

import "testing"

func TestZeroOneBag(t *testing.T) {
	items := []int{9, 3, 6, 7}
	limitedWeight := 20
	maxWeight := -1
	zeroOneBag(0, 0, items, limitedWeight, &maxWeight)
	t.Log(maxWeight)
}
