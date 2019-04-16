package array

type SortedArray struct {
	capacity int
	data     []Value
	len      int
	lessFn   func(Value, Value) bool
}

func NewSortedArray(capacity int, lessFn func(Value, Value) bool) *SortedArray {
	return &SortedArray{
		capacity: capacity,
		data:     make([]Value, capacity),
		lessFn:   lessFn,
	}
}

func (s *SortedArray) Len() int {
	return s.len
}

func (s *SortedArray) Capacity() int {
	return s.capacity
}

func (s *SortedArray) Find(Value, func(Value, Value) bool) int {

}

func (s *SortedArray) Add(v Value) error {
}

func (s *SortedArray) Remove(v Value) error          {}
func (s *SortedArray) Replace(src, dest Value) error {}
func (s *SortedArray) String() string                {}
