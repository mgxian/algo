package strmatch

import "testing"

func TestBFSearch(t *testing.T) {
	s := "a"
	pattern := "ab"
	expected := -1
	actual := BFSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "aba"
	pattern = "ab"
	expected = 0
	actual = BFSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "acbab"
	pattern = "ab"
	expected = 3
	actual = BFSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}

func TestRKSearch(t *testing.T) {
	s := "a"
	pattern := "ab"
	expected := -1
	actual := RKSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "aba"
	pattern = "ab"
	expected = 0
	actual = RKSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "acbab"
	pattern = "ab"
	expected = 3
	actual = RKSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}

func TestBMSearch(t *testing.T) {
	s := "a"
	pattern := "ab"
	expected := -1
	actual := BMSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "aba"
	pattern = "ab"
	expected = 0
	actual = BMSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "acbab"
	pattern = "ab"
	expected = 3
	actual = BMSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}
