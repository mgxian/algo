package list

// CircularLinkedList represents a singly linked list.
type CircularLinkedList struct {
	head Element // sentinel list element
	len  int     // current list length excluding (this) sentinel element
}

// NewCircularLinkedList create a new initial CircularLinkedList
func NewCircularLinkedList() *CircularLinkedList {
	l := new(CircularLinkedList)
	l.head.prev = &l.head
	l.head.next = &l.head
	l.len = 0
	return l
}
