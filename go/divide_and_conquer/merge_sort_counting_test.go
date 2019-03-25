package divide

import "testing"

func TestMergeSortCounting(t *testing.T) {
	numbers := []int{5, 3, 1, 2, 4}
	expected := []int{1, 2, 3, 4, 5}
	count := 6
	n := countReversePairs(numbers)
	for i, num := range numbers {
		if num != expected[i] {
			t.Errorf("expected %d but got %d", expected[i], num)
		}
	}

	if count != n {
		t.Errorf("expected %d but got %d", count, n)
	}
}
