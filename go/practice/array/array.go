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
