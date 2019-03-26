package dynamic

import "testing"

func TestLcsLength(t *testing.T) {
	a := "mitcmu"
	b := "mtacnu"
	max := lcsLength(a, b)
	t.Log(max)
}
