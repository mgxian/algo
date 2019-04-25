package queue

// Queue is a interface that a queue must implements.
type Queue interface {
	// Enqueue add element to end of queue.
	Enqueue(e interface{}) interface{}
	// Dequeue remove element from front of queue.
	Dequeue() interface{}
	// IsEmpty check if queue is empty.
	IsEmpty() bool
	// IsFull check if queue is full.
	IsFull() bool
	// Peek get the value of the front of queue without removing it.
	Peek() interface{}
}
