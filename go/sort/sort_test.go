package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	BubbleSort(MyInts(numbers))
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	InsertionSort(MyInts(numbers))
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	SelectionSort(MyInts(numbers))
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestShellSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4, 8, 6, 9, 0, 7}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	ShellSort(MyInts(numbers))
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestMergeSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4, 8, 6, 9, 0, 7}
	// numbers := []int{5, 3, 1}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	MergeSort(MyInts(numbers), 0, len(numbers)-1)
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestQuickSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4, 8, 6, 9, 0, 7}
	// numbers := []int{5, 3, 1}
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	QuickSort(MyInts(numbers), 0, len(numbers)-1)
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestBucketSort(t *testing.T) {
	numbers := []int{11, 22, 100, 34, 5, 47, 59, 60, 72, 88, 95, 15, 14, 16}
	expected := []int{5, 11, 14, 15, 16, 22, 34, 47, 59, 60, 72, 88, 95, 100}
	BucketSort(numbers)
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestCountingSort(t *testing.T) {
	numbers := []int{11, 22, 100, 34, 5, 47, 59, 60, 72, 88, 95, 15, 14, 16}
	expected := []int{5, 11, 14, 15, 16, 22, 34, 47, 59, 60, 72, 88, 95, 100}
	CountingSort(numbers)
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestRadixSort(t *testing.T) {
	numbers := []int{11, 22, 100, 34, 5, 47, 59, 60, 72, 88, 95, 15, 14, 16}
	expected := []int{5, 11, 14, 15, 16, 22, 34, 47, 59, 60, 72, 88, 95, 100}
	RadixSort(numbers)
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}
