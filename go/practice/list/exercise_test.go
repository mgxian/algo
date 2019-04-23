package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergeSortedSinglyLinkedList(t *testing.T) {
	Convey("Test singly linked list mergeSortedSinglyLinkedList", t, func() {
		Convey("Setup", func() {
			lessFn := func(x, y interface{}) bool { return x.(int) < y.(int) }
			l1 := NewSinglyLinkedList()
			l2 := NewSinglyLinkedList()

			Convey("Test two empty singly linked list", func() {
				l3 := mergeSortedSinglyLinkedList(l1, l2, lessFn)
				So(checkSinglyLinkedListPointers(l3, []*Element{}), ShouldEqual, true)
			})

			Convey("Test one empty singly linked list", func() {
				l3 := mergeSortedSinglyLinkedList(l1, l2, lessFn)
				So(checkSinglyLinkedListPointers(l3, []*Element{}), ShouldEqual, true)
			})

			Convey("Test two single element singly linked list", func() {
				e1 := l1.PushBack(1)
				e2 := l2.PushBack(2)
				l3 := mergeSortedSinglyLinkedList(l1, l2, lessFn)
				So(checkSinglyLinkedListPointers(l3, []*Element{e1, e2}), ShouldEqual, true)
			})

			Convey("Test two more elements singly linked list", func() {
				e1 := l1.PushBack(1)
				e3 := l1.PushBack(3)
				e5 := l1.PushBack(5)

				e2 := l2.PushBack(2)
				e4 := l2.PushBack(4)
				e6 := l2.PushBack(6)

				l3 := mergeSortedSinglyLinkedList(l1, l2, lessFn)
				t.Log(l3)
				So(checkSinglyLinkedListPointers(l3, []*Element{e1, e2, e3, e4, e5, e6}), ShouldEqual, true)
			})

			Convey("Test two more elements and have equal element singly linked list", func() {
				e1 := l1.PushBack(1)
				e3 := l1.PushBack(3)
				e5 := l1.PushBack(5)

				e2 := l2.PushBack(2)
				e4 := l2.PushBack(3)
				e6 := l2.PushBack(6)

				l3 := mergeSortedSinglyLinkedList(l1, l2, lessFn)
				t.Log(l3)
				So(checkSinglyLinkedListPointers(l3, []*Element{e1, e2, e3, e4, e5, e6}), ShouldEqual, true)
			})
		})
	})
}
