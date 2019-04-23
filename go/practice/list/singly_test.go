package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func checkSinglyLinkedListLen(l *SinglyLinkedList, es []*Element) bool {
	return l.Len() == len(es)
}

func checkSinglyLinkedListNextPointers(l *SinglyLinkedList, es []*Element) bool {
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

func checkSinglyLinkedListPrevPointers(l *SinglyLinkedList, es []*Element) bool {
	back := es[len(es)-1]
	if l.head.prev != back {
		return false
	}

	return true
}

func checkEmptySinglyLinkedListPointers(l *SinglyLinkedList) bool {
	if (l.head.next != nil && l.head.next != &l.head) || (l.head.prev != nil && l.head.prev != &l.head) {
		return false
	}
	return true
}

func checkSinglyLinkedListPointers(l *SinglyLinkedList, es []*Element) bool {
	if !checkSinglyLinkedListLen(l, es) {
		return false
	}

	if len(es) == 0 {
		return checkEmptySinglyLinkedListPointers(l)
	}

	if !checkSinglyLinkedListNextPointers(l, es) {
		return false
	}

	if !checkSinglyLinkedListPrevPointers(l, es) {
		return false
	}

	return true
}

func TestSinglyLinkedListPushBack(t *testing.T) {
	Convey("Test singly linked list PushBack", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test empty singly linked list", func() {
				So(checkSinglyLinkedListPointers(l, []*Element{}), ShouldEqual, true)
			})

			Convey("Test push back single element singly linked list", func() {
				e := l.PushBack(1)
				So(checkSinglyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test push back more elements singly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				So(checkSinglyLinkedListPointers(l, []*Element{e1, e2, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestSinglyLinkedListPushFront(t *testing.T) {
	Convey("Test singly linked list PushFront", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test push front single element singly linked list", func() {
				e := l.PushFront(1)
				So(checkSinglyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test push front more elements singly linked list", func() {
				e1 := l.PushFront(1)
				e2 := l.PushFront(2)
				e3 := l.PushFront(3)
				So(checkSinglyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestSinglyLinkedListFrontBack(t *testing.T) {
	Convey("Test singly linked list Front and Back", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test empty singly linked list", func() {
				So(l.Front(), ShouldBeNil)
				So(l.Back(), ShouldBeNil)
			})

			Convey("Test push front single element singly linked list", func() {
				e := l.PushFront(1)
				So(l.Front(), ShouldEqual, e)
				So(l.Back(), ShouldEqual, e)
			})

			Convey("Test push back single element singly linked list", func() {
				e := l.PushBack(1)
				So(l.Front(), ShouldEqual, e)
				So(l.Back(), ShouldEqual, e)
			})

			Convey("Test push front more elements singly linked list", func() {
				e1 := l.PushFront(1)
				l.PushFront(2)
				e3 := l.PushFront(3)
				So(l.Front(), ShouldEqual, e3)
				So(l.Back(), ShouldEqual, e1)
			})

			Convey("Test push back more elements singly linked list", func() {
				e1 := l.PushBack(1)
				l.PushBack(2)
				e3 := l.PushBack(3)
				So(l.Front(), ShouldEqual, e1)
				So(l.Back(), ShouldEqual, e3)
			})

			Convey("Test push front and back more elements singly linked list", func() {
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

func TestSinglyLinkedListInsertAfter(t *testing.T) {
	Convey("Test singly linked list InsertAfter", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test InsertAfter single element singly linked list", func() {
				e := l.InsertAfter(1, &l.head)
				So(checkSinglyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test InsertAfter more elements singly linked list", func() {
				e1 := l.InsertAfter(1, &l.head)
				e2 := l.InsertAfter(2, e1)
				e3 := l.InsertAfter(3, e2)
				So(checkSinglyLinkedListPointers(l, []*Element{e1, e2, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestSinglyLinkedListInsertBefore(t *testing.T) {
	Convey("Test singly linked list InsertBefore", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test InsertBefore single element singly linked list", func() {
				e2 := l.InsertAfter(2, &l.head)
				e1 := l.InsertBefore(1, e2)
				So(checkSinglyLinkedListPointers(l, []*Element{e1, e2}), ShouldEqual, true)
			})

			Convey("Test InsertBefore more elements singly linked list", func() {
				e1 := l.InsertAfter(1, &l.head)
				e2 := l.InsertBefore(2, e1)
				e3 := l.InsertBefore(3, e2)
				So(checkSinglyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestSinglyLinkedListMoveToBack(t *testing.T) {
	Convey("Test singly linked list MoveToBack", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test single element singly linked list", func() {
				e := l.PushBack(1)
				l.MoveToBack(e)
				So(checkSinglyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test more elements singly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.MoveToBack(e1)
				So(checkSinglyLinkedListPointers(l, []*Element{e2, e3, e1}), ShouldEqual, true)
			})
		})
	})
}

func TestSinglyLinkedListMoveToFront(t *testing.T) {
	Convey("Test singly linked list MoveToFront", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test single element singly linked list", func() {
				e := l.PushBack(1)
				l.MoveToFront(e)
				So(checkSinglyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test more elements singly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.MoveToFront(e3)
				So(checkSinglyLinkedListPointers(l, []*Element{e3, e1, e2}), ShouldEqual, true)
			})
		})
	})
}

func TestSinglyLinkedListRemove(t *testing.T) {
	Convey("Test singly linked list Remove", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test single element singly linked list", func() {
				e := l.PushBack(1)
				l.Remove(e)
				So(checkSinglyLinkedListPointers(l, []*Element{}), ShouldEqual, true)
			})

			Convey("Test more elements singly linked list remove first element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e1)
				So(checkSinglyLinkedListPointers(l, []*Element{e2, e3}), ShouldEqual, true)
			})

			Convey("Test more elements singly linked list remove last element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e3)
				So(checkSinglyLinkedListPointers(l, []*Element{e1, e2}), ShouldEqual, true)
			})

			Convey("Test more elements singly linked list remove middle element", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Remove(e2)
				So(checkSinglyLinkedListPointers(l, []*Element{e1, e3}), ShouldEqual, true)
			})
		})
	})
}

func TestSinglyLinkedListReverse(t *testing.T) {
	Convey("Test singly linked list Reverse", t, func() {
		Convey("Setup", func() {
			l := NewSinglyLinkedList()

			Convey("Test reverse single element singly linked list", func() {
				e := l.PushBack(1)
				l.Reverse()
				So(checkSinglyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
			})

			Convey("Test reverse more elements singly linked list", func() {
				e1 := l.PushBack(1)
				e2 := l.PushBack(2)
				e3 := l.PushBack(3)
				l.Reverse()
				So(checkSinglyLinkedListPointers(l, []*Element{e3, e2, e1}), ShouldEqual, true)
			})
		})
	})
}
