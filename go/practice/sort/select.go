package sort

func selectSort(data sortable) {
	for i := 0; i < data.Len()-1; i++ {
		min := i
		for j := i + 1; j < data.Len(); j++ {
			if data.Less(j, min) {
				min = j
			}
		}
		data.Swap(i, min)
	}
}
