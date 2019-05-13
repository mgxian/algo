package sort

func partition2(data sortable, start, end int) int {
	pivot := end
	i := start
	for j := start; j < end; j++ {
		if data.Less(j, pivot) {
			data.Swap(i, j)
			i++
		}
	}
	data.Swap(i, pivot)
	return i
}

func findKth(data sortable, k int) int {
	rightIdx := data.Len() - k
	start := 0
	end := data.Len() - 1
	for {
		p := partition2(data, start, end)
		if p == rightIdx {
			return p
		}

		if p < rightIdx {
			start = p + 1
		} else {
			end = p - 1
		}
	}
}
