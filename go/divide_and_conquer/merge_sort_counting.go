package divide

func countReversePairs(data []int) int {
	n := mergeSortCounting(data, 0, len(data)-1)
	return n
}

func mergeSortCounting(data []int, start, end int) int {
	if start >= end {
		return 0
	}

	mid := start + (end-start)/2
	a := mergeSortCounting(data, start, mid)
	b := mergeSortCounting(data, mid+1, end)
	n := merge(data, start, mid, end)

	return a + b + n
}

func merge(data []int, start, mid, end int) int {
	i := start
	j := mid + 1
	k := 0
	n := 0
	tmp := make([]int, end-start+1)
	for i <= mid && j <= end {
		if data[i] <= data[j] {
			tmp[k] = data[i]
			i++
		} else {
			n += mid - i + 1
			tmp[k] = data[j]
			j++
		}
		k++
	}

	for i <= mid {
		tmp[k] = data[i]
		i++
		k++
	}

	for j <= end {
		tmp[k] = data[j]
		j++
		k++
	}

	k = 0
	for i := start; i <= end; i++ {
		data[i] = tmp[k]
		k++
	}

	return n
}
