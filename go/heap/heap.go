package heap

import "fmt"

// Heap is heap data structure
type Heap struct {
	data  []int
	count int
	size  int
}

// NewHeap create a new heap
func NewHeap(size int) *Heap {
	return &Heap{
		data: make([]int, size),
		size: size,
	}
}

func getParentIndex(i int) int {
	return (i - 1) / 2
}

// Insert save a item to the heap
func (h *Heap) Insert(val int) {
	if h.count >= h.size {
		return
	}

	h.data[h.count] = val
	currentIndex := h.count
	parentIndex := getParentIndex(h.count)
	for parentIndex >= 0 && h.data[currentIndex] > h.data[parentIndex] {
		h.data[parentIndex], h.data[currentIndex] = h.data[currentIndex], h.data[parentIndex]
		currentIndex = parentIndex
		parentIndex = getParentIndex(currentIndex)
	}

	h.count++
}

// RemoveMax remove a item from the top of the heap
func (h *Heap) RemoveMax() int {
	if h.count == 0 {
		return -1
	}

	ret := h.data[0]
	h.data[0] = h.data[h.count-1]
	h.count--
	h.heapify()
	return ret
}

func (h *Heap) heapify() {
	current := 0
	for {
		left := 2*current + 1
		right := 2*current + 2
		maxIndex := current
		if left < h.count && h.data[maxIndex] < h.data[left] {
			maxIndex = left
		}

		if right < h.count && h.data[maxIndex] < h.data[right] {
			maxIndex = right
		}

		if maxIndex == current {
			break
		}

		h.data[current], h.data[maxIndex] = h.data[maxIndex], h.data[current]
		current = maxIndex
	}
}

func (h *Heap) String() string {
	ret := "Heap: "
	for i := 0; i < h.count; i++ {
		ret += fmt.Sprintf("%d-->", h.data[i])
	}

	ret += "\n"
	return ret
}
