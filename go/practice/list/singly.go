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
