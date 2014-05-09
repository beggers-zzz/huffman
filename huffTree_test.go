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

func TestBasicTree(t *testing.T) {
	node := &huffNode{char: 'x', count: 10}
	nodes := []*huffNode{node}
	tree := huffNode(*makeTreeFromNodeSlice(nodes))
	if tree.char != 'x' || tree.count != 10 {
		t.Error("Unexpected! Got node with count: ", tree.count, " and char: ",
			tree.char, "instead of count: ", node.count,
			" and char: ", node.char, ".")
	}
}
