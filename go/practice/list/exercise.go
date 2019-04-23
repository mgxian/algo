package list

func mergeSortedSinglyLinkedList(l1, l2 *SinglyLinkedList, lessFn func(interface{}, interface{}) bool) *SinglyLinkedList {
	l3 := NewSinglyLinkedList()

	p1 := l1.head.next
	p2 := l2.head.next

	for p1 != nil && p2 != nil {
		if !lessFn(p2.Value, p1.Value) {
			n := p1.next
			l3.insert(p1, l3.head.prev)
			p1 = n
		} else {
			n := p2.next
			l3.insert(p2, l3.head.prev)
			p2 = n
		}
	}

	p := p1
	if p1 == nil {
		p = p2
	}
	for p != nil {
		n := p.next
		l3.insert(p, l3.head.prev)
		p = n
	}

	return l3
}
