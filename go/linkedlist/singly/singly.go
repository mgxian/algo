package singly

// ListNode linkedlist node
type ListNode struct {
	val  interface{}
	Next *ListNode
}

// LinkedList linkedlist
type LinkedList struct {
	Head *ListNode
}

func newListNode(val interface{}) *ListNode {
	return &ListNode{
		val:  val,
		Next: nil,
	}
}

// NewLinkedList create a linkedlist
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// InsertToHead insert a node to linkedlist head
func (l *LinkedList) InsertToHead(val interface{}) {
	node := newListNode(val)
	node.Next = l.Head
	l.Head = node
}

// InsertToTail insert a node to linkedlist tail
func (l *LinkedList) InsertToTail(val interface{}) {
	node := newListNode(val)
	p := l.Head

	if p == nil {
		l.Head = node
		return
	}

	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
}

// InsertBefore insert val to linkedlist before a exist val
func (l *LinkedList) InsertBefore(existVal, val interface{}) {
	p := l.Head
	if p == nil {
		return
	}

	if l.Head.val == existVal {
		l.InsertToHead(val)
		return
	}

	for p.Next != nil {
		if p.Next.val == existVal {
			node := newListNode(val)
			node.Next = p.Next
			p.Next = node
			return
		}
		p = p.Next
	}
}

// InsertAfter insert val to linkedlist after a exist val
func (l *LinkedList) InsertAfter(existVal, val interface{}) {
	p := l.Head
	if p == nil {
		return
	}

	for p != nil {
		if p.val == existVal {
			node := newListNode(val)
			node.Next = p.Next
			p.Next = node
		}
		p = p.Next
	}
}

// Delete delete all val from linkedlist
func (l *LinkedList) Delete(val interface{}) {
	if l.Head == nil {
		return
	}

	for {
		if l.Head.val == val {
			l.Head = l.Head.Next
			continue
		}
		break
	}

	p := l.Head
	for p.Next != nil {
		if p.Next.val == val {
			p.Next = p.Next.Next
			continue
		}
		p = p.Next
	}
}

// Find find if val in linkedlist
func (l LinkedList) Find(val interface{}) bool {
	if l.Head == nil {
		return false
	}

	p := l.Head

	for p != nil {
		if p.val == val {
			return true
		}
		p = p.Next
	}

	return false
}

// Traverse print all val in linkedlist
func (l LinkedList) Traverse() (ret []interface{}) {
	p := l.Head
	if p == nil {
		return
	}

	for p != nil {
		ret = append(ret, p.val)
		p = p.Next
	}

	return
}
