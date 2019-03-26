package dynamic

import "testing"

func TestDouble11Adavance(t *testing.T) {
	items := []int{2, 2, 4, 6, 3}
	quota := 5
	choice := double11Adavance(items, quota)
	t.Log(choice)
}
