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
