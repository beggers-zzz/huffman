// Ben Eggers
// GNU GPL'd

package huffman

// Tests the huffTree

import (
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
// makeTreeFromNodeSlice tests
////////////////////////////////////////////////////////////////////////////////

func TestMakeTreeFromNodesOneNode(t *testing.T) {
	node := &huffNode{char: 'x', count: 10}
	nodes := []*huffNode{node}
	tree := huffNode(*makeTreeFromNodeSlice(nodes))
	if tree.char != 'x' || tree.count != 10 {
		t.Error("Unexpected! Got node with count: ", tree.count, " and char: ",
			tree.char, "instead of count: ", node.count,
			" and char: ", node.char, ".")
	}
}

func TestMakeTreeFromNodesBasicTree(t *testing.T) {
	nodes := []*huffNode{{char: 'x', count: 2}, {char: 'x', count: 2}}
	tree := huffNode(*makeTreeFromNodeSlice(nodes))
	if tree.count != 4 {
		t.Error("Tree root count should have been 4, was: ", tree.count, ".")
	}
	if tree.left.count != 2 || tree.left.char != 'x' {
		t.Error("Tree's left node was wrong! Expected {'x', 2}, got {'",
			tree.left.char, "', ", tree.left.count, "}")
	}
	if tree.right.count != 2 || tree.right.char != 'x' {
		t.Error("Tree's right node was wrong! Expected {'x', 2}, got {'",
			tree.right.char, "', ", tree.right.count, "}")
	}
}
