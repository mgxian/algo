package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4}
	myInts := make([]MyInt, 0)
	for _, num := range numbers {
		myInts = append(myInts, NewMyInt(num))
	}

	BubbleSort(myInts)

	t.Log(myInts)
}
