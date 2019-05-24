package binarytree

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type treeNode struct {
	value       int
	left, right *treeNode
}

func newTreeNode(v int) *treeNode {
	return &treeNode{
		value: v,
	}
}

type tree struct {
	root *treeNode
}

func newTree() *tree {
	return &tree{}
}

func (t *tree) insert(v int) {
	if t.root == nil {
		t.root = newTreeNode(v)
		return
	}

	prev := t.root
	curr := t.root
	isLeft := false
	for curr != nil {
		prev = curr
		if v < curr.value {
			curr = curr.left
			isLeft = true
			continue
		}
		curr = curr.right
		isLeft = false
	}

	if isLeft {
		prev.left = newTreeNode(v)
		return
	}
	prev.right = newTreeNode(v)
}

func (t *tree) getTreeNode(v int) (prev, curr *treeNode) {
	prev = t.root
	curr = t.root
	for curr != nil {
		if curr.value == v {
			break
		}
		prev = curr
		if v < curr.value {
			curr = curr.left
			continue
		}
		curr = curr.right
	}
	return
}

func (t *tree) remove(v int) {
	prev, curr := t.getTreeNode(v)
	if curr == nil {
		return
	}

	if curr == t.root {
		t.removeRoot()
		return
	}

	if curr == prev.left {
		t.removeLeftChild(prev)
		return
	}

	t.removeRightChild(prev)
}

func (t *tree) removeRoot() {
	if t.root.left == nil && t.root.right == nil {
		t.root = nil
		return
	}

	if t.root.left == nil && t.root.right != nil {
		t.root = t.root.right
		return
	}

	if t.root.left != nil && t.root.right == nil {
		t.root = t.root.left
		return
	}

	prev, curr := getMaxTreeNode(t.root.left)
	t.root.value = curr.value
	t.removeRightChild(prev)
}

func (t *tree) removeLeftChild(tn *treeNode) {
	deleteTreeNode := tn.left
	if deleteTreeNode.left == nil && deleteTreeNode.right == nil {
		tn.left = nil
		return
	}

	if deleteTreeNode.left == nil && deleteTreeNode.right != nil {
		tn.left = deleteTreeNode.right
		return
	}

	if deleteTreeNode.left != nil && deleteTreeNode.right == nil {
		tn.left = deleteTreeNode.left
		return
	}

	prev, curr := getMaxTreeNode(deleteTreeNode)
	deleteTreeNode.value = curr.value
	t.removeRightChild(prev)
}

func (t *tree) removeRightChild(tn *treeNode) {
	deleteTreeNode := tn.right
	if deleteTreeNode.left == nil && deleteTreeNode.right == nil {
		tn.right = nil
		return
	}

	if deleteTreeNode.left == nil && deleteTreeNode.right != nil {
		tn.right = deleteTreeNode.right
		return
	}

	if deleteTreeNode.left != nil && deleteTreeNode.right == nil {
		tn.right = deleteTreeNode.left
		return
	}

	prev, curr := getMaxTreeNode(deleteTreeNode)
	deleteTreeNode.value = curr.value
	t.removeRightChild(prev)
}

func (t *tree) find(v int) bool {
	curr := t.root
	for curr != nil {
		if curr.value == v {
			return true
		}
		if v < curr.value {
			curr = curr.left
			continue
		}
		curr = curr.right
	}

	return false
}

func (t *tree) preOrderTraverse() {
	preOrderTraverse(t.root)
}

func (t *tree) inOrderTraverse() {
	inOrderTraverse(t.root)
}

func (t *tree) postOrderTraverse() {
	postOrderTraverse(t.root)
}

func preOrderTraverse(tn *treeNode) {
	if tn == nil {
		return
	}
	fmt.Fprintln(outputWriter, tn.value)
	preOrderTraverse(tn.left)
	preOrderTraverse(tn.right)
}

func inOrderTraverse(tn *treeNode) {
	if tn == nil {
		return
	}
	inOrderTraverse(tn.left)
	fmt.Fprintln(outputWriter, tn.value)
	inOrderTraverse(tn.right)
}

func postOrderTraverse(tn *treeNode) {
	if tn == nil {
		return
	}

	postOrderTraverse(tn.left)
	postOrderTraverse(tn.right)
	fmt.Fprintln(outputWriter, tn.value)
}

func getMaxTreeNode(tn *treeNode) (prev, curr *treeNode) {
	if tn == nil {
		return
	}
	curr = tn
	for curr.right != nil {
		prev = curr
		curr = curr.right
	}
	return
}
