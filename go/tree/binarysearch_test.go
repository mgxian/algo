package tree

import "testing"

func TestBinarySearchTree(t *testing.T) {
	numbers := []int{9, 8, 7, 6, 9, 5, 4, 6, 3, 2, 1}
	tree := NewBinarySearchTree()
	for _, v := range numbers {
		tree.Insert(v)
	}
	tree.InOrderTraverse()
	t.Log(tree.FindMax().GetVal())
	t.Log(tree.FindMin().GetVal())
	tree.Delete(9)
	tree.InOrderTraverse()
	tree.Delete(6)
	tree.InOrderTraverse()
}
