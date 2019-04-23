package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func checkCircularDoublyLinkedListLen(l *CircularDoublyLinkedList, es []*Element) bool {
	return l.Len() == len(es)
}

func checkCircularDoublyLinkedListNextPointers(l *CircularDoublyLinkedList, es []*Element) bool {
	prev := &l.head
	for _, e := range es {
		if prev.next != e {
			return false
		}
		prev = prev.next
	}

	back := es[len(es)-1]
	if back.next != &l.head {
		return false
	}

	return true
}

func checkCircularDoublyLinkedListPrevPointers(l *CircularDoublyLinkedList, es []*Element) bool {
	prev := &l.head
	for _, e := range es {
		if e.prev != prev {
			return false
		}
		prev = prev.next
	}

	back := es[len(es)-1]
	if l.head.prev != back {
		return false
	}

	return true
}

func checkEmptyCircularDoublyLinkedListPointers(l *CircularDoublyLinkedList) bool {
	if (l.head.next != nil && l.head.next != &l.head) || (l.head.prev != nil && l.head.prev != &l.head) {
		return false
	}
	return true
}

func checkCircularDoublyLinkedListPointers(l *CircularDoublyLinkedList, es []*Element) bool {
	if !checkCircularDoublyLinkedListLen(l, es) {
		return false
	}

	if len(es) == 0 {
		return checkEmptyCircularDoublyLinkedListPointers(l)
	}

	if !checkCircularDoublyLinkedListNextPointers(l, es) {
		return false
	}

	if !checkCircularDoublyLinkedListPrevPointers(l, es) {
		return false
	}

	return true
}

func TestCircularDoublyLinkedListPushBack(t *testing.T) {
	Convey("Test circular doubly linked list PushBack", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test empty circular doubly linked list", func() {
				So(checkCircularDoublyLinkedListPointers(l, []*Element{}), ShouldEqual, true)
			})

			Convey("Test push back single element circular doubly linked list", func() {
				e := l.PushBack(1)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test push back more elements circular doubly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e1, e2, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestCircularDoublyLinkedListPushFront(t *testing.T) {
	Convey("Test circular doubly linked list PushFront", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test push front single element circular doubly linked list", func() {
				e := l.PushFront(1)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test push front more elements circular doubly linked list", func() {
				e1 := l.PushFront(1)
				e2 := l.PushFront(2)
				e3 := l.PushFront(3)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestCircularDoublyLinkedListFrontBack(t *testing.T) {
	Convey("Test circular doubly linked list Front and Back", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test empty circular doubly linked list", func() {
				So(l.Front(), ShouldBeNil)
				So(l.Back(), ShouldBeNil)
			})

			Convey("Test push front single element circular doubly linked list", func() {
				e := l.PushFront(1)
				So(l.Front(), ShouldEqual, e)
				So(l.Back(), ShouldEqual, e)
			})

			Convey("Test push back single element circular doubly linked list", func() {
				e := l.PushBack(1)
				So(l.Front(), ShouldEqual, e)
				So(l.Back(), ShouldEqual, e)
			})

			Convey("Test push front more elements circular doubly linked list", func() {
				e1 := l.PushFront(1)
				l.PushFront(2)
				e3 := l.PushFront(3)
				So(l.Front(), ShouldEqual, e3)
				So(l.Back(), ShouldEqual, e1)
			})

			Convey("Test push back more elements circular doubly linked list", func() {
				e1 := l.PushBack(1)
				l.PushBack(2)
				e3 := l.PushBack(3)
				So(l.Front(), ShouldEqual, e1)
				So(l.Back(), ShouldEqual, e3)
			})

			Convey("Test push front and back more elements circular doubly linked list", func() {
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

func TestCircularDoublyLinkedListInsertAfter(t *testing.T) {
	Convey("Test circular doubly linked list InsertAfter", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test InsertAfter single element circular doubly linked list", func() {
				e := l.InsertAfter(1, &l.head)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test InsertAfter more elements circular doubly linked list", func() {
				e1 := l.InsertAfter(1, &l.head)
				e2 := l.InsertAfter(2, e1)
				e3 := l.InsertAfter(3, e2)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e1, e2, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestCircularDoublyLinkedListInsertBefore(t *testing.T) {
	Convey("Test circular doubly linked list InsertBefore", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test InsertBefore single element circular doubly linked list", func() {
				e := l.InsertBefore(1, &l.head)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test InsertBefore more elements circular doubly linked list", func() {
				e1 := l.InsertAfter(1, &l.head)
				e2 := l.InsertBefore(2, e1)
				e3 := l.InsertBefore(3, e2)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestCircularDoublyLinkedListMoveToBack(t *testing.T) {
	Convey("Test circular doubly linked list MoveToBack", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test single element circular doubly linked list", func() {
				e := l.PushBack(1)
				l.MoveToBack(e)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test more elements circular doubly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.MoveToBack(e1)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e2, e3, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestCircularDoublyLinkedListMoveToFront(t *testing.T) {
	Convey("Test circular doubly linked list MoveToFront", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test single element circular doubly linked list", func() {
				e := l.PushBack(1)
				l.MoveToFront(e)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test more elements circular doubly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.MoveToFront(e3)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e3, e1, e2}), ShouldEqual, true)
			})
		})
	})
}

func TestCircularDoublyLinkedListRemove(t *testing.T) {
	Convey("Test circular doubly linked list Remove", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test single element circular doubly linked list", func() {
				e := l.PushBack(1)
				l.Remove(e)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{}), ShouldEqual, true)
			})

			Convey("Test more elements circular doubly linked list remove first element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e1)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e2, e3}), ShouldEqual, true)
			})

			Convey("Test more elements circular doubly linked list remove last element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e3)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e1, e2}), ShouldEqual, true)
			})

			Convey("Test more elements circular doubly linked list remove middle element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e2)
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e1, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestCircularDoublyLinkedListReverse(t *testing.T) {
	Convey("Test circular doubly linked list Reverse", t, func() {
		Convey("Setup", func() {
			l := NewCircularDoublyLinkedList()

			Convey("Test reverse single element circular doubly linked list", func() {
				e := l.PushBack(1)
				l.Reverse()
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test reverse more elements circular doubly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Reverse()
				So(checkCircularDoublyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}
