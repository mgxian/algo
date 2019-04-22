package list

// DoublyLinkedList represents a singly linked list.
type DoublyLinkedList struct {
	head Element // sentinel list element
	len  int     // current list length excluding (this) sentinel element
}

// NewDoublyLinkedList create a new initial DoublyLinkedList
func NewDoublyLinkedList() *DoublyLinkedList {
	l := new(DoublyLinkedList)
	l.head.prev = &l.head
	l.head.next = nil
	l.len = 0
	return l
}

// Front returns the first element of the linked list or nil if the linked list is empty.
func (l *DoublyLinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// Back returns the last element of the linked list or nil if the linked list is empty.
func (l *DoublyLinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.prev
}

// Len returns the number of elements of the linked list.
func (l *DoublyLinkedList) Len() int {
	return l.len
}

func (l *DoublyLinkedList) insert(e, at *Element) *Element {
	if at.next == nil {
		l.head.prev = e
	}

	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	if n != nil {
		n.prev = e
	}

	l.len++
	return e
}

func (l *DoublyLinkedList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

func (l *DoublyLinkedList) remove(e *Element) *Element {
	if e.next == nil {
		l.head.prev = e.prev
	}

	e.prev.next = e.next
	if e.next != nil {
		e.next.prev = e.prev
	}

	l.len--
	return e
}

func (l *DoublyLinkedList) move(e, at *Element) *Element {
	if e == at {
		return e
	}

	if at.next == nil {
		l.head.prev = e
	}

	if e.next == nil {
		l.head.prev = e.prev
	}

	e.prev.next = e.next
	if e.next != nil {
		e.next.prev = e.prev
	}

	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	if n != nil {
		n.prev = e
	}

	return e
}

// InsertAfter inserts a new element e with value v after mark and returns e.
func (l *DoublyLinkedList) InsertAfter(v interface{}, mark *Element) *Element {
	return l.insertValue(v, mark)
}

// InsertBefore inserts a new element e with value v before mark and returns e.
func (l *DoublyLinkedList) InsertBefore(v interface{}, mark *Element) *Element {
	return l.insertValue(v, mark.prev)
}

// MoveToBack moves element e to the back of the linked list.
func (l *DoublyLinkedList) MoveToBack(e *Element) {
	if l.head.prev == e {
		return
	}
	l.move(e, l.head.prev)
}

// MoveToFront moves element e to the front of the linked list.
func (l *DoublyLinkedList) MoveToFront(e *Element) {
	if l.head.next == e {
		return
	}
	l.move(e, &l.head)
}

// PushBack inserts a new element e with value v to the back of the linked list.
func (l *DoublyLinkedList) PushBack(v interface{}) *Element {
	return l.insertValue(v, l.head.prev)
}

// PushFront inserts a new element e with value v to the front of the linked list.
func (l *DoublyLinkedList) PushFront(v interface{}) *Element {
	return l.insertValue(v, &l.head)
}

// Remove remove element e from the linked list and returns the element value.
func (l *DoublyLinkedList) Remove(e *Element) interface{} {
	l.remove(e)
	return e.Value
}
