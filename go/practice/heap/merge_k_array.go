package heap

import (
	"errors"
)

type arrayValue struct {
	v     int
	index int
}

type minHeap struct {
	capacity int
	size     int
	data     []arrayValue
}

func newMinHeap(n int) *minHeap {
	mh := new(minHeap)
	mh.data = make([]arrayValue, n)
	mh.capacity = n
	return mh
}

func (mh *minHeap) minParentAndChildren(i int) int {
	min := i
	left := 2*i + 1
	if left < mh.size && mh.data[left].v < mh.data[min].v {
		min = left
	}

	right := 2*i + 2
	if right < mh.size && mh.data[right].v < mh.data[min].v {
		min = right
	}

	return min
}

func (mh *minHeap) heapifyFromBottom(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if mh.data[p].v <= mh.data[i].v {
			break
		}
		mh.data[p], mh.data[i] = mh.data[i], mh.data[p]
		i = p
	}
}

func (mh *minHeap) heapifyFromTop(i int) {
	if i < 0 {
		return
	}

	for {
		min := mh.minParentAndChildren(i)
		if min == i {
			break
		}
		mh.data[i], mh.data[min] = mh.data[min], mh.data[i]
		i = min
	}
}

func (mh *minHeap) insert(v arrayValue) error {
	if mh.size >= mh.capacity {
		return errors.New("heap is full")
	}
	mh.data[mh.size] = v
	mh.heapifyFromBottom(mh.size)
	mh.size++
	return nil
}

func (mh *minHeap) top() (arrayValue, error) {
	if mh.size <= 0 {
		return arrayValue{}, errors.New("heap is empty")
	}
	v := mh.data[0]
	mh.data[0] = mh.data[mh.size-1]
	mh.size--
	mh.heapifyFromTop(0)
	return v, nil
}

func mergeKSortedArray(numbersToMerge ...[]int) []int {
	result := make([]int, 0)
	k := len(numbersToMerge)
	mh := newMinHeap(k)
	arrayCursors := make([]int, k)
	for i, nums := range numbersToMerge {
		v := arrayValue{nums[arrayCursors[i]], i}
		mh.insert(v)
		arrayCursors[i]++
	}

	for {
		v, err := mh.top()
		// fmt.Println("top :", v)
		// fmt.Printf("top after %v \n\n", mh.data)
		if err != nil {
			break
		}
		result = append(result, v.v)
		idx := arrayCursors[v.index]
		if idx > len(numbersToMerge[v.index])-1 {
			continue
		}
		iv := numbersToMerge[v.index][idx]
		nv := arrayValue{iv, v.index}
		mh.insert(nv)
		arrayCursors[v.index]++
		// fmt.Println("insert :", nv)
		// fmt.Printf("insert after %v \n\n", mh.data)
	}

	return result
}
