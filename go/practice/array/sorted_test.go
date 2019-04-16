package array

import "testing"

func TestSortedArray(t *testing.T) {
	var aSortedArray SortedArrayInterface

	lessFn := func(x, y Value) bool { return x.(int) < y.(int) }
	equalFn := func(x, y Value) bool { return x.(int) == y.(int) }

	// single element SortedArray
	aSortedArray = NewSortedArray(1, lessFn, equalFn)
	aSortedArray.Add(1)
	t.Log(aSortedArray)

	// bigger SortedArray
	aSortedArray = NewSortedArray(5, lessFn, equalFn)
	aSortedArray.Add(4)
	aSortedArray.Add(5)
	aSortedArray.Add(1)
	aSortedArray.Add(3)
	aSortedArray.Add(2)

	t.Log(aSortedArray)
}

func TestSortedArrayInsert(t *testing.T) {
	var aSortedArray SortedArrayInterface

	lessFn := func(x, y Value) bool { return x.(int) < y.(int) }
	equalFn := func(x, y Value) bool { return x.(int) == y.(int) }

	aSortedArray = NewSortedArray(3, lessFn, equalFn)
	aSortedArray.Add(1)
	aSortedArray.Add(3)
	aSortedArray.Add(2)
	t.Log(aSortedArray)
}

func TestSortedArrayRemove(t *testing.T) {
	var aSortedArray SortedArrayInterface

	lessFn := func(x, y Value) bool { return x.(int) < y.(int) }
	equalFn := func(x, y Value) bool { return x.(int) == y.(int) }

	aSortedArray = NewSortedArray(6, lessFn, equalFn)

	aSortedArray.Add(1)
	aSortedArray.Add(2)
	aSortedArray.Add(3)
	aSortedArray.Add(4)
	aSortedArray.Add(5)
	aSortedArray.Add(3)
	t.Log(aSortedArray)
	aSortedArray.RemoveAll(2)
	t.Log(aSortedArray)
	aSortedArray.RemoveAll(3)
	t.Log(aSortedArray)
}

func TestSortedArrayFind(t *testing.T) {
	var aSortedArray SortedArrayInterface

	lessFn := func(x, y Value) bool { return x.(int) < y.(int) }
	equalFn := func(x, y Value) bool { return x.(int) == y.(int) }

	aSortedArray = NewSortedArray(7, lessFn, equalFn)

	aSortedArray.Add(1)
	aSortedArray.Add(2)
	aSortedArray.Add(3)
	aSortedArray.Add(4)
	aSortedArray.Add(5)
	aSortedArray.Add(6)
	aSortedArray.Add(3)

	t.Log(aSortedArray)
	t.Log(aSortedArray.Find(3))
	t.Log(aSortedArray.Find(4))
	t.Log(aSortedArray.FindAll(3))
}

func TestSortedArrayReplace(t *testing.T) {
	var aSortedArray SortedArrayInterface

	lessFn := func(x, y Value) bool { return x.(int) < y.(int) }
	equalFn := func(x, y Value) bool { return x.(int) == y.(int) }

	aSortedArray = NewSortedArray(7, lessFn, equalFn)
	aSortedArray.Add(1)
	aSortedArray.Add(2)
	aSortedArray.Add(3)
	aSortedArray.Add(4)
	aSortedArray.Add(5)
	aSortedArray.Add(6)
	aSortedArray.Add(3)
	t.Log(aSortedArray)

	aSortedArray.Replace(4, 5)
	t.Log(aSortedArray)

	aSortedArray.Replace(3, 4)
	t.Log(aSortedArray)
}
