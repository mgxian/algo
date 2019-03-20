package strmatch

const charMAXSIZE = 26 // suppose only have 26 characters

// TrieNode is trie node data structure
type TrieNode struct {
	val          byte
	children     []*TrieNode
	isEndingChar bool
}

// NewTrieNode create a new trie node
func NewTrieNode(val byte) *TrieNode {
	return &TrieNode{
		val:      val,
		children: make([]*TrieNode, charMAXSIZE),
	}
}

// TrieTree is trie tree data structure
type TrieTree struct {
	root *TrieNode
}

// NewTrieTree create a new trie tree
func NewTrieTree() *TrieTree {
	root := NewTrieNode('/')
	return &TrieTree{
		root: root,
	}
}

// Insert save a word to the trie tree
func (t *TrieTree) Insert(word string) {
	p := t.root
	for _, v := range word {
		char := byte(v)
		index := char - 'a'
		if p.children[index] == nil {
			node := NewTrieNode(char)
			p.children[index] = node
		}
		p = p.children[index]
	}

	p.isEndingChar = true
}

// Search word in the trie tree
func (t *TrieTree) Search(word string) bool {
	p := t.root
	for _, v := range word {
		char := byte(v)
		index := char - 'a'
		p = p.children[index]
		if p == nil {
			return false
		}
	}

	if p.isEndingChar {
		return true
	}

	return false
}
