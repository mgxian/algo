package queue

// ArrayQueue represents a array queue.
type ArrayQueue struct {
	capacity int
	front    int
	rear     int
	data     []interface{}
}

// NewArrayQueue create a initial array queue.
func NewArrayQueue(n int) *ArrayQueue {
	if n < 1 {
		return nil
	}
	return &ArrayQueue{
		capacity: n,
		front:    -1,
		rear:     -1,
		data:     make([]interface{}, n),
	}
}

// Enqueue add element to end of queue.
func (q *ArrayQueue) Enqueue(e interface{}) interface{} {
	if q.IsFull() {
		return nil
	}

	q.rear++
	q.data[q.rear] = e
	return e
}

// Dequeue remove element from front of queue.
func (q *ArrayQueue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	q.front++
	result := q.data[q.front]
	if q.front == q.rear {
		q.front = -1
		q.rear = -1
	}
	return result
}

// IsEmpty check if queue is empty.
func (q *ArrayQueue) IsEmpty() bool {
	return q.rear-q.front == 0
}

// IsFull check if queue is full.
func (q *ArrayQueue) IsFull() bool {
	return q.rear == q.capacity-1
}

// Peek get the value of the front of queue without removing it.
func (q *ArrayQueue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}

	return q.data[q.front+1]
}
