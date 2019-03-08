package queue

import (
	"errors"
	"fmt"
)

// SimpleQueue is a slice based queue
type SimpleQueue struct {
	data  []interface{}
	size  int
	front int
	rear  int
}

// NewSimpleQueue creates a new SimpleQueue
func NewSimpleQueue(size int) *SimpleQueue {
	return &SimpleQueue{
		data:  make([]interface{}, size),
		size:  size + 1,
		front: 0,
		rear:  0,
	}
}

// EnQueue add a item to the queue
func (q *SimpleQueue) EnQueue(val interface{}) (err error) {
	if q.IsFull() {
		err = errors.New("queue is full")
		return
	}

	q.data[q.rear] = val
	q.rear = (q.rear + 1) % q.size
	return
}

// DeQueue remove a item from the queue
func (q *SimpleQueue) DeQueue() (val interface{}, err error) {
	if q.IsEmpty() {
		err = errors.New("queue is empty")
		return
	}
	val = q.data[q.front]
	q.front = (q.front + 1) % q.size
	return
}

// Peek gets the element at the front of the queue without removing it
func (q SimpleQueue) Peek() (val interface{}, err error) {
	if q.IsEmpty() {
		err = errors.New("queue is empty")
		return
	}

	val = q.data[q.front]
	return
}

// IsEmpty checks if the queue is empty
func (q SimpleQueue) IsEmpty() (empty bool) {
	if q.rear == q.front {
		empty = true
	}
	return
}

// IsFull checks if the queue is full
func (q SimpleQueue) IsFull() (full bool) {
	if (q.rear+1)%q.size == q.front {
		full = true
	}
	return
}

func (q SimpleQueue) String() (ret string) {
	start := q.front

	if q.IsEmpty() {
		ret = "queue is empty"
		return
	}

	ret = "queue:"
	for start != q.rear {
		ret += fmt.Sprintf(" <- %v", q.data[start])
		start = (start + 1) % q.size
	}
	return
}
