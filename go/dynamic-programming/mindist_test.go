package dynamic

import (
	"testing"
)

func TestMinDist(t *testing.T) {
	w := [][]int{{1, 3, 5, 9}, {2, 1, 3, 4}, {5, 2, 6, 7}, {6, 8, 4, 3}}
	min := minDist(w)
	t.Log(min)
}

func TestMinDist2(t *testing.T) {
	w := [][]int{{1, 3, 5, 9}, {2, 1, 3, 4}, {5, 2, 6, 7}, {6, 8, 4, 3}}
	cache := make([][]int, len(w))
	for i := range cache {
		cache[i] = make([]int, len(w))
	}

	min := minDist2(len(w)-1, len(w)-1, w, cache)
	t.Log(min)
}
