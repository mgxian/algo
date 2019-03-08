package queue

import (
	"errors"
	"fmt"
)

// ListNode is the linkedlist's node
type ListNode struct {
	val  interface{}
	next *ListNode
}

// LinkedListQueue is the queue based on linkedlist
type LinkedListQueue struct {
	front  *ListNode
	rear   *ListNode
	length int
	size   int
}

func newListNode(val interface{}) *ListNode {
	return &ListNode{
		val: val,
	}
}

// NewLinkedListQueue create a new LinkedListQueue
func NewLinkedListQueue(size int) *LinkedListQueue {
	return &LinkedListQueue{
		size: size,
	}
}

// IsFull checks if the queue is full
func (q LinkedListQueue) IsFull() (full bool) {
	if q.length >= q.size {
		full = true
	}
	return
}

// IsEmpty checks if the queue if empty
func (q LinkedListQueue) IsEmpty() (empty bool) {
	if q.length <= 0 || q.front == nil {
		empty = true
	}
	return
}

// EnQueue add a item to the queue
func (q *LinkedListQueue) EnQueue(val interface{}) (err error) {
	if q.IsFull() {
		err = errors.New("queue is full")
		return
	}

	node := newListNode(val)
	if q.IsEmpty() {
		q.rear = node
		q.front = node
	} else {
		q.rear.next = node
		q.rear = node
	}

	q.length++
	return
}

// DeQueue remove a item from the queue
func (q *LinkedListQueue) DeQueue() (val interface{}, err error) {
	if q.IsEmpty() {
		err = errors.New("queue is empty")
		return
	}

	val = q.front.val
	q.front = q.front.next

	if q.front == nil {
		q.rear = nil
	}

	q.length--
	return
}

// Peek gets the elements at the front of the queue without removing it
func (q LinkedListQueue) Peek() (val interface{}, err error) {
	if q.IsEmpty() {
		err = errors.New("queue is empty")
		return
	}

	val = q.front.val
	return
}

func (q LinkedListQueue) String() (ret string) {
	if q.IsEmpty() {
		ret = "queue is empty"
		return
	}

	start := q.front
	ret = "queue:"
	for start != nil {
		ret += fmt.Sprintf(" <- %v", start.val)
		start = start.next
	}
	return
}
