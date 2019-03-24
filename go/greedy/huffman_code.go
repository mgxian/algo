package greedy

import (
	"fmt"

	"github.com/mgxian/algo/go/sort"
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

func (p Chars) String() string {
	ret := "chars: "
	for _, c := range p {
		ret += fmt.Sprintf("(%s: %d)", string(c.val), c.count)
	}

	return ret
}

func NewChars(s string) Chars {
	chars := make([]int, CharacterSize)
	for _, v := range s {
		index := v - 'a'
		chars[index]++
	}

	var aChars Chars
	for c, n := range chars {
		if n != 0 {
			aChar := NewChar(byte(c+'a'), n)
			aChars = append(aChars, aChar)
		}
	}

	sort.QuickSort(aChars, 0, len(aChars)-1)

	return aChars
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
// }
