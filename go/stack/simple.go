package stack

import "errors"

// SimpleStack simple array based stack
type SimpleStack struct {
	data []interface{}
	size int
	top  int
}

// NewSimpleStack create a simple array based stack
func NewSimpleStack(size int) *SimpleStack {
	return &SimpleStack{
		data: make([]interface{}, size),
		size: size,
		top:  -1,
	}
}

// Push save a val to the stack
func (s *SimpleStack) Push(val interface{}) (err error) {
	if s.isFull() {
		err = errors.New("stack is full")
		return
	}
	s.top++
	s.data[s.top] = val
	return
}

// Pop delete a val from the stack
func (s *SimpleStack) Pop() (val interface{}, err error) {
	if s.isEmpty() {
		err = errors.New("stack is empty")
		return
	}
	val = s.data[s.top]
	s.top--
	return
}

// Top get the stack top val
func (s SimpleStack) Top() (val interface{}, err error) {
	if s.isEmpty() {
		err = errors.New("stack is empty")
		return
	}
	val = s.data[s.top]
	return
}

// isEmpty is the stack is empty
func (s SimpleStack) isEmpty() (empty bool) {
	if s.top == -1 {
		return true
	}
	return
}

// isFull is the stack is full
func (s SimpleStack) isFull() (full bool) {
	if s.top == s.size-1 {
		return true
	}
	return
}
