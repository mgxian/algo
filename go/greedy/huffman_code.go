package greedy

import (
	"container/heap"
)

const CharacterSize = 27

type Char struct {
	val   byte
	count int
}

func NewChar(val byte, n int) Char {
	return Char{
		val:   val,
		count: n,
	}
}

type Chars []Char

func (p Chars) Len() int {
	return len(p)
}

func (p Chars) Less(i, j int) bool {
	return p[i].count < p[j].count
}

func (p Chars) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Chars) Push(x interface{}) {
	*p = append(*p, x.(Char))
}

func (p *Chars) Pop() interface{} {
	old := *p
	n := len(old)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func (p *Chars) String() string {
	ret := "chars: "

	for {
		c := heap.Pop(p)
		if c == nil {
			break
		}
		// aChar := c.(Char)
		// ret += fmt.Sprintf("(%s: %d)", string(aChar.val), aChar.count)
	}

	return ret
}

func NewChars(s string) *Chars {
	chars := make([]int, CharacterSize)
	for _, v := range s {
		index := v - 'a'
		chars[index]++
	}

	aChars := make(Chars, 0)
	heap.Init(&aChars)
	for c, n := range chars {
		if n != 0 {
			aChar := NewChar(byte(c+'a'), n)
			heap.Push(&aChars, aChar)
		}
	}
	return &aChars
}

type treeNode struct {
	data  Char
	left  *treeNode
	right *treeNode
}

func newTreeNode(c Char) *treeNode {
	return &treeNode{
		data: c,
	}
}

// func HuffmanCode(aChars Chars) {
// 	aChar := NewChar('/', aChars[0].count+aChars[1].count)
// 	node := newTreeNode(aChar)
// 	for i := 2; i < aChars.Len(); i++ {

// 	}

// 	heap.pu
// }
