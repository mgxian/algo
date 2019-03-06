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

func main() {
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
