package backtracking

import "testing"

func TestZeroOneBag(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	limitedWeight := 9
	maxWeight := -1
	zeroOneBag(0, 0, items, limitedWeight, &maxWeight)
	t.Log(maxWeight)
}

func TestZeroOneBagWithCache(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	limitedWeight := 9
	maxWeight := -1
	cache := make([][]bool, len(items))
	for i := range cache {
		cache[i] = make([]bool, limitedWeight+1)
	}
	zeroOneBagWithCache(0, 0, items, limitedWeight, &maxWeight, cache)
	t.Log(maxWeight)
}
