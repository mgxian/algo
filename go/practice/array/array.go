package array

// Value is array element data type
type Value interface{}

// DynamicInterface is dynamic array interface
type DynamicInterface interface {
	Len() int
	Capacity() int

	Insert(int, Value) error
	InsertFront(Value) error
	InsertBack(Value) error

	Remove(int) (Value, error)
	RemoveFront() (Value, error)
	RemoveBack() (Value, error)

	Get(int) (Value, error)
	Find(Value, func(Value, Value) bool) int
	FindAll(Value, func(Value, Value) bool) []int
	Contains(Value, func(Value, Value) bool) bool

	Set(int, Value) error

	String() string
}

// SortedArrayInterface is sorted array interface
type SortedArrayInterface interface {
	Len() int
	Capacity() int
	Add(Value) error
	Remove(Value) error
	RemoveAll(Value) (int, error)
	Replace(Value, Value) error
	Find(Value) int
	FindAll(Value) []int
	String() string
}
