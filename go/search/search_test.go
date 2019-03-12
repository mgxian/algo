package search

import "testing"

func TestBinarySearch(t *testing.T) {
	numbers := []int{5, 11, 14, 15, 16, 22, 34, 47, 59, 60, 72, 88, 95, 100}
	find := []int{5, 15, 34, 72, 100, 2}
	expected := []int{0, 3, 6, 10, 13, -1}

	for i, v := range find {
		index := BinarySearch(numbers, v)
		if index != expected[i] {
			t.Errorf("expected %d, but got %d", expected[i], index)
		}
	}
}

func TestBinarySearchFirstEqual(t *testing.T) {
	numbers := []int{5, 11, 14, 15, 16, 22, 34, 34, 47, 59, 60, 72, 88, 95, 100, 100}
	find := []int{5, 15, 34, 72, 100, 2}
	expected := []int{0, 3, 6, 11, 14, -1}

	for i, v := range find {
		index := BinarySearchFirstEqual(numbers, v)
		if index != expected[i] {
			t.Errorf("expected %d, but got %d", expected[i], index)
		}
	}
}

func TestBinarySearchLastEqual(t *testing.T) {
	numbers := []int{5, 11, 14, 15, 16, 22, 34, 34, 47, 59, 60, 72, 88, 95, 100, 100}
	find := []int{5, 15, 34, 72, 100, 2}
	expected := []int{0, 3, 7, 11, 15, -1}

	for i, v := range find {
		index := BinarySearchLastEqual(numbers, v)
		if index != expected[i] {
			t.Errorf("expected %d, but got %d", expected[i], index)
		}
	}
}

func TestBinarySearchFirstGreaterThanOrEqualTo(t *testing.T) {
	numbers := []int{5, 11, 14, 15, 16, 22, 34, 34, 47, 59, 60, 72, 88, 95, 100, 100}
	find := []int{5, 15, 34, 72, 100, 2}
	expected := []int{0, 3, 7, 11, 15, -1}

	for i, v := range find {
		index := BinarySearchFirstGreaterThanOrEqualTo(numbers, v)
		if index != expected[i] {
			t.Errorf("expected %d, but got %d", expected[i], index)
		}
	}
}

func TestBinarySearchLastLessThanOrEqualTo(t *testing.T) {
	numbers := []int{5, 11, 14, 15, 16, 22, 34, 34, 47, 59, 60, 72, 88, 95, 100, 100}
	find := []int{5, 15, 34, 72, 100, 2}
	expected := []int{0, 3, 7, 11, 15, -1}

	for i, v := range find {
		index := BinarySearchLastLessThanOrEqualTo(numbers, v)
		if index != expected[i] {
			t.Errorf("expected %d, but got %d", expected[i], index)
		}
	}
}
