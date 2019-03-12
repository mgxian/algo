package sort

import (
	"math"
	"reflect"
)

// Sortable is interface that collections that want sort need implement
type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type MyInts []int

func (p MyInts) Len() int           { return len(p) }
func (p MyInts) Less(i, j int) bool { return p[i] < p[j] }
func (p MyInts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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

// InsertionSort is a implement of insertion sort
func InsertionSort(data Sortable) {
	length := data.Len()
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}

// SelectionSort is a implement of select sort
func SelectionSort(data Sortable) {
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
	for j := start; j < end; j++ {
		if data.Less(j, pivot) {
			data.Swap(j, i)
			i++
		}
	}
	data.Swap(i, pivot)

	QuickSort(data, start, i-1)
	QuickSort(data, i+1, end)
}

// BucketSort is a implement of bucket sort
func BucketSort(data []int) {
	min, max := 0, 0
	length := len(data)
	for i := 1; i < length; i++ {
		if data[i] < data[min] {
			min = i
			continue
		}
		if data[max] < data[i] {
			max = i
		}
	}

	bucketSize := 10
	bucketCount := int((data[max]-data[min])/bucketSize) + 1
	bucket := make([][]int, bucketCount)
	for i := 0; i < length; i++ {
		bucketNum := int((data[i] - data[min]) / bucketSize)
		bucket[bucketNum] = append(bucket[bucketNum], data[i])
	}

	index := 0
	for i := 0; i < bucketCount; i++ {
		InsertionSort(MyInts(bucket[i]))
		for _, v := range bucket[i] {
			data[index] = v
			index++
		}
	}
}

// CountingSort is a implement of counting sort
func CountingSort(data []int) {
	max := 0
	length := len(data)
	for i := 1; i < length; i++ {
		if data[max] < data[i] {
			max = i
		}
	}

	count := make([]int, data[max]+1)
	for i := 0; i < length; i++ {
		count[data[i]]++
	}

	for i := 1; i < data[max]+1; i++ {
		count[i] = count[i-1] + count[i]
	}

	tmp := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		index := count[data[i]] - 1
		tmp[index] = data[i]
		count[data[i]]--
	}

	for i := 0; i < length; i++ {
		data[i] = tmp[i]
	}
}

// RadixSort is a implement of radix sort
func RadixSort(data []int) {
	max := data[0]
	length := len(data)
	for i := 1; i < length; i++ {
		if max < data[i] {
			max = data[i]
		}
	}

	maxDigitalNumber := 1
	for (max / 10) > 0 {
		maxDigitalNumber++
		max = max / 10
	}

	for n := 1; n <= maxDigitalNumber; n++ {
		count := make([]int, 10)
		for i := 0; i < length; i++ {
			tmpDigital := data[i]/int(math.Pow10(n-1)) - (data[i]/int(math.Pow10(n)))*10
			count[tmpDigital]++
		}

		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}

		tmp := make([]int, length)
		for i := length - 1; i >= 0; i-- {
			tmpDigital := data[i]/int(math.Pow10(n-1)) - (data[i]/int(math.Pow10(n)))*10
			index := count[tmpDigital] - 1
			tmp[index] = data[i]
			count[tmpDigital]--
		}

		for i := 0; i < length; i++ {
			data[i] = tmp[i]
		}
	}
}
