package dynamic

import "testing"

func TestLevenshteinDistance(t *testing.T) {
	a := "mitcmu"
	b := "mtacnu"
	minDist := levenshteinDistance(a, b)
	t.Log(minDist)
}
