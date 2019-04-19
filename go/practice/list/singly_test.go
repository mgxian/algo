package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func checkSinglyLinkedListLen(l *SinglyLinkedList, es []*Element) bool {
	if len(es) != l.Len() {
		return false
	}
	return true
}

func checkSinglyLinkedListPointers(l *SinglyLinkedList, es []*Element) bool {
	if !checkSinglyLinkedListLen(l, es) {
		return false
	}

	if len(es) == 0 {
		if (l.head.next != nil && l.head.next != &l.head) || (l.head.prev != nil && l.head.prev != &l.head) {
			return false
		}
		return true
	}

	result := true
	prev := &l.head
	for _, e := range es {
		if prev.next != e {
			result = false
			break
		}
		prev = e
	}

	back := es[len(es)-1]
	if back.next != nil || l.head.prev != back {
		result = false
	}
	return result
}

func TestSinglyLinkedList(t *testing.T) {
	Convey("Test Single element SinglyLinkedList", t, func() {
		Convey("Test PushFront", func() {
			Convey("Given an empty SinglyLinkedList", func() {
				l := NewSinglyLinkedList()

				So(l.Front(), ShouldBeNil)
				So(l.Back(), ShouldBeNil)

				Convey("When PushFront a element to the SinglyLinkedList", func() {
					e := l.PushFront("a")

					Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
						So(checkSinglyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
					})

					So(l.Front(), ShouldEqual, e)
					So(l.Back(), ShouldEqual, e)

					Convey("When Remove element from the SinglyLinkedList", func() {
						l.Remove(e)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{}), ShouldEqual, true)
						})

						So(l.Front(), ShouldBeNil)
						So(l.Back(), ShouldBeNil)
					})
				})
			})
		})

		Convey("Test PushBack", func() {
			Convey("Given an empty SinglyLinkedList", func() {
				l := NewSinglyLinkedList()

				Convey("When PushBack a element to the SinglyLinkedList", func() {
					e := l.PushBack("a")

					Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
						So(checkSinglyLinkedListPointers(l, []*Element{e}), ShouldEqual, true)
					})

					So(l.Front(), ShouldEqual, e)
					So(l.Back(), ShouldEqual, e)

					Convey("When Remove element from the SinglyLinkedList", func() {
						l.Remove(e)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{}), ShouldEqual, true)
						})

						So(l.Front(), ShouldBeNil)
						So(l.Back(), ShouldBeNil)
					})
				})
			})
		})
	})

	Convey("Test multiple elements SinglyLinkedList", t, func() {
		Convey("Test PushFront", func() {
			Convey("Given an empty SinglyLinkedList", func() {
				l := NewSinglyLinkedList()

				Convey("When PushFront multiple elements to the SinglyLinkedList", func() {
					e1 := l.PushFront(1)
					e2 := l.PushFront(2)
					e3 := l.PushFront(3)
					e4 := l.PushFront(4)
					e5 := l.PushFront("a")

					Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
						So(checkSinglyLinkedListPointers(l, []*Element{e5, e4, e3, e2, e1}), ShouldEqual, true)
					})

					So(l.Front(), ShouldEqual, e5)
					So(l.Back(), ShouldEqual, e1)

					Convey("When Remove first element from the SinglyLinkedList", func() {
						l.Remove(e5)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{e4, e3, e2, e1}), ShouldEqual, true)
						})

						So(l.Front(), ShouldEqual, e4)
						So(l.Back(), ShouldEqual, e1)
					})

					Convey("When Remove last element from the SinglyLinkedList", func() {
						l.Remove(e1)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{e5, e4, e3, e2}), ShouldEqual, true)
						})

						So(l.Front(), ShouldEqual, e5)
						So(l.Back(), ShouldEqual, e2)
					})
				})
			})
		})

		Convey("Test PushBack", func() {
			Convey("Given an empty SinglyLinkedList", func() {
				l := NewSinglyLinkedList()

				Convey("When PushBack multiple elements to the SinglyLinkedList", func() {
					e1 := l.PushBack(1)
					e2 := l.PushBack(2)
					e3 := l.PushBack(3)
					e4 := l.PushBack(4)
					e5 := l.PushBack("a")

					Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
						So(checkSinglyLinkedListPointers(l, []*Element{e1, e2, e3, e4, e5}), ShouldEqual, true)
					})

					So(l.Front(), ShouldEqual, e1)
					So(l.Back(), ShouldEqual, e5)

					Convey("When insert a element after an element to the SinglyLinkedList", func() {
						e6 := l.InsertAfter(5, e4)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{e1, e2, e3, e4, e6, e5}), ShouldEqual, true)
						})

						So(l.Front(), ShouldEqual, e1)
						So(l.Back(), ShouldEqual, e5)
					})

					Convey("When insert a element before an element to the SinglyLinkedList", func() {
						e6 := l.InsertBefore(5, e5)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{e1, e2, e3, e4, e6, e5}), ShouldEqual, true)
						})

						So(l.Front(), ShouldEqual, e1)
						So(l.Back(), ShouldEqual, e5)
					})

					Convey("When insert a element to the front of the SinglyLinkedList", func() {
						e0 := l.InsertBefore(0, e1)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{e0, e1, e2, e3, e4, e5}), ShouldEqual, true)
						})

						So(l.Front(), ShouldEqual, e0)
						So(l.Back(), ShouldEqual, e5)
					})

					Convey("When insert a element to the back of the SinglyLinkedList", func() {
						e6 := l.InsertAfter(6, e5)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{e1, e2, e3, e4, e5, e6}), ShouldEqual, true)
						})

						So(l.Front(), ShouldEqual, e1)
						So(l.Back(), ShouldEqual, e6)
					})
				})
			})
		})

		Convey("Test PushFront and PushBack", func() {
			Convey("Given an empty SinglyLinkedList", func() {
				l := NewSinglyLinkedList()

				Convey("When PushFront multiple elements to the SinglyLinkedList", func() {
					e1 := l.PushFront(1)
					e2 := l.PushBack(2)
					e3 := l.PushBack(3)
					e4 := l.PushFront(4)
					e5 := l.PushBack("a")

					Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
						So(checkSinglyLinkedListPointers(l, []*Element{e4, e1, e2, e3, e5}), ShouldEqual, true)
					})

					So(l.Front(), ShouldEqual, e4)
					So(l.Back(), ShouldEqual, e5)

					Convey("When Move element to the front the SinglyLinkedList", func() {
						l.MoveToFront(e1)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{e1, e4, e2, e3, e5}), ShouldEqual, true)
						})

						So(l.Front(), ShouldEqual, e1)
						So(l.Back(), ShouldEqual, e5)
					})

					Convey("When Move element to the back the SinglyLinkedList", func() {
						l.MoveToBack(e4)

						Convey("Then SinglyLinkedList should pass checkSinglyLinkedListPointers", func() {
							So(checkSinglyLinkedListPointers(l, []*Element{e1, e2, e3, e5, e4}), ShouldEqual, true)
						})

						So(l.Front(), ShouldEqual, e1)
						So(l.Back(), ShouldEqual, e4)
					})
				})
			})
		})
	})
}
