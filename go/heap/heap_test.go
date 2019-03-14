package heap

import "testing"

func TestHeap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
	aHeap := NewHeap(10)

	for _, v := range numbers {
		aHeap.Insert(v)
	}

	t.Log(aHeap)

	for range numbers {
		t.Log(aHeap.RemoveMax())
		t.Log(aHeap)
	}
}
