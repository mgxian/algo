package backtracking

import "testing"

func TestLcsLength(t *testing.T) {
	a := "mitcmu"
	b := "mtacnu"
	max := lcsLength(len(a)-1, len(b)-1, a, b)
	t.Log(max)

	a = "mitcmu"
	b = "mitcmu"
	max = lcsLength(len(a)-1, len(b)-1, a, b)
	t.Log(max)
}
