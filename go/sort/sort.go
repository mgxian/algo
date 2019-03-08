package sort

// Element is the wrap of the need sorted data
type Element struct{}

// Sortable is the wrap of the Less
type Sortable interface {
	Less(e Element) bool
}

type MyInt struct {
	Element
	val int
}

func NewMyInt(val int) MyInt {
	return MyInt{
		val: val,
	}
}

func (m MyInt) Less(aMyInt MyInt) bool {
	if m.val < aMyInt.val {
		return true
	}

	return false
}

func BubbleSort(data []Sortable) {
	length := len(data)
	for i := 0; i < length-1; i++ {
		hasSwap := false
		for j := length - 1; j > i; j-- {
			if data[j].Less(data[j-1]) {
				hasSwap = true
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
		if !hasSwap {
			break
		}
	}
}
