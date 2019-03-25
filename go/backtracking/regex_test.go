package backtracking

import "testing"

func TestRegexMatch(t *testing.T) {
	text := "abcdefghijklmn"
	matched := false

	pattern := "abc*lmn"
	expected := true
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}

	pattern = "abcdefg?hijklmn"
	expected = true
	matched = false
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}

	pattern = "abcdef?hijklmn"
	expected = true
	matched = false
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}

	pattern = "*abcdefghijklmn"
	expected = true
	matched = false
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}

	pattern = "abcdefghijklmn*"
	expected = true
	matched = false
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}

	pattern = "?abcdefghijklmn"
	expected = true
	matched = false
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}

	pattern = "abcdefghijklmn?"
	expected = true
	matched = false
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}

	pattern = "abcdefgh?lmn"
	expected = false
	matched = false
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}

	pattern = "*"
	expected = true
	matched = false
	regexMatch(text, pattern, 0, 0, &matched)
	if expected != matched {
		t.Errorf("expected %v, got %v", expected, matched)
	}
}
