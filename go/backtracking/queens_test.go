package backtracking

import "testing"

func TestNQueens(t *testing.T) {
	n := 8
	ret := make([]int, n)
	nQueens(0, ret)
}
