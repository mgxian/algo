package strmatch

import (
	"github.com/mgxian/algo/go/queue"
)

// ACNode is AC node data structure
type ACNode struct {
	val          byte
	children     []*ACNode
	isEndingChar bool
	fail         *ACNode
	length       int
}

// NewACNode create a new AC node
func NewACNode(val byte) *ACNode {
	return &ACNode{
		val:      val,
		children: make([]*ACNode, charMAXSIZE),
	}
}

// ACAutomation is AC automation data structure
type ACAutomation struct {
	root  *ACNode
	count int
}

// NewACAutomation create a new AC automation
func NewACAutomation() *ACAutomation {
	root := NewACNode('/')
	return &ACAutomation{
		root: root,
	}
}

// Insert save a word to the AC automation
func (ac *ACAutomation) Insert(word string) {
	p := ac.root
	for _, v := range word {
		char := byte(v)
		index := char - 'a'
		if p.children[index] == nil {
			node := NewACNode(char)
			node.length = p.length + 1
			p.children[index] = node
		}
		p = p.children[index]
	}

	p.isEndingChar = true
	ac.count += len(word)
}

// BuildACAutomation build ac automation
func (ac *ACAutomation) BuildACAutomation() {
	ac.root.fail = nil
	aQueue := queue.NewSimpleQueue(ac.count)
	aQueue.EnQueue(ac.root)
	for !aQueue.IsEmpty() {
		data, _ := aQueue.DeQueue()
		p := data.(*ACNode)
		for _, pc := range p.children {
			if pc == nil {
				continue
			}

			if p == ac.root {
				pc.fail = ac.root
			} else {
				q := p.fail
				for q != nil {
					qc := q.children[pc.val-'a']
					if qc != nil {
						pc.fail = qc
						break
					}
					q = q.fail
				}
				if q == nil {
					pc.fail = ac.root
				}
			}
			aQueue.EnQueue(pc)
		}
	}
}

// Match search match string in the text
func (ac *ACAutomation) Match(text string) (matches [][]int) {
	p := ac.root
	for i, v := range text {
		index := v - 'a'
		for p.children[index] == nil && p != ac.root {
			p = p.fail
		}

		p = p.children[index]

		if p == nil {
			p = ac.root
		}

		tmp := p
		for tmp != ac.root {
			if tmp.isEndingChar == true {
				pos := i - tmp.length + 1
				matches = append(matches, []int{pos, tmp.length})
			}
			tmp = tmp.fail
		}
	}
	return
}
