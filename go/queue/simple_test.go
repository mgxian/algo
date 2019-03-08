package queue

import "testing"

func TestSimpleQueueEnqueue(t *testing.T) {
	var aQueue Queue
	aQueue = NewSimpleQueue(2)
	aQueue.EnQueue(1)
	aQueue.EnQueue(2)
	if err := aQueue.EnQueue(3); err == nil {
		t.Error("EnQueue should fail, but it succeed")
	}

	val, _ := aQueue.Peek()
	if val.(int) != 1 {
		t.Errorf("expected 1, but got %d", val)
	}

	t.Log(aQueue)
}

func TestSimpleQueueDequeue(t *testing.T) {
	var aQueue Queue
	aQueue = NewSimpleQueue(2)
	aQueue.EnQueue(1)
	aQueue.EnQueue(2)

	t.Log(aQueue)

	val, _ := aQueue.DeQueue()
	if val.(int) != 1 {
		t.Errorf("expected 1, but got %d", val)
	}

	t.Log(aQueue)

	val, _ = aQueue.DeQueue()
	if val.(int) != 2 {
		t.Errorf("expected 2, but got %d", val)
	}

	t.Log(aQueue)
}
