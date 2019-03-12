package search

// BinarySearch is a simple binary search
func BinarySearch(data []int, n int) int {
	length := len(data)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] == n {
			return mid
		}

		if data[mid] > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchFirstEqual find the first value that equal to the specified value from the slice return it's index
func BinarySearchFirstEqual(data []int, val int) int {
	length := len(data)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > val {
			high = mid - 1
		} else if data[mid] < val {
			low = mid + 1
		} else {
			if mid == 0 || data[mid-1] != val {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

// BinarySearchLastEqual find the last value that equal to the specified value from the slice return it's index
func BinarySearchLastEqual(data []int, val int) int {
	length := len(data)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > val {
			high = mid - 1
		} else if data[mid] < val {
			low = mid + 1
		} else {
			if mid == length-1 || data[mid+1] != val {
				return mid
			}
			low = mid + 1
		}
	}

	return -1
}

// BinarySearchFirstGreaterThanOrEqualTo find the first value that grater than or equal to the specified value from the slice return it's index
func BinarySearchFirstGreaterThanOrEqualTo(data []int, val int) int {
	length := len(data)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] < val {
			low = mid + 1
		} else {
			if mid == 0 || data[mid-1] < val {
				return mid
			}
			high = mid - 1
		}
	}

	return -1
}

// BinarySearchLastLessThanOrEqualTo find the last value that less than or equal to the specified value from the slice return it's index
func BinarySearchLastLessThanOrEqualTo(data []int, val int) int {
	length := len(data)
	low, high := 0, length-1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > val {
			high = mid - 1
		} else {
			if mid == length-1 || data[mid+1] > val {
				return mid
			}
			low = mid + 1
		}
	}

	return -1
}
