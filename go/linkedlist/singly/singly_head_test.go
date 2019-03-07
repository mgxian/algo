package singly

import (
	"fmt"
	"testing"
)

func TestListInsertToHead(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}

	linkedList := NewLinkedList()

	for _, num := range numbers {
		linkedList.InsertToHead(num)
	}

	actual := linkedList.Traverse()
	for i, n := range expected {
		if n != actual[i] {
			t.Error(fmt.Sprintf("Expected the linkedlist of %v to be %d but instead got %d!", numbers, expected, actual))
		}
	}
}

func TestListInsertToTail(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	linkedList := NewLinkedList()

	for _, num := range numbers {
		linkedList.InsertToTail(num)
	}

	actual := linkedList.Traverse()
	for i, n := range expected {
		if n != actual[i] {
			t.Error(fmt.Sprintf("Expected the linkedlist of %v to be %d but instead got %d!", numbers, expected, actual))
		}
	}
}

func TestListFind(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := map[int]bool{3: true, 6: false}

	linkedList := NewLinkedList()

	for _, num := range numbers {
		linkedList.InsertToTail(num)
	}

	for n, exist := range expected {
		actual := linkedList.Find(n)
		if exist != actual {
			t.Error(fmt.Sprintf("Expected the linkedlist of %v to be %v but instead got %v!", numbers, expected, actual))
		}
	}
}

func TestListDelete(t *testing.T) {
	numbers := []int{1, 1, 2, 3, 4, 3, 5, 5}
	deleted := []int{1, 3, 5}
	expected := []int{2, 4}

	linkedList := NewLinkedList()

	for _, num := range numbers {
		linkedList.InsertToTail(num)
	}

	for _, d := range deleted {
		linkedList.Delete(d)
	}

	actual := linkedList.Traverse()
	for i, n := range expected {
		if n != actual[i] {
			t.Error(fmt.Sprintf("Expected the linkedlist of %v to be %d but instead got %d!", numbers, expected, actual))
		}
	}
}

func TestListInsertBefore(t *testing.T) {
	numbers := []int{3, 4, 5}
	existVal := 3
	inserted := []int{1, 2}
	expected := []int{1, 2, 3, 4, 5}

	linkedList := NewLinkedList()

	for _, num := range numbers {
		linkedList.InsertToTail(num)
	}

	for _, val := range inserted {
		linkedList.InsertBefore(existVal, val)
	}

	actual := linkedList.Traverse()
	for i, n := range expected {
		if n != actual[i] {
			t.Error(fmt.Sprintf("Expected the linkedlist of %v to be %d but instead got %d!", numbers, expected, actual))
		}
	}
}

func TestListInsertAfter(t *testing.T) {
	numbers := []int{1, 3, 5}
	inserted := map[int]int{1: 2, 3: 4, 5: 6}
	expected := []int{1, 2, 3, 4, 5, 6}

	linkedList := NewLinkedList()

	for _, num := range numbers {
		linkedList.InsertToTail(num)
	}

	for existVal, val := range inserted {
		linkedList.InsertAfter(existVal, val)
	}

	actual := linkedList.Traverse()
	for i, n := range expected {
		if n != actual[i] {
			t.Error(fmt.Sprintf("Expected the linkedlist of %v to be %d but instead got %d!", numbers, expected, actual))
		}
	}
}
