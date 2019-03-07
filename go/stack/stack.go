package stack

// Stack stack interface define
type Stack interface {
	Push(val interface{}) (err error)
	Pop() (val interface{}, err error)
	Top() (val interface{}, err error)
	isEmpty() (empty bool)
	isFull() (full bool)
}
