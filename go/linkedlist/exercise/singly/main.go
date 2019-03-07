package main

import (
	"fmt"

	"github.com/mgxian/algo/go/linkedlist/singly"
)

func reverseList(l *singly.LinkedList) {
	p := l.Head

	cur := p.Next
	p.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = p
		p = cur
		cur = next
	}

	l.Head = p
}

func middleNode(l *singly.LinkedList) *singly.ListNode {
	p := l.Head
	slow, fast := p, p
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == nil {
			break
		}
	}

	return slow
}

func hasCycle(l *singly.LinkedList) bool {
	p := l.Head
	slow, fast := p, p
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == nil {
			break
		}

		if slow == fast {
			return true
		}
	}

	return false
}

func removeNthFromEnd(l *singly.LinkedList, n int) {
	p := l.Head
	slow, fast := p, p
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		l.Head = p.Next
		return
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next

		if fast == nil {
			break
		}
	}

	slow.Next = slow.Next.Next
}

func mergeTwoLists(l1, l2 *singly.LinkedList) (l3 *singly.LinkedList) {
	p, q := l1.Head, l2.Head
	l3 = singly.NewLinkedList()
	var cur *singly.ListNode
	for p != nil && q != nil {
		if p.Val.(int) <= q.Val.(int) {
			if l3.Head == nil {
				l3.Head = p
			} else {
				cur.Next = p
			}
			cur = p
			p = p.Next
		} else {
			if l3.Head == nil {
				l3.Head = q
			} else {
				cur.Next = q
			}
			cur = q
			q = q.Next
		}
	}

	if p != nil {
		cur.Next = p
	} else {
		cur.Next = q
	}

	return
}

func generateLinkedList(numbers []int) *singly.LinkedList {
	// generate a linkedlist
	linkedList := singly.NewLinkedList()

	for _, num := range numbers {
		linkedList.InsertToTail(num)
	}

	return linkedList
}

func testLinkedList() {
	var linkedList *singly.LinkedList

	// reverse linkedlist
	linkedList = generateLinkedList([]int{1, 2, 3, 4, 5})
	reverseList(linkedList)
	fmt.Println(linkedList.Traverse())

	// get the middle node of the linkedlist
	linkedList = generateLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Printf("middleNode: %v \n", middleNode(linkedList))

	// if linkedlist has a cycle
	linkedList = generateLinkedList([]int{1, 2, 3, 4, 5})
	fmt.Println(hasCycle(linkedList))
	forthListNode := linkedList.Head.Next.Next.Next
	forthListNode.Next = linkedList.Head.Next
	fmt.Println(hasCycle(linkedList))

	// remove Nth from end
	linkedList = generateLinkedList([]int{1, 2, 3, 4, 5})
	removeNthFromEnd(linkedList, 1)
	fmt.Println(linkedList.Traverse())
	removeNthFromEnd(linkedList, 4)
	fmt.Println(linkedList.Traverse())
	removeNthFromEnd(linkedList, 2)
	fmt.Println(linkedList.Traverse())

	// merge two linkedlist
	l1 := generateLinkedList([]int{1, 2, 4})
	l2 := generateLinkedList([]int{1, 3, 4})
	fmt.Println(mergeTwoLists(l1, l2).Traverse())
}

func reverseListWithDummyHead(l *singly.LinkedListWithDummyHead) {
	dummyHead := l.DummyHead
	p := dummyHead.Next
	cur := p.Next
	p.Next = nil
	for cur != nil {
		next := cur.Next
		cur.Next = dummyHead.Next
		dummyHead.Next = cur
		cur = next
	}
}

func middleNodeListWithDummyHead(l *singly.LinkedListWithDummyHead) *singly.ListNode {
	p := l.DummyHead.Next
	slow, fast := p, p

	for fast != nil {
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func hasCycleListWithDummyHead(l *singly.LinkedListWithDummyHead) bool {
	p := l.DummyHead.Next
	slow, fast := p, p

	for fast != nil {
		if fast.Next == nil {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func removeNthFromEndListWithDummyHead(l *singly.LinkedListWithDummyHead, n int) {
	p := l.DummyHead.Next
	slow, fast := p, p
	for fast != nil && n > 0 {
		fast = fast.Next
		n--
	}

	for fast != nil {
		fast = fast.Next
		if fast == nil {
			break
		}
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

}

func mergeTwoListsWithDummyHead(l1, l2 *singly.LinkedListWithDummyHead) (l3 *singly.LinkedListWithDummyHead) {
	p, q := l1.DummyHead.Next, l2.DummyHead.Next
	l3 = singly.NewLinkedListWithDummyHead()
	cur := l3.DummyHead

	for p != nil && q != nil {
		if p.Val.(int) <= q.Val.(int) {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}

	if p != nil {
		cur.Next = p
	} else {
		cur.Next = q
	}

	return
}

func generateLinkedListWithDummyHead(numbers []int) *singly.LinkedListWithDummyHead {
	linkedListWithDummyHead := singly.NewLinkedListWithDummyHead()

	for _, num := range numbers {
		linkedListWithDummyHead.InsertToTail(num)
	}

	return linkedListWithDummyHead
}

func testLinkedListWithDummyHead() {
	var linkedListWithDummyHead *singly.LinkedListWithDummyHead
	linkedListWithDummyHead = generateLinkedListWithDummyHead([]int{1, 2, 3, 4, 5})
	reverseListWithDummyHead(linkedListWithDummyHead)
	fmt.Println(linkedListWithDummyHead.Tranverse())

	linkedListWithDummyHead = generateLinkedListWithDummyHead([]int{1, 2, 3, 4, 5})
	fmt.Println(middleNodeListWithDummyHead(linkedListWithDummyHead))
	linkedListWithDummyHead = generateLinkedListWithDummyHead([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(middleNodeListWithDummyHead(linkedListWithDummyHead))

	linkedListWithDummyHead = generateLinkedListWithDummyHead([]int{1, 2, 3, 4, 5})
	fmt.Println(hasCycleListWithDummyHead(linkedListWithDummyHead))
	fifthListNode := linkedListWithDummyHead.DummyHead.Next.Next.Next.Next.Next
	fifthListNode.Next = linkedListWithDummyHead.DummyHead.Next.Next
	fmt.Println(hasCycleListWithDummyHead(linkedListWithDummyHead))

	linkedListWithDummyHead = generateLinkedListWithDummyHead([]int{1, 2, 3, 4, 5})
	removeNthFromEndListWithDummyHead(linkedListWithDummyHead, 5)
	fmt.Println(linkedListWithDummyHead.Tranverse())
	removeNthFromEndListWithDummyHead(linkedListWithDummyHead, 1)
	fmt.Println(linkedListWithDummyHead.Tranverse())
	removeNthFromEndListWithDummyHead(linkedListWithDummyHead, 2)
	fmt.Println(linkedListWithDummyHead.Tranverse())

	l1 := generateLinkedListWithDummyHead([]int{1, 2, 4})
	l2 := generateLinkedListWithDummyHead([]int{1, 3, 4})
	fmt.Println(mergeTwoListsWithDummyHead(l1, l2).Tranverse())
}

func main() {
	testLinkedList()
	// testLinkedListWithDummyHead()
}
