package trie

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrieTree(t *testing.T) {
	strings := []string{"hello", "hi", "hey", "world", "how", "why", "start", ""}
	findStrings := []string{"hi", "hey", "star", "abc", "", "notfound"}
	expected := []bool{true, true, false, false, false, false}

	tt := newTrieTree()
	for _, aString := range strings {
		tt.addString(aString)
	}

	for i, aString := range findStrings {
		assert.Equal(t, expected[i], tt.find(aString))
	}
}
