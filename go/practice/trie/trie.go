package trie

type trieNode struct {
	char     byte
	isWord   bool
	children []*trieNode
}

func newTrieNode(c byte) *trieNode {
	return &trieNode{
		char:     c,
		children: make([]*trieNode, 26),
	}
}

func (tn *trieNode) addChild(c byte) *trieNode {
	index := int(c) - int('a')
	child := newTrieNode(c)
	tn.children[index] = child
	return child
}

func (tn *trieNode) getChild(c byte) *trieNode {
	index := int(c) - int('a')
	return tn.children[index]
}

type trieTree struct {
	root *trieNode
}

func newTrieTree() *trieTree {
	return &trieTree{
		root: newTrieNode('/'),
	}
}

func (tt *trieTree) addString(aString string) {
	n := len(aString)
	prev := tt.root
	for i, c := range aString {
		bc := byte(c)
		child := prev.getChild(bc)
		if child == nil {
			child = prev.addChild(bc)
		}

		if i == n-1 {
			child.isWord = true
		}

		prev = child
	}
}

func (tt *trieTree) find(aString string) bool {
	n := len(aString)
	prev := tt.root
	for i, c := range aString {
		bc := byte(c)
		child := prev.getChild(bc)
		if child == nil {
			return false
		}

		if i == n-1 && child.isWord == true {
			return true
		}

		prev = child
	}
	return false
}
