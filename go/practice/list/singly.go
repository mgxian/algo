package list

// SinglyLinkedList represents a singly linked list.
type SinglyLinkedList struct {
	// sentinel list element
	// head.next is the first element of the list
	// head.pre is the last element of the list
	head Element
	len  int // current list length excluding (this) sentinel element
}

// NewSinglyLinkedList create a new initial SinglyLinkedList
func NewSinglyLinkedList() *SinglyLinkedList {
	l := new(SinglyLinkedList)
	l.head.prev = &l.head
	l.head.next = nil
	l.len = 0
	return l
}

// Front returns the first element of the linked list or nil if the linked list is empty.
func (l *SinglyLinkedList) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.next
}

// Back returns the last element of the linked list or nil if the linked list is empty.
func (l *SinglyLinkedList) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.head.prev
}

// Len returns the number of elements of the linked list.
func (l *SinglyLinkedList) Len() int {
	return l.len
}

// insert inserts e after at, increments l.len, and returns e.
func (l *SinglyLinkedList) insert(e, at *Element) *Element {
	if at.next == nil {
		l.head.prev = e
	}
	e.next = at.next
	at.next = e
	l.len++
	return e
}

// remove removes e from the linked list, decrements l.len, and return e.
func (l *SinglyLinkedList) remove(e *Element) *Element {
	p := &l.head
	for p.next != e {
		p = p.next
	}
	p.next = e.next
	if e.next == nil {
		l.head.prev = p
	}
	l.len--
	return e
}

// move moves e to next to at and returns e.
func (l *SinglyLinkedList) move(e, at *Element) *Element {
	if at.next == nil {
		l.head.prev = e
	}
	p := &l.head
	for p.next != e {
		p = p.next
	}
	p.next = e.next
	e.next = at.next
	at.next = e
	return e
}

// insertValue inserts element e with value v after at and returns e.
func (l *SinglyLinkedList) insertValue(v interface{}, at *Element) *Element {
	return l.insert(&Element{Value: v}, at)
}

// InsertAfter inserts a new element e with value v after mark and returns e.
func (l *SinglyLinkedList) InsertAfter(v interface{}, mark *Element) *Element {
	return l.insertValue(v, mark)
}

// InsertBefore inserts a new element e with value v before mark and returns e.
func (l *SinglyLinkedList) InsertBefore(v interface{}, mark *Element) *Element {
	p := &l.head
	for p.next != mark {
		p = p.next
	}
	return l.insertValue(v, p)
}

// MoveToBack moves element e to the back of the linked list.
func (l *SinglyLinkedList) MoveToBack(e *Element) {
	l.move(e, l.head.prev)
}

// MoveToFront moves element e to the front of the linked list.
func (l *SinglyLinkedList) MoveToFront(e *Element) {
	l.move(e, &l.head)
}

// PushBack inserts a new element e with value v to the back of the linked list.
func (l *SinglyLinkedList) PushBack(v interface{}) *Element {
	return l.insertValue(v, l.head.prev)
}

// PushFront inserts a new element e with value v to the front of the linked list.
func (l *SinglyLinkedList) PushFront(v interface{}) *Element {
	return l.insertValue(v, &l.head)
}

// Remove remove element e from the linked list and returns the element value.
func (l *SinglyLinkedList) Remove(e *Element) interface{} {
	l.remove(e)
	return e.Value
}
