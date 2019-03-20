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

	s = "acdabcde"
	pattern = "bc"
	expected = 4
	actual = BFSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "abcacababcababcdc"
	pattern = "abcababc"
	expected = 7
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

	s = "acdabcde"
	pattern = "bc"
	expected = 4
	actual = RKSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "abcacababcababcdc"
	pattern = "abcababc"
	expected = 7
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

	s = "acdabcde"
	pattern = "bc"
	expected = 4
	actual = BMSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "abcacababcababcdc"
	pattern = "abcababc"
	expected = 7
	actual = BMSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}

func TestKMPSearch(t *testing.T) {
	s := "a"
	pattern := "ab"
	expected := -1
	actual := KMPSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "aba"
	pattern = "ab"
	expected = 0
	actual = KMPSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "acbab"
	pattern = "ab"
	expected = 3
	actual = KMPSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "abcacabdc"
	pattern = "abd"
	expected = 5
	actual = KMPSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}

	s = "abcacababcababcdc"
	pattern = "abcababc"
	expected = 7
	actual = KMPSearch(s, pattern)
	if expected != actual {
		t.Errorf("expected %d, but got %d", expected, actual)
	}
}

func TestTrieTreeSearch(t *testing.T) {
	strings := []string{"hi", "how", "her", "hello", "so", "see"}
	aTrieTree := NewTrieTree()
	for _, v := range strings {
		aTrieTree.Insert(v)
	}

	s := "so"
	expected := true
	actual := aTrieTree.Search(s)
	if expected != actual {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	s = "hi"
	expected = true
	actual = aTrieTree.Search(s)
	if expected != actual {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	s = "his"
	expected = false
	actual = aTrieTree.Search(s)
	if expected != actual {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	s = "hell"
	expected = false
	actual = aTrieTree.Search(s)
	if expected != actual {
		t.Errorf("expected %v, but got %v", expected, actual)
	}

	s = "hello"
	expected = true
	actual = aTrieTree.Search(s)
	if expected != actual {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}
