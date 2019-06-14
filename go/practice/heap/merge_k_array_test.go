package heap

import (
	"reflect"
	"testing"
)

func TestMinHeap(t *testing.T) {
	mh := newMinHeap(10)
	want := true
	for i := 9; i >= 0; i-- {
		mh.insert(arrayValue{i, -1})
		vs := make([]int, 0)
		for _, v := range mh.data[:mh.size] {
			vs = append(vs, v.v)
		}
		got := isMinHeap(vs)
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestMergeKSortedArray(t *testing.T) {
	tests := []struct {
		numbersToMerge [][]int
		want           []int
	}{
		{
			[][]int{
				[]int{2, 5, 8, 9},
			},
			[]int{2, 5, 8, 9},
		},
		{
			[][]int{
				[]int{1, 4},
				[]int{3, 6, 7},
			},
			[]int{1, 3, 4, 6, 7},
		},
		{
			[][]int{
				[]int{1, 4},
				[]int{3, 6, 7},
				[]int{2, 5, 8, 9},
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		got := mergeKSortedArray(tt.numbersToMerge...)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
