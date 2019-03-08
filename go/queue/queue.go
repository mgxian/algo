package queue

// Queue is an interface used by queue
type Queue interface {
	EnQueue(val interface{}) (err error)
	DeQueue() (val interface{}, err error)
	IsEmpty() (empty bool)
	IsFull() (full bool)
	Peek() (val interface{}, err error)
}
