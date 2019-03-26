package dynamic

import (
	"testing"
)

func TestMinDist(t *testing.T) {
	w := [][]int{{1, 3, 5, 9}, {2, 1, 3, 4}, {5, 2, 6, 7}, {6, 8, 4, 3}}
	min := minDist(w)
	t.Log(min)
}
