package queue

type RingQueue struct {
	capacity int // number of elements
	size     int // capacity + 1, data slice size, ring queue waste one postion.
	front    int
	rear     int
	data     []interface{}
}

func NewRingQueue(n int) *RingQueue {
	return &RingQueue{
		capacity: n,
		size:     n + 1,
		front:    0,
		rear:     0,
		data:     make([]interface{}, n+1),
	}
}

// Enqueue add element to end of queue.
func (q *RingQueue) Enqueue(e interface{}) interface{} {
	if q.IsFull() {
		return nil
	}
	q.data[q.rear] = e
	q.rear = (q.rear + 1) % q.size
	return e
}

// Dequeue remove element from front of queue.
func (q *RingQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	e := q.data[q.front]
	q.front = (q.front + 1) % q.size
	return e
}

// IsEmpty check if queue is empty.
func (q *RingQueue) IsEmpty() bool {
	return q.front == q.rear
}

// IsFull check if queue is full.
func (q *RingQueue) IsFull() bool {
	return (q.rear+1)%q.size == q.front
}

// Peek get the value of the front of queue without removing it.
func (q *RingQueue) Peek() interface{} {}
