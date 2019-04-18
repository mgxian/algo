package array

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMergSortedArray(t *testing.T) {
	Convey("Given two sorted array and less compare function", t, func() {
		arr1 := []Value{1, 2, 6, 7, 9}
		arr2 := []Value{0, 3, 4, 5, 8}
		lessFn := func(x, y Value) bool { return x.(int) < y.(int) }

		Convey("When merge the two sorted array", func() {
			actual := mergeSortedArray(arr1, arr2, len(arr1), len(arr2), lessFn)

			Convey("The result should be a merged sorted array", func() {
				expeted := []Value{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
				So(actual, ShouldResemble, expeted)
			})
		})
	})
}
