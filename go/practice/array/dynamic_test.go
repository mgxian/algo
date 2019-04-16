package array

import "testing"

func TestDynamicArray(t *testing.T) {
	var aDynamicArray DynamicInterface

	// single element DynamicArray
	aDynamicArray = NewDynamicArray(1)
	aDynamicArray.InsertBack(1)
	aDynamicArray.InsertBack(2)
	aDynamicArray.InsertBack(3)
	t.Log(aDynamicArray)

	// bigger DynamicArray
	aDynamicArray = NewDynamicArray(3)
	aDynamicArray.InsertBack(1)
	aDynamicArray.InsertBack(2)
	aDynamicArray.InsertBack(3)
	aDynamicArray.InsertBack(4)
	aDynamicArray.InsertBack(5)
	t.Log(aDynamicArray)
}

func TestDynamicArrayInsert(t *testing.T) {
	var aDynamicArray DynamicInterface
	aDynamicArray = NewDynamicArray(2)
	aDynamicArray.InsertBack(1)
	aDynamicArray.InsertBack(3)
	aDynamicArray.InsertFront(0)
	aDynamicArray.Insert(2, 2)
	t.Log(aDynamicArray)
}

func TestDynamicArrayRemove(t *testing.T) {
	var aDynamicArray DynamicInterface
	aDynamicArray = NewDynamicArray(5)
	aDynamicArray.InsertBack(1)
	aDynamicArray.InsertBack(2)
	aDynamicArray.InsertBack(3)
	aDynamicArray.InsertBack(4)
	aDynamicArray.InsertBack(5)
	aDynamicArray.InsertBack(6)
	aDynamicArray.RemoveFront()
	aDynamicArray.RemoveBack()
	aDynamicArray.Remove(2)

	t.Log(aDynamicArray)
}

func TestDynamicArrayAccess(t *testing.T) {
	var aDynamicArray DynamicInterface
	aDynamicArray = NewDynamicArray(5)
	aDynamicArray.InsertBack(1)
	aDynamicArray.InsertBack(2)
	aDynamicArray.InsertBack(3)
	aDynamicArray.InsertBack(4)
	aDynamicArray.InsertBack(5)
	aDynamicArray.InsertBack(6)
	aDynamicArray.InsertBack(3)

	v, _ := aDynamicArray.Get(3)
	t.Log(v)

	fn := func(x, y Value) bool { return x.(int) == y.(int) }
	t.Log(aDynamicArray.Find(2, fn))
	t.Log(aDynamicArray.FindAll(3, fn))
	t.Log(aDynamicArray.Contains(5, fn))

	t.Log(aDynamicArray.Find(10, fn))
	t.Log(aDynamicArray.FindAll(10, fn))
	t.Log(aDynamicArray.Contains(10, fn))
}

func TestDynamicArraySet(t *testing.T) {
	var aDynamicArray DynamicInterface
	aDynamicArray = NewDynamicArray(2)
	aDynamicArray.InsertBack(1)
	aDynamicArray.InsertBack(3)
	aDynamicArray.InsertFront(0)
	aDynamicArray.Set(2, 2)
	t.Log(aDynamicArray)
}
