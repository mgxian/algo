package stack

// ArrayStack represents a array stack.
type ArrayStack struct {
	top      int
	capacity int
	data     []interface{}
}

// NewArrayStack create a new initial array stack of n elements.
func NewArrayStack(n int) *ArrayStack {
	if n < 1 {
		return nil
	}
	return &ArrayStack{
		top:      -1,
		capacity: n,
		data:     make([]interface{}, n),
	}
}

// Push add a element to top of stack.
func (s *ArrayStack) Push(e interface{}) interface{} {
	if s.IsFull() {
		return nil
	}
	s.top++
	s.data[s.top] = e
	return e
}

// Pop remove a element from top of stack.
func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty(){
		return nil
	}
	result := s.data[s.top]
	s.top--
	return result
}

// IsEmpty check if stack is empty.
func (s *ArrayStack) IsEmpty() bool {
	return s.top == -1
}

// IsFull check if stack is full.
func (s *ArrayStack) IsFull() bool {
	return s.top == s.capacity-1
}

// Peek get the value of the top element without removing it.
func (s *ArrayStack) Peek() interface{} {
	if s.IsEmpty(){
		return nil
	}
	return s.data[s.top]
}
