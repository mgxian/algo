package sort

func insertSort(data sortable) {
	for i := 1; i < data.Len(); i++ {
		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}
