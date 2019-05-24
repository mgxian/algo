package binarytree

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchTree(t *testing.T) {
	outputWriterOri := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = outputWriterOri }()
	aTree := newTree()
	aTree.insert(10)
	aTree.insert(3)
	aTree.insert(5)
	aTree.insert(1)
	aTree.insert(15)
	aTree.insert(12)
	aTree.insert(19)
	aTree.insert(2)
	aTree.insert(0)
	aTree.insert(11)
	aTree.insert(20)

	aTree.inOrderTraverse()
	assert.Equal(t, "0\n1\n2\n3\n5\n10\n11\n12\n15\n19\n20\n", outputWriter.(*strings.Builder).String())

	assert.Equal(t, false, aTree.find(-1))
	assert.Equal(t, true, aTree.find(3))
	assert.Equal(t, true, aTree.find(10))
	assert.Equal(t, true, aTree.find(19))
	assert.Equal(t, true, aTree.find(12))
	assert.Equal(t, false, aTree.find(21))

	outputWriter.(*strings.Builder).Reset()
	aTree.remove(4)
	aTree.inOrderTraverse()
	assert.Equal(t, "0\n1\n2\n3\n5\n10\n11\n12\n15\n19\n20\n", outputWriter.(*strings.Builder).String())

	outputWriter.(*strings.Builder).Reset()
	aTree.remove(0)
	aTree.inOrderTraverse()
	assert.Equal(t, "1\n2\n3\n5\n10\n11\n12\n15\n19\n20\n", outputWriter.(*strings.Builder).String())

	outputWriter.(*strings.Builder).Reset()
	aTree.remove(20)
	aTree.inOrderTraverse()
	assert.Equal(t, "1\n2\n3\n5\n10\n11\n12\n15\n19\n", outputWriter.(*strings.Builder).String())

	outputWriter.(*strings.Builder).Reset()
	aTree.remove(10)
	aTree.inOrderTraverse()
	assert.Equal(t, "1\n2\n3\n5\n11\n12\n15\n19\n", outputWriter.(*strings.Builder).String())

	outputWriter.(*strings.Builder).Reset()
	aTree.remove(11)
	aTree.inOrderTraverse()
	assert.Equal(t, "1\n2\n3\n5\n12\n15\n19\n", outputWriter.(*strings.Builder).String())
}
