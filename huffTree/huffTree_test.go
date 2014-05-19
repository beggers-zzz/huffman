// Ben Eggers
// GNU GPL'd

package huffTree

// Tests the huffTree

import (
	"io/ioutil"
	"os"
	"testing"
)

var filename = ".test"

////////////////////////////////////////////////////////////////////////////////
// MakeTreeFromText tests
////////////////////////////////////////////////////////////////////////////////

func TestMakeTreeFromTextEmpty(t *testing.T) {
	b := make([]byte, 0)
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	tree, err := MakeTreeFromText(filename)
	if err == nil {
		t.Error("Got nil error from MakeTreeFromText! Should be 'Text file empty'")
	}

	root := tree.root
	if root != nil {
		t.Error("Got non-nil root.")
	}
}

func TestMakeTreeFromTextSingleChar(t *testing.T) {
	b := []byte{0}
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	tree, err := MakeTreeFromText(filename)
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
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	tree, err := MakeTreeFromText(filename)
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
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	tree, err := MakeTreeFromText(filename)
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
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	tree, err := MakeTreeFromText(filename)
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
	tree, err := makeTreeFromNodeSlice(nodes)

	if err == nil {
		t.Error("Got nil error when we shouldn't have.")
	}

	root := tree.root
	if root != nil {
		t.Error("Got non-nil root.")
	}
}

func TestMakeTreeFromNodesOneNode(t *testing.T) {
	node := &huffNode{char: 120, count: 10}
	nodes := []*huffNode{node}
	tree, err := makeTreeFromNodeSlice(nodes)
	if err != nil {
		t.Error(err)
	}

	root := tree.root
	if root.char != 120 || root.count != 10 {
		t.Error("Unexpected! Got node with count:", root.count, "and char:",
			root.char, "instead of count:", node.count, "and char:", node.char)
	}
}

func TestMakeTreeFromNodesBasicTree(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 2}, {char: 120, count: 2}}
	tree, err := makeTreeFromNodeSlice(nodes)
	if err != nil {
		t.Error(err)
	}

	root := tree.root
	if root.count != 4 {
		t.Error("Tree's root count should have been 4, was: ", root.count, ".")
	}
	if root.left.count != 2 || root.left.char != 120 {
		t.Error("Tree's left node was wrong! Expected { 120, 2 }, got {",
			root.left.char, ",", root.left.count, "}")
	}
	if root.right.count != 2 || root.right.char != 120 {
		t.Error("Tree's right node was wrong! Expected { 120, 2 }, got {",
			root.right.char, ",", root.right.count, "}")
	}
}

// This test is fairly tied to the implementation, but tests of something
// internal (like this) often have to be :(
func TestMakeTreeFromNodesMultiLevelTree(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 2},
		{char: 120, count: 2},
		{char: 121, count: 3}}
	tree, err := makeTreeFromNodeSlice(nodes)
	if err != nil {
		t.Error(err)
	}

	root := tree.root
	if root.count != 7 {
		t.Error("Tree's root count should have been 7, was:", root.count)
	}
	if root.left.count != 3 || root.left.char != 121 {
		t.Error("Tree's left node was wrong! Expected { 121, 3 }, got {",
			root.left.char, ", ", root.left.count, "}")
	}
	if root.right.count != 4 {
		t.Error("Right subtree count should have been 4, was:", root.right.count)
	}
	if root.right.right.count != 2 || root.right.right.char != 120 {
		t.Error("Tree's right node's right node was wrong! Expected { 120, 2 }, got {",
			root.right.right.char, ",", root.right.right.count)
	}
	if root.right.left.count != 2 || root.right.left.char != 120 {
		t.Error("Tree's right node's left node was wrong!",
			"Expected { 120, 2 }, got { ", root.right.left.char,
			",", root.right.left.count, "}")
	}
}
