package stack

import (
	"errors"

	"github.com/mgxian/algo/go/linkedlist/singly"
)

// LinkedListStack linkedlist based stack
type LinkedListStack struct {
	list *singly.LinkedListWithDummyHead
	size int
}

// NewLinkedListStack create a linkedlist based stack
func NewLinkedListStack(size int) *LinkedListStack {
	return &LinkedListStack{
		list: singly.NewLinkedListWithDummyHead(),
		size: size,
	}
}

// Push save a val to the stack
func (s LinkedListStack) Push(val interface{}) (err error) {
	if s.isFull() {
		err = errors.New("stack is full")
		return
	}

	s.list.InsertToHead(val)
	return
}

// Pop delete a val at top of the stack and delete the val
func (s LinkedListStack) Pop() (val interface{}, err error) {
	if s.isEmpty() {
		err = errors.New("stack is empty")
		return
	}

	val = s.list.DeleteHead()
	return
}

// Top get the val at top of the stack without delete the val
func (s LinkedListStack) Top() (val interface{}, err error) {
	if s.isEmpty() {
		err = errors.New("stack is empty")
		return
	}
	val = s.list.GetHead()
	return
}

// isEmpty if the stack is empty
func (s LinkedListStack) isEmpty() (empty bool) {
	if s.list.Size() == 0 {
		empty = true
	}
	return
}

// isFull if the stack is full
func (s LinkedListStack) isFull() (full bool) {
	if s.list.Size() == s.size {
		full = true
	}
	return
}
