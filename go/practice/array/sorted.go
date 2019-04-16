package array

import (
	"bytes"
	"errors"
	"fmt"
)

type compareFn func(Value, Value) bool

// SortedArray is a sorted array
type SortedArray struct {
	capacity int
	data     []Value
	len      int
	lessFn   compareFn
	equalFn  compareFn
}

// NewSortedArray creates a new SortedArray
func NewSortedArray(capacity int, lessFn, equalFn compareFn) *SortedArray {
	return &SortedArray{
		capacity: capacity,
		data:     make([]Value, capacity),
		lessFn:   lessFn,
		equalFn:  equalFn,
	}
}

// Len returns number of the SortedArray
func (s *SortedArray) Len() int {
	return s.len
}

// Capacity returns the capacity of the SortedArray
func (s *SortedArray) Capacity() int {
	return s.capacity
}

// Find returns the postion of element that first equal to the specified element in the SortedArray
func (s *SortedArray) Find(v Value) int {
	for i := 0; i < s.Len(); i++ {
		if s.equalFn(s.data[i], v) {
			return i
		}
	}
	return -1
}

// FindAll returns the postions of element that equal to the specified element in the SortedArray
func (s *SortedArray) FindAll(v Value) []int {
	result := make([]int, 0)
	find := false
	for i := 0; i < s.Len(); i++ {
		if s.equalFn(s.data[i], v) {
			result = append(result, i)
			find = true
		} else if find {
			break
		}
	}
	return result
}

// Add save a item to the SortedArray
func (s *SortedArray) Add(v Value) error {
	if s.Len() >= s.Capacity() {
		return errors.New("sorted array is full")
	}

	pos := -1
	for i := 0; i < s.Len(); i++ {
		if !s.lessFn(s.data[i], v) {
			pos = i
			break
		}
	}

	if pos == -1 {
		s.data[s.Len()] = v
		s.len++
		return nil
	}

	for i := s.Len(); i > pos; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[pos] = v
	s.len++
	return nil
}

// Remove delete first the specified element from the SortedArray
func (s *SortedArray) Remove(v Value) error {
	pos := s.Find(v)
	if pos == -1 {
		return errors.New("value not exist in sorted array")
	}

	for i := pos; i < s.Len()-1; i++ {
		s.data[i] = s.data[i+1]
	}

	s.len--
	return nil
}

// RemoveAll delete all the specified element from the SortedArray
func (s *SortedArray) RemoveAll(v Value) (int, error) {
	pos := s.FindAll(v)
	if len(pos) == 0 {
		return 0, errors.New("value not exist in sorted array")
	}

	count := len(pos)
	for i := pos[0]; i < s.Len()-count; i++ {
		s.data[i] = s.data[i+count]
	}

	s.len -= count
	return count, nil
}

// Replace replace the src element as dest element in the SortedArray
func (s *SortedArray) Replace(src, dest Value) error {
	n, err := s.RemoveAll(src)
	if err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		err = s.Add(dest)
		if err != nil {
			break
		}
	}

	return err
}

// String returns the readable format string for the SortedArray
func (s *SortedArray) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", s.Len(), s.Capacity()))
	buffer.WriteString("[")
	for i := 0; i < s.Len(); i++ {
		buffer.WriteString(fmt.Sprint(s.data[i]))
		if i != s.Len()-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
