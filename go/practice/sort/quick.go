package sort

func quickSort(data sortable, start, end int) {
	if start >= end {
		return
	}

	q := partition(data, start, end)
	quickSort(data, start, q-1)
	quickSort(data, q+1, end)
}

func partition(data sortable, start, end int) int {
	pivot := end
	i := start
	for j := start; j < end; j++ {
		if data.Less(j, pivot) {
			data.Swap(i, j)
			i++
		}
	}
	data.Swap(i, end)
	return i
}
