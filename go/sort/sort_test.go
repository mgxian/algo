package sort

import "testing"

type MyInts []int

func (p MyInts) Len() int           { return len(p) }
func (p MyInts) Less(i, j int) bool { return p[i] < p[j] }
func (p MyInts) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

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

func TestInsertSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	InsertSort(MyInts(numbers))
	t.Log(numbers)

	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}
}

func TestSelectSort(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	SelectSort(MyInts(numbers))
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
