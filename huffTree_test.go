// Ben Eggers
// GNU GPL'd

package huffman

// Tests the huffTree

import (
	"io/ioutil"
	"os"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
// MakeTreeFromText tests
////////////////////////////////////////////////////////////////////////////////

func TestMakeTreeFromTextEmpty(t *testing.T) {
	b := make([]byte, 0)
	err := ioutil.WriteFile(".test", b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(".test")

	tree, err := MakeTreeFromText(".test")
	if err == nil {
		t.Error("Got nil error from MakeTreeFromText! Should be 'Text file empty'")
	}
}

func TestMakeTreeFromTextSingleChar(t *testing.T) {
	b := []byte{0}
	err := ioutil.WriteFile(".test", b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(".test")

	tree, err := MakeTreeFromText(".test")
	if err != nil {
		t.Error("Got non-nil error from MakeTreeFromText: ", err)
	}
	
	root := tree.root
	if root.count != 1 || root.char != 0 {
		t.Error("Tree was built improperly! Expected: { char: 0, count: 1 },",
			"got { char:", root.char, ", count:", root.count, "}")
	}
}

func TestMakeTreeFromTextTwoOfSameChar(t *testing.T) {
	b := []byte{0, 0}
	err := ioutil.WriteFile(".test", b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(".test")
	
	tree, err := MakeTreeFromText(".test")
	if err != nil {
		t.Error("Got non-nil error from MakeTreeFromText: ", err)
	}
	
	root := tree.root
	if root.count != 2 || root.char != 0 {
		t.Error("Tree was built improperly! Expected: { char: 0, count: 2 },",
			"got { char:", root.char, ", count:", root.count, "}")
	}
}

func TestMakeTreeFromTextBasicTree(t *testing.T) {
	b := []byte{0, 0, 2}
	err := ioutil.WriteFile(".test", b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(".test")
	
	tree, err := MakeTreeFromText(".test")
	if err != nil {
		t.Error("Got non-nil error from MakeTreeFromText: ", err)
	}
	
	root := tree.root
	if root.count != 3 {
		t.Error("Tree's count was wrong! Should be 3, was", root.count)
	}
	if root.left.count != 1 || root.left.char != 2 {
		t.Error("Tree was built improperly! Expected: { char: 2, count: 1 },",
			"got { char:", root.left.char, ", count:", root.left.count, "}")
	}
	if root.right.count != 2 || root.right.char != 0 {
		t.Error("Tree was built improperly! Expected: { char: 0, count: 2 },",
			"got { char:", root.right.char, ", count:", root.right.count, "}")
	}
}

func TestMakeTreeFromTextMultiLevelTree(t *testing.T) {
	b := []byte{0, 0, 1, 1, 2, 2, 2}
	err := ioutil.WriteFile(".test", b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(".test")
	
	tree, err := MakeTreeFromText(".test")
	if err != nil {
		t.Error("Got non-nil error from MakeTreeFromText: ", err)
	}
	
	root := tree.root
	if root.count != 7 {
		t.Error("Tree's count was wrong! Should be 7, was", root.count)
	}
	if root.left.count != 3 || root.left.char != 2 {
		t.Error("Tree was built improperly! Expected: { char: 2, count: 3 },",
			"got { char:", root.left.char, ", count:", root.left.count, "}")
	}
	if root.right.count != 4 {
		t.Error("Tree.right's count was wrong! Should be 4, was", root.right.count)
	}
	if root.right.left.count != 2 || root.right.left.char != 0 {
		t.Error("Tree was built improperly! On tree.right.left: Expected:,",
			"{ char: 0, count: 2 },", "got { char:", root.right.char,
			", count:", root.right.count, "}")
	}
	if root.right.right.count != 2 || root.right.right.char != 1 {
		t.Error("Tree was built improperly! On tree.right.left: Expected:,",
			"{ char: 1, count: 2 },", "got { char:", root.right.char,
			", count:", root.right.count, "}")
	}
}

////////////////////////////////////////////////////////////////////////////////
// makeTreeFromNodeSlice tests
////////////////////////////////////////////////////////////////////////////////

func TestMakeTreeFromNodesEmpty(t *testing.T) {
	nodes := []*huffNode{}
	tree := makeTreeFromNodeSlice(nodes)
	if tree != nil {
		t.Error("Tree wasn't nil! tree: ", tree, ".")
	}
}

func TestMakeTreeFromNodesOneNode(t *testing.T) {
	node := &huffNode{char: 120, count: 10}
	nodes := []*huffNode{node}
	tree := makeTreeFromNodeSlice(nodes)
	if tree.char != 120 || tree.count != 10 {
		t.Error("Unexpected! Got node with count:", tree.count, "and char:",
			tree.char, "instead of count:", node.count, "and char:", node.char)
	}
}

func TestMakeTreeFromNodesBasicTree(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 2}, {char: 120, count: 2}}
	tree := makeTreeFromNodeSlice(nodes)
	if tree.count != 4 {
		t.Error("Tree root count should have been 4, was: ", tree.count, ".")
	}
	if tree.left.count != 2 || tree.left.char != 120 {
		t.Error("Tree's left node was wrong! Expected { 120, 2 }, got {",
			tree.left.char, ",", tree.left.count, "}")
	}
	if tree.right.count != 2 || tree.right.char != 120 {
		t.Error("Tree's right node was wrong! Expected { 120, 2 }, got {",
			tree.right.char, ",", tree.right.count, "}")
	}
}

// This test is fairly tied to the implementation, but tests of something
// internal (like this) often have to be :(
func TestMakeTreeFromNodesMultiLevelTree(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 2},
		{char: 120, count: 2},
		{char: 121, count: 3}}
	tree := makeTreeFromNodeSlice(nodes)
	if tree.count != 7 {
		t.Error("Tree root count should have been 7, was:", tree.count)
	}
	if tree.left.count != 3 || tree.left.char != 121 {
		t.Error("Tree's left node was wrong! Expected { 121, 3 }, got {",
			tree.left.char, ", ", tree.left.count, "}")
	}
	if tree.right.count != 4 {
		t.Error("Right subtree count should have been 4, was:", tree.right.count)
	}
	if tree.right.right.count != 2 || tree.right.right.char != 120 {
		t.Error("Tree's right node's right node was wrong! Expected { 120, 2 }, got {",
			tree.right.right.char, ",", tree.right.right.count)
	}
	if tree.right.left.count != 2 || tree.right.left.char != 120 {
		t.Error("Tree's right node's left node left node was wrong!",
			"Expected { 120, 2 }, got { ", tree.right.left.char,
			",", tree.right.left.count, "}")
	}
}
