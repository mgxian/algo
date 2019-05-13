package bsearch

func binarySearch(data []int, num int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > num {
			high = mid - 1
		} else if data[mid] < num {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func binarySearchFirstEqual(data []int, num int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > num {
			high = mid - 1
		} else if data[mid] < num {
			low = mid + 1
		} else {
			if mid == 0 || data[mid-1] != num {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func binarySearchLastEqual(data []int, num int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > num {
			high = mid - 1
		} else if data[mid] < num {
			low = mid + 1
		} else {
			if mid == len(data)-1 || data[mid+1] != num {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}

func binarySearchFirstEqualOrGreaterThan(data []int, num int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] < num {
			low = mid + 1
		} else {
			if mid == 0 || data[mid-1] < num {
				return mid
			}
			high = mid - 1
		}
	}
	return -1
}

func binarySearchLastEqualOrLessThan(data []int, num int) int {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := low + (high-low)/2
		if data[mid] > num {
			high = mid - 1
		} else {
			if mid == len(data)-1 || data[mid+1] > num {
				return mid
			}
			low = mid + 1
		}
	}
	return -1
}
