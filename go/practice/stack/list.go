package stack

import (
	"github.com/mgxian/algo/go/practice/list"
)

type ListStack struct {
	capacity int
	l        list.List
}

func NewListStack(n int) *ListStack {
	if n < 1 {
		return nil
	}
	return &ListStack{
		capacity: n,
		l:        list.NewSinglyLinkedList(),
	}
}

// Push add a element to top of stack
func (s *ListStack) Push(e interface{}) interface{} {
	if s.IsFull() {
		return nil
	}
	s.l.PushBack(e)
	return e
}

// Pop remove a element from top of stack
func (s *ListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	if result := s.l.Back(); result != nil {
		s.l.Remove(result)
		return result.Value
	}
	return nil
}

// IsEmpty check if stack is empty
func (s *ListStack) IsEmpty() bool {
	return s.l.Len() == 0
}

// IsFull check if stack is full
func (s *ListStack) IsFull() bool {
	return s.l.Len() == s.capacity
}

// Peek get the value of the top element without removing it
func (s *ListStack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	if result := s.l.Back(); result != nil {
		return result.Value
	}
	return nil
}
