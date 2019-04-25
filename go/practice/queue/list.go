package queue

import (
	"github.com/mgxian/algo/go/practice/list"
)

// ListQueue represents a list queue.
type ListQueue struct {
	l        list.List
	capacity int
}

// NewListQueue creates a new initial list queue.
func NewListQueue(n int) *ListQueue {
	if n < 1 {
		return nil
	}
	return &ListQueue{
		l:        list.NewSinglyLinkedList(),
		capacity: n,
	}
}

// Enqueue add element to end of queue.
func (q *ListQueue) Enqueue(e interface{}) interface{} {
	if q.IsFull() {
		return nil
	}

	if result := q.l.PushBack(e); result != nil {
		return e
	}

	return nil
}

// Dequeue remove element from front of queue.
func (q *ListQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	if result := q.l.Front(); result != nil {
		q.l.Remove(result)
		return result.Value
	}

	return nil
}

// IsEmpty check if queue is empty.
func (q *ListQueue) IsEmpty() bool {
	return q.l.Len() == 0
}

// IsFull check if queue is full.
func (q *ListQueue) IsFull() bool {
	return q.l.Len() == q.capacity
}

// Peek get the value of the front of queue without removing it.
func (q *ListQueue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	if result := q.l.Front(); result != nil {
		return result.Value
	}

	return nil
}
