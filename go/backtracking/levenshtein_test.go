package backtracking

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	a := "mitcmu"
	b := "mtacnu"
	minDist := len(a) + len(b)
	levenshteinDistance(0, 0, 0, a, b, &minDist)
	t.Log(minDist)
}
