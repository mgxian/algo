package singly

// LinkedListWithDummyHead  linkedlist with dummy head
type LinkedListWithDummyHead struct {
	DummyHead *ListNode
}

// NewLinkedListWithDummyHead create linkedlist with dummy head
func NewLinkedListWithDummyHead() *LinkedListWithDummyHead {
	dummyHead := NewListNode(0)
	return &LinkedListWithDummyHead{
		DummyHead: dummyHead,
	}
}

// InsertToTail insert a val to list tail
func (l *LinkedListWithDummyHead) InsertToTail(val interface{}) {
	p := l.DummyHead
	for p.Next != nil {
		p = p.Next
	}
	p.Next = NewListNode(val)
}

// InsertToHead insert a val to list head
func (l *LinkedListWithDummyHead) InsertToHead(val interface{}) {
	p := l.DummyHead
	node := NewListNode(val)
	node.Next = p.Next
	p.Next = node
}

// InsertBefore insert a val to list before a val
func (l *LinkedListWithDummyHead) InsertBefore(existVal, val interface{}) {
	p := l.DummyHead
	for p.Next != nil {
		if p.Next.Val == existVal {
			node := NewListNode(val)
			node.Next = p.Next
			p.Next = node
			return
		}
		p = p.Next
	}

}

// InsertAfter insert a val to list after a val
func (l *LinkedListWithDummyHead) InsertAfter(existVal, val interface{}) {
	p := l.DummyHead.Next
	for p != nil {
		if p.Val == val {
			node := NewListNode(val)
			node.Next = p.Next
			p.Next = node
			return
		}
		p = p.Next
	}
}

// Delete delete a val from list
func (l *LinkedListWithDummyHead) Delete(val interface{}) {
	p := l.DummyHead
	for p.Next != nil {
		if p.Next.Val == val {
			p.Next = p.Next.Next
			return
		}
		p = p.Next
	}
}

// Find find if a val in the list
func (l LinkedListWithDummyHead) Find(val interface{}) bool {
	p := l.DummyHead.Next
	for p != nil {
		if p.Val == val {
			return true
		}
		p = p.Next
	}
	return false
}

// Tranverse return all val in list
func (l LinkedListWithDummyHead) Tranverse() (ret []interface{}) {
	p := l.DummyHead.Next
	for p != nil {
		ret = append(ret, p.Val)
		p = p.Next
	}
	return
}

// Size get the size of the list
func (l LinkedListWithDummyHead) Size() (n int) {
	p := l.DummyHead.Next
	for p != nil {
		p = p.Next
		n++
	}
	return
}

// DeleteTail delete last val from the list
func (l *LinkedListWithDummyHead) DeleteTail() {
	p := l.DummyHead
	for p.Next != nil {
		if p.Next.Next == nil {
			break
		}
		p = p.Next
	}
	p.Next = p.Next.Next
}
