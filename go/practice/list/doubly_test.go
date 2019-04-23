package list

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func checkDoublyLinkedListLen(l *DoublyLinkedList, es []*Element) bool {
	return l.Len() == len(es)
}

func checkDoublyLinkedListNextPointers(l *DoublyLinkedList, es []*Element) bool {
	prev := &l.head
	for _, e := range es {
		if prev.next != e {
			return false
		}
		prev = prev.next
	}
	back := es[len(es)-1]
	if back.next != nil {
		return false
	}
	return true
}

func checkDoublyLinkedListPrevPointers(l *DoublyLinkedList, es []*Element) bool {
	for i, e := range es[:len(es)-1] {
		if es[i+1].prev != e {
			fmt.Println("middle")
			return false
		}
	}
	front := es[0]
	back := es[len(es)-1]
	if front.prev != nil && front.prev != &l.head || l.head.prev != back {
		fmt.Println("head tail")
		fmt.Println(front.prev, l.head.prev, back)
		return false
	}
	return true
}

func checkEmptyDoublyLinkedListPointers(l *DoublyLinkedList) bool {
	if (l.head.next != nil && l.head.next != &l.head) || (l.head.prev != nil && l.head.prev != &l.head) {
		return false
	}
	return true
}

func checkDoublyLinkedListPointers(l *DoublyLinkedList, es []*Element) bool {
	if !checkDoublyLinkedListLen(l, es) {
		fmt.Println("len")
		return false
	}

	if len(es) == 0 {
		return checkEmptyDoublyLinkedListPointers(l)
	}

	if !checkDoublyLinkedListNextPointers(l, es) {
		fmt.Println("next")
		return false
	}

	if !checkDoublyLinkedListPrevPointers(l, es) {
		fmt.Println("prev")
		return false
	}

	return true
}

func TestDoublyLinkedListPushBack(t *testing.T) {
	Convey("Test doubly linked list PushBack", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test empty doubly linked list", func() {
				So(checkDoublyLinkedListPointers(l, []*Element{}), ShouldEqual, true)
			})

			Convey("Test push back single element doubly linked list", func() {
				e := l.PushBack(1)
				So(checkDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test push back more elements doubly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				So(checkDoublyLinkedListPointers(l, []*Element{e1, e2, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestDoublyLinkedListPushFront(t *testing.T) {
	Convey("Test doubly linked list PushFront", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test push front single element doubly linked list", func() {
				e := l.PushFront(1)
				So(checkDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test push front more elements doubly linked list", func() {
				e1 := l.PushFront(1)
				e2 := l.PushFront(2)
				e3 := l.PushFront(3)
				So(checkDoublyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestDoublyLinkedListFrontBack(t *testing.T) {
	Convey("Test doubly linked list Front and Back", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test empty doubly linked list", func() {
				So(l.Front(), ShouldBeNil)
				So(l.Back(), ShouldBeNil)
			})

			Convey("Test push front single element doubly linked list", func() {
				e := l.PushFront(1)
				So(l.Front(), ShouldEqual, e)
				So(l.Back(), ShouldEqual, e)
			})

			Convey("Test push back single element doubly linked list", func() {
				e := l.PushBack(1)
				So(l.Front(), ShouldEqual, e)
				So(l.Back(), ShouldEqual, e)
			})

			Convey("Test push front more elements doubly linked list", func() {
				e1 := l.PushFront(1)
				l.PushFront(2)
				e3 := l.PushFront(3)
				So(l.Front(), ShouldEqual, e3)
				So(l.Back(), ShouldEqual, e1)
			})

			Convey("Test push back more elements doubly linked list", func() {
				e1 := l.PushBack(1)
				l.PushBack(2)
				e3 := l.PushBack(3)
				So(l.Front(), ShouldEqual, e1)
				So(l.Back(), ShouldEqual, e3)
			})

			Convey("Test push front and back more elements doubly linked list", func() {
				l.PushFront(1)
				l.PushFront(2)
				l.PushFront(3)
				e4 := l.PushBack(4)
				e5 := l.PushFront(5)
				So(l.Front(), ShouldEqual, e5)
				So(l.Back(), ShouldEqual, e4)
			})
		})
	})
}

func TestDoublyLinkedListInsertAfter(t *testing.T) {
	Convey("Test doubly linked list InsertAfter", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test InsertAfter single element doubly linked list", func() {
				e := l.InsertAfter(1, &l.head)
				So(checkDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test InsertAfter more elements doubly linked list", func() {
				e1 := l.InsertAfter(1, &l.head)
				e2 := l.InsertAfter(2, e1)
				e3 := l.InsertAfter(3, e2)
				So(checkDoublyLinkedListPointers(l, []*Element{e1, e2, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestDoublyLinkedListInsertBefore(t *testing.T) {
	Convey("Test doubly linked list InsertBefore", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test InsertBefore single element doubly linked list", func() {
				e := l.InsertBefore(1, &l.head)
				So(checkDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test InsertBefore more elements doubly linked list", func() {
				e1 := l.InsertAfter(1, &l.head)
				e2 := l.InsertBefore(2, e1)
				e3 := l.InsertBefore(3, e2)
				So(checkDoublyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestDoublyLinkedListMoveToBack(t *testing.T) {
	Convey("Test doubly linked list MoveToBack", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test single element doubly linked list", func() {
				e := l.PushBack(1)
				l.MoveToBack(e)
				So(checkDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test more elements doubly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.MoveToBack(e1)
				So(checkDoublyLinkedListPointers(l, []*Element{e2, e3, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestDoublyLinkedListMoveToFront(t *testing.T) {
	Convey("Test doubly linked list MoveToFront", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test single element doubly linked list", func() {
				e := l.PushBack(1)
				l.MoveToFront(e)
				So(checkDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test more elements doubly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.MoveToFront(e3)
				So(checkDoublyLinkedListPointers(l, []*Element{e3, e1, e2}), ShouldEqual, true)
			})
		})
	})
}

func TestDoublyLinkedListRemove(t *testing.T) {
	Convey("Test doubly linked list Remove", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test single element doubly linked list", func() {
				e := l.PushBack(1)
				l.Remove(e)
				So(checkDoublyLinkedListPointers(l, []*Element{}), ShouldEqual, true)
			})

			Convey("Test more elements doubly linked list remove first element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e1)
				So(checkDoublyLinkedListPointers(l, []*Element{e2, e3}), ShouldEqual, true)
			})

			Convey("Test more elements doubly linked list remove last element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e3)
				So(checkDoublyLinkedListPointers(l, []*Element{e1, e2}), ShouldEqual, true)
			})

			Convey("Test more elements doubly linked list remove middle element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e2)
				So(checkDoublyLinkedListPointers(l, []*Element{e1, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestDoublyLinkedListReverse(t *testing.T) {
	Convey("Test doubly linked list Reverse", t, func() {
		Convey("Setup", func() {
			l := NewDoublyLinkedList()

			Convey("Test reverse single element doubly linked list", func() {
				e := l.PushBack(1)
				l.Reverse()
				So(checkDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test reverse more elements doubly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Reverse()
				So(checkDoublyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}
