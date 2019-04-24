package queue

// Queue is a interface that a queue must implements.
type Queue interface {
	// InQueue inqueue a element to the tail of queue
	InQueue(e interface{}) interface{}
	DeQueue() interface{}
	IsEmpty() bool
	IsFull() bool
}