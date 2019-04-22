package list

// CircularDoublyLinkedList represents a singly linked list.
type CircularDoublyLinkedList struct {
	head Element // sentinel list element
	len  int     // current list length excluding (this) sentinel element
}

// NewCircularDoublyLinkedList create a new initial CircularDoublyLinkedList
func NewCircularDoublyLinkedList() *CircularDoublyLinkedList {
	l := new(CircularDoublyLinkedList)
	l.head.prev = &l.head
	l.head.next = &l.head
	l.len = 0
	return l
}

// Front returns the first element of the linked list or nil if the linked list is empty.
func (l *CircularDoublyLinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// Back returns the last element of the linked list or nil if the linked list is empty.
func (l *CircularDoublyLinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.prev
}

// Len returns the number of elements of the linked list.
func (l *CircularDoublyLinkedList) Len() int {
	return l.len
}

func (l *CircularDoublyLinkedList) insert(e, at *Element) *Element {
	n := at.next
	at.next = e
	e.prev = at

	e.next = n
	n.prev = e

	l.len++
	return e
}

func (l *CircularDoublyLinkedList) remove(e *Element) *Element {
	e.prev.next = e.next
	e.next.prev = e.prev

	l.len--
	return e
}

func (l *CircularDoublyLinkedList) move(e, at *Element) *Element {
	if e == at {
		return e
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	n := at.next
	at.next = e
	e.prev = at

	e.next = n
	n.prev = e

	return e
}

func (l *CircularDoublyLinkedList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// InsertAfter inserts a new element e with value v after mark and returns e.
func (l *CircularDoublyLinkedList) InsertAfter(v interface{}, mark *Element) *Element {
	return l.insertValue(v, mark)
}

// InsertBefore inserts a new element e with value v before mark and returns e.
func (l *CircularDoublyLinkedList) InsertBefore(v interface{}, mark *Element) *Element {
	return l.insertValue(v, mark.prev)
}

// MoveToBack moves element e to the back of the linked list.
func (l *CircularDoublyLinkedList) MoveToBack(e *Element) {
	if l.head.prev == e {
		return
	}
	l.move(e, l.head.prev)
}

// MoveToFront moves element e to the front of the linked list.
func (l *CircularDoublyLinkedList) MoveToFront(e *Element) {
	if l.head.next == e {
		return
	}
	l.move(e, &l.head)
}

// PushBack inserts a new element e with value v to the back of the linked list.
func (l *CircularDoublyLinkedList) PushBack(v interface{}) *Element {
	return l.insertValue(v, l.head.prev)
}

// PushFront inserts a new element e with value v to the front of the linked list.
func (l *CircularDoublyLinkedList) PushFront(v interface{}) *Element {
	return l.insertValue(v, &l.head)
}

// Remove remove element e from the linked list and returns the element value.
func (l *CircularDoublyLinkedList) Remove(e *Element) interface{} {
	l.remove(e)
	return e.Value
}
