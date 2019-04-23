package stack

// Stack is a interface that a stack must implements
type Stack interface {
	// Push add a element to top of stack
	Push(e interface{}) interface{}
	// Pop remove a element from top of stack
	Pop() interface{}
	// IsEmpty check if stack is empty
	IsEmpty() bool
	// IsFull check if stack is full
	IsFull() bool
	// Peek get the value of the top element without removing it
	Peek() interface{}
}
