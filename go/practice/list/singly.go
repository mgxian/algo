package list

// SinglyLinkedList represents a singly linked list.
type SinglyLinkedList struct {
	head Element // sentinel list element
	len  int     // current list length excluding (this) sentinel element
}

// NewSinglyLinkedList create a new initial SinglyLinkedList
func NewSinglyLinkedList() *SinglyLinkedList {
	l := new(SinglyLinkedList)
	l.head.prev = &l.head
	l.head.next = &l.head
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

// InsertAfter inserts a new element e with value v after mark and returns e.
func (l *SinglyLinkedList) InsertAfter(v interface{}, mark *Element) *Element {
}

// InsertBefore inserts a new element e with value v before mark and returns e.
func (l *SinglyLinkedList) InsertBefore(v interface{}, mark *Element) *Element {}

// MoveToBack moves element e to the back of the linked list.
func (l *SinglyLinkedList) MoveToBack(e *Element) {}

// MoveToFront moves element e to the front of the linked list.
func (l *SinglyLinkedList) MoveToFront(e *Element) {}

// PushBack inserts a new element e with value v to the back of the linked list.
func (l *SinglyLinkedList) PushBack(v interface{}) *Element {}

// PushFront inserts a new element e with value v to the front of the linked list.
func (l *SinglyLinkedList) PushFront(v interface{}) *Element {}

// Remove remove element e from the linked list and returns the element value.
func (l *SinglyLinkedList) Remove(e *Element) {}
