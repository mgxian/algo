package queue

import "errors"

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
		size:  size,
		front: 0,
		rear:  0,
	}
}

func (q *SimpleQueue) enQueue(val interface{}) (err error) {
	if q.isFull() {
		err = errors.New("queue is full")
		return
	}

	q.data[q.rear] = val
	q.rear = (q.rear + 1) % q.size
	return
}

func (q *SimpleQueue) deQueue() (val interface{}, err error) {
	if q.isEmpty() {
		err = errors.New("queue is empty")
		return
	}
	val = q.data[q.front]
	q.front = (q.front + 1) % q.size
	return
}

func (q SimpleQueue) isEmpty() (empty bool) {
	if q.rear == q.front {
		empty = true
	}
	return
}

func (q SimpleQueue) isFull() (full bool) {
	if (q.rear+1)%q.size == q.front {
		full = true
	}
	return
}
