package array

func mergeSortedArray(arr1, arr2 []Value, m, n int, lessFn compareFn) []Value {
	i, j, k := 0, 0, 0
	result := make([]Value, m+n)
	for i < m && j < n {
		if !lessFn(arr2[j], arr1[i]) {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}

	if i < m {
		for i < m {
			result[k] = arr1[i]
			i++
			k++
		}
	} else {
		for j < n {
			result[k] = arr2[j]
			j++
			k++
		}
	}

	return result
}
