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

func main() {
	linkedList := singly.NewLinkedList()
	numbers := []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		linkedList.InsertToTail(num)
	}

	reverseList(linkedList)
	fmt.Println(linkedList.Traverse())
}
