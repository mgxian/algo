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
	l.head.next = &l.head
	l.len = 0
	return l
}
