package sort

import (
	"reflect"
)

// Sortable is interface that collections that want sort need implement
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// BubbleSort is a implement of bubble sort
func BubbleSort(data Sortable) {
	length := data.Len()
	for i := 0; i < length-1; i++ {
		hasSwap := false
		for j := length - 1; j > i; j-- {
			if data.Less(j, j-1) {
				hasSwap = true
				data.Swap(j, j-1)
			}
		}
		if !hasSwap {
			break
		}
	}
}

// InsertSort is a implement of insertion sort
func InsertSort(data Sortable) {
	length := data.Len()
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

// SelectSort is a implement of select sort
func SelectSort(data Sortable) {
	length := data.Len()
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if data.Less(j, minIndex) {
				minIndex = j
			}
		}
		data.Swap(i, minIndex)
	}
}

// ShellSort is a implement of shell sort
func ShellSort(data Sortable) {
	length := data.Len()
	step := length / 2
	for step > 0 {
		for i := step; i < length; i += step {
			for j := i; j > 0; j -= step {
				if data.Less(j, j-step) {
					data.Swap(j, j-step)
				}
			}
		}
		step = step / 2
	}
}

// MergeSort is a implement of merge sort
func MergeSort(data Sortable, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2
	MergeSort(data, start, mid)
	MergeSort(data, mid+1, end)

	tmp := reflect.MakeSlice(reflect.TypeOf(data), 0, end-start+1)
	i, j := start, mid+1
	for i <= mid && j <= end {
		if data.Less(j, i) {
			tmp = reflect.Append(tmp, reflect.ValueOf(data).Index(j))
			j++
		} else {
			tmp = reflect.Append(tmp, reflect.ValueOf(data).Index(i))
			i++
		}
	}

	m, n := i, mid
	if j <= end {
		m, n = j, end
	}

	for i := m; i <= n; i++ {
		v := reflect.ValueOf(data).Index(i)
		tmp = reflect.Append(tmp, v)
	}

	for i := 0; i <= end-start; i++ {
		reflect.ValueOf(data).Index(start + i).Set(tmp.Index(i))
	}
}

// QuickSort is a implement of quick sort
func QuickSort(data Sortable, start, end int) {
	if start >= end {
		return
	}

	pivot := end
	i := start
	for j := start + 1; j <= end; j++ {
		if data.Less(j, pivot) {
			data.Swap(j, i)
			i++
		}
	}
}
