package list

// Element is an element of a linked list.
type Element struct {
	prev, next *Element
	Value      interface{}
}

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element {
	return e.prev
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element {
	return e.next
}

// List is a interface that a linked list must implements.
type List interface {
	// Front returns the first element of linked list or nil if the linked list is empty.
	Front() *Element
	// Back returns the last element of linked list or nil if the linked list is empty.
	Back() *Element
	// Len returns the number of elements of the linked list.
	Len() int
	// InsertAfter inserts a new element e with value v after mark and returns e.
	InsertAfter(v interface{}, mark *Element) *Element
	// InsertBefore inserts a new element e with value v before mark and returns e.
	InsertBefore(v interface{}, mark *Element) *Element
	// MoveToBack moves element e to the back of the linked list.
	MoveToBack(e *Element)
	// MoveToFront moves element e to the front of the linked list.
	MoveToFront(e *Element)
	// PushBack inserts a new element e with value v to the back of the linked list.
	PushBack(v interface{}) *Element
	// PushFront inserts a new element e with value v to the front of the linked list.
	PushFront(v interface{}) *Element
	// Remove remove element e from the linked list and returns the element value.
	Remove(e *Element)
}
