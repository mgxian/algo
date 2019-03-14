package tree

// Node is a tree node structure
type Node struct {
	left, right *Node
	val         int
}

// NewNode create a new tree node
func NewNode(val int) *Node {
	return &Node{
		val: val,
	}
}

// GetVal return the node's val
func (n Node) GetVal() int {
	return n.val
}

// Tree is collection of tree operation
type Tree interface {
	Insert(val int)
	Delete(val int)
	Find(val int) (exist bool)
	FindMax() (node *Node)
	FindMin() (node *Node)
}
