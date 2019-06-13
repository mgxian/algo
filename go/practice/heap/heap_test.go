package heap

import (
	"testing"
	"reflect"
)

func TestIsHeap(t *testing.T) {
	tests := []struct {
		nums   []int
		isHeap bool
	}{
		{[]int{10, 9, 8, 6, 5, 7, 4, 3}, true},
		{[]int{10, 8, 9, 7, 3, 5, 6, 4}, true},
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{2, 1}, true},
		{[]int{2, 3}, false},
	}

	for _, tt := range tests {
		got := isHeap(tt.nums)
		if got != tt.isHeap {
			t.Errorf("got %t, want %t", got, tt.isHeap)
		}
	}
}

func TestHeap(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := isHeap(nums)
	if got != false {
		t.Errorf("got %v, want %v", got, false)
	}
	heap(nums)
	got = isHeap(nums)
	if got != true {
		t.Errorf("got %v, want %v", got, true)
	}
	t.Log(nums)
}


func TestBuildHeap(t *testing.T){
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	buildHeap(nums)
	got := isHeap(nums)
	if got != true {
		t.Errorf("got %v, want %v", got, true)
	}
	t.Log(nums)
}

func TestHeapSort(t *testing.T){
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := []int{0,1,2,3,4,5,6,7,8,9}
	heapSort(nums)
	if ! reflect.DeepEqual(nums,want) {
		t.Errorf("got %v, want %v", nums, want)
	}
	t.Log(nums)
}