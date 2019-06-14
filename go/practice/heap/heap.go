package heap

func parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func leftChild(i int) int {
	return 2*i + 1
}

func rightChild(i int) int {
	return 2*i + 2
}

func isLess(data []int, p, c int) bool {
	if c < len(data) && data[p] < data[c] {
		return true
	}
	return false
}

func isGreater(data []int, p, c int) bool {
	if c < len(data) && data[p] > data[c] {
		return true
	}
	return false
}

type compareFunc func([]int, int, int) bool

func isHeap(data []int, fn compareFunc) bool {
	if len(data) < 2 {
		return true
	}

	for p := 0; p < len(data); p++ {
		if fn(data, p, leftChild(p)) {
			return false
		}

		if fn(data, p, rightChild(p)) {
			return false
		}
	}

	return true
}

func isMaxHeap(data []int) bool {
	return isHeap(data, isLess)
}

func isMinHeap(data []int) bool {
	return isHeap(data, isGreater)
}

func heapifyFromBottomToTop(data []int, i int) {
	if i == 0 {
		return
	}

	for {
		p := parent(i)
		if p < 0 {
			break
		}
		if data[p] < data[i] {
			data[p], data[i] = data[i], data[p]
		}
		i = p
	}
}

func heapifyFromTopToBottom(data []int, n, i int) {
	for {
		max := i
		left := leftChild(i)
		right := rightChild(i)
		if left < n && data[left] > data[i] {
			max = left
		}

		if right < n && data[right] > data[max] {
			max = right
		}

		if max == i {
			break
		}

		data[i], data[max] = data[max], data[i]
		i = max
	}
}

func heap(data []int) {
	for i := 0; i < len(data); i++ {
		heapifyFromBottomToTop(data, i)
	}
}

func buildHeap(data []int) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		heapifyFromTopToBottom(data, len(data), i)
	}
}

func heapSort(data []int) {
	buildHeap(data)
	length := len(data)
	k := length - 1
	for k > 0 {
		data[k], data[0] = data[0], data[k]
		k--
		heapifyFromTopToBottom(data, k+1, 0)
	}
}
