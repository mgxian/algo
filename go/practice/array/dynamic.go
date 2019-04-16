package array

import (
	"bytes"
	"errors"
	"fmt"
)

// DynamicArray is a dynamic grow array
type DynamicArray struct {
	capacity int
	data     []Value
	len      int
}

// NewDynamicArray create a new DynamicArray
func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		capacity: capacity,
		data:     make([]Value, capacity),
	}
}

// Len returns the number of elements in the DynamicArray
func (d *DynamicArray) Len() int {
	return d.len
}

// Capacity returns the capacity fo the DynamicArray
func (d *DynamicArray) Capacity() int {
	return d.capacity
}

func (d *DynamicArray) resize(capacity int) {
	data := make([]Value, capacity)
	for i, v := range d.data {
		data[i] = v
	}
	d.capacity = capacity
	d.data = data
}

// Insert inserts the specified element at the specified position in the DynamicArray
func (d *DynamicArray) Insert(i int, v Value) error {
	if i < 0 || i > d.Len() {
		return errors.New("index out of range")
	}

	if d.Len() == d.Capacity() {
		d.resize(d.Capacity() * 2)
	}

	for j := d.Len(); j > i; j-- {
		d.data[j] = d.data[j-1]
	}

	d.data[i] = v
	d.len++

	return nil
}

// InsertFront inserts the specified element at the front of the DynamicArray
func (d *DynamicArray) InsertFront(v Value) error {
	return d.Insert(0, v)
}

// InsertBack inserts the specified element at the back of the DynamicArray
func (d *DynamicArray) InsertBack(v Value) error {
	return d.Insert(d.Len(), v)
}

// Remove delete the element at the specified position in the DynamicArray
func (d *DynamicArray) Remove(i int) (Value, error) {
	if i < 0 || i >= d.Len() {
		return nil, errors.New("index out of range")
	}

	v := d.data[i]
	for j := i; j < d.Len()-1; j++ {
		d.data[j] = d.data[j+1]
	}

	d.len--
	return v, nil
}

// RemoveFront delete the element at the front of the DynamicArray
func (d *DynamicArray) RemoveFront() (Value, error) {
	return d.Remove(0)
}

// RemoveBack delete the element at the back of the DynamicArray
func (d *DynamicArray) RemoveBack() (Value, error) {
	return d.Remove(d.Len() - 1)
}

// Get returns the element at the specified position in the DynamicArray
func (d *DynamicArray) Get(i int) (Value, error) {
	if i < 0 || i >= d.Len() {
		return nil, errors.New("index out of range")
	}
	return d.data[i], nil
}

// Find returns the postion of first element that equal to specified element in the DynamicArray
func (d *DynamicArray) Find(v Value, equalFn func(x, y Value) bool) int {
	for i := 0; i < d.Len(); i++ {
		if equalFn(v, d.data[i]) {
			return i
		}
	}

	return -1
}

// FindAll returns the postion of elements that equal to specified element in the DynamicArray
func (d *DynamicArray) FindAll(v Value, equalFn func(x, y Value) bool) []int {
	result := make([]int, 0)
	for i := 0; i < d.Len(); i++ {
		if equalFn(v, d.data[i]) {
			result = append(result, i)
		}
	}

	return result
}

// Contains check if the DynamicArray has element than equal to the specified element
func (d *DynamicArray) Contains(v Value, equalFn func(x, y Value) bool) bool {
	if d.Find(v, equalFn) != -1 {
		return true
	}
	return false
}

// Set replaces the element at the specified position in the DynamicArray with the specified element.
func (d *DynamicArray) Set(i int, v Value) error {
	if i < 0 || i >= d.Len() {
		return errors.New("index out of range")
	}
	d.data[i] = v
	return nil
}

func (d *DynamicArray) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", d.Len(), d.Capacity()))
	buffer.WriteString("[")
	for i := 0; i < d.Len(); i++ {
		buffer.WriteString(fmt.Sprint(d.data[i]))
		if i != d.Len()-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
