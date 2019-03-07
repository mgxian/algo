package queue

// Queue is an interface used by queue
type Queue interface {
	enQueue(val interface{}) (err error)
	deQueue() (val interface{}, err error)
	isEmpty() (empty bool)
	isFull() (full bool)
}
