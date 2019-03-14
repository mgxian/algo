package tree

import "fmt"

// BinarySearchTree is a binary search tree
type BinarySearchTree struct {
	root *Node
}

// NewBinarySearchTree create a new BinarySearchTree
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

// Insert save a item to the tree
func (tree *BinarySearchTree) Insert(val int) {
	p := tree.root
	if p == nil {
		tree.root = NewNode(val)
	}

	for p != nil {
		if val < p.val {
			if p.left == nil {
				p.left = NewNode(val)
				return
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = NewNode(val)
				return
			}
			p = p.right
		}
	}
}

// Delete remove a item to the tree
func (tree *BinarySearchTree) Delete(val int) {
	p := tree.root
	if p == nil {
		return
	}

	var pp *Node
	for p != nil && p.val != val {
		pp = p
		if val < p.val {
			p = p.left
		} else {
			p = p.right
		}
	}

	if p == nil {
		return
	}

	if p.left != nil && p.right != nil {
		minP, minPP := p.right, p
		for minP.left != nil {
			minPP = minP
			minP = p.left
		}
		p.val = minP.val
		p = minP
		pp = minPP
	}

	var child *Node
	if p.left != nil {
		child = p.left
	} else if p.right != nil {
		child = p.right
	} else {
		child = nil
	}

	if pp == nil {
		tree.root = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}
}

// InOrderTraverse in order traverse the tree
func (tree BinarySearchTree) InOrderTraverse() {
	p := tree.root
	if p == nil {
		return
	}

	if p.left != nil {
		tmpTree := NewBinarySearchTree()
		tmpTree.root = p.left
		tmpTree.InOrderTraverse()
	}
	fmt.Printf("-->%d\n", p.val)
	if p.right != nil {
		tmpTree := NewBinarySearchTree()
		tmpTree.root = p.right
		tmpTree.InOrderTraverse()
	}
}

// FindMax find max item from the tree
func (tree BinarySearchTree) FindMax() (node *Node) {
	p := tree.root
	for p.right != nil {
		p = p.right
	}
	node = p
	return
}

// FindMin find min item from the tree
func (tree BinarySearchTree) FindMin() (node *Node) {
	p := tree.root
	for p.left != nil {
		p = p.left
	}
	node = p
	return
}
