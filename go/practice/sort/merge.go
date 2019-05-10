package sort

func mergeSort(data []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2
	mergeSort(data, start, mid)
	mergeSort(data, mid+1, end)

	result := merge(data, start, end)

	for i := start; i <= end; i++ {
		data[i] = result[i-start]
	}
}

func merge(data []int, start int, end int) []int {
	mid := start + (end-start)/2
	i, j, k := start, mid+1, 0
	result := make([]int, end-start+1)

	for i <= mid && j <= end {
		if data[i] <= data[j] {
			result[k] = data[i]
			i++
		} else {
			result[k] = data[j]
			j++
		}
		k++
	}

	m, n := i, mid
	if i > mid {
		m, n = j, end
	}

	for i := m; i <= n; i++ {
		result[k] = data[i]
		k++
	}

	return result
}
