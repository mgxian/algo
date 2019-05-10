package sort

func bubbleSort(data sortable) {
	for i := 0; i < data.Len()-1; i++ {
		for j := data.Len() - 1; j > i; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
