// Ben Eggers
// GNU GPL'd

package huffTree

// Tests the huffTree. Tests are (ostensibly) ordered in the order the function
// tested depend on each other, so failures are apparent at the lowest level.

import (
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
// makeTreeFromNodeSlice tests
////////////////////////////////////////////////////////////////////////////////

func TestMakeTreeFromNodesEmpty(t *testing.T) {
	nodes := []*huffNode{}
	root, err := makeTreeFromNodeSlice(nodes)

	if err == nil {
		t.Error("Got nil error when we shouldn't have.")
	}

	if root != nil {
		t.Error("Got non-nil root.")
	}
}

func TestMakeTreeFromNodesOneNode(t *testing.T) {
	node := &huffNode{char: 120, count: 10}
	nodes := []*huffNode{node}
	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

	if root.char != 120 || root.count != 10 {
		t.Error("Unexpected! Got node with count:", root.count, "and char:",
			root.char, "instead of count:", node.count, "and char:", node.char)
	}
}

func TestMakeTreeFromNodesBasicTree(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 2}, {char: 120, count: 2}}
	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

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
	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

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

////////////////////////////////////////////////////////////////////////////////
// makeTreeFromText tests
////////////////////////////////////////////////////////////////////////////////

func TestMakeTreeFromTextEmpty(t *testing.T) {
	filename := string(rand.Int63())
	b := make([]byte, 0)
	err := ioutil.WriteFile(filename, b, 0644)
	errorIfNecessary(t, err)
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	if err == nil {
		t.Error("Got nil error from makeTreeFromText! Should be 'Text file empty'")
	}

	if root != nil {
		t.Error("Got non-nil tree.")
	}
}

func TestMakeTreeFromTextSingleChar(t *testing.T) {
	filename := string(rand.Int63())
	b := []byte{0}
	err := ioutil.WriteFile(filename, b, 0644)
	errorIfNecessary(t, err)
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	errorIfNecessary(t, err)

	if root.count != 1 || root.char != 0 {
		t.Error("Tree was built improperly! Expected: { char: 0, count: 1 },",
			"got { char:", root.char, ", count:", root.count, "}")
	}
}

func TestMakeTreeFromTextTwoOfSameChar(t *testing.T) {
	filename := string(rand.Int63())
	b := []byte{0, 0}
	err := ioutil.WriteFile(filename, b, 0644)
	errorIfNecessary(t, err)
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	errorIfNecessary(t, err)

	if root.count != 2 || root.char != 0 {
		t.Error("Tree was built improperly! Expected: { char: 0, count: 2 },",
			"got { char:", root.char, ", count:", root.count, "}")
	}
}

func TestMakeTreeFromTextBasicTree(t *testing.T) {
	filename := string(rand.Int63())
	b := []byte{0, 0, 2}
	err := ioutil.WriteFile(filename, b, 0644)
	errorIfNecessary(t, err)
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	errorIfNecessary(t, err)

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
	filename := string(rand.Int63())
	b := []byte{0, 0, 1, 1, 2, 2, 2}
	err := ioutil.WriteFile(filename, b, 0644)
	errorIfNecessary(t, err)
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	errorIfNecessary(t, err)

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
// getByteMap tests
////////////////////////////////////////////////////////////////////////////////

func TestGetByteMapOneByte(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 1}}
	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

	bytes := root.getByteMap()

	if len(bytes) != 1 {
		t.Error("Wrong number of bytes in byteCount map:", len(bytes))
	}

	if bytes[120] != "" {
		t.Error("Got non-empty string for only byte:", bytes[120])
	}
}

func TestGetByteMapTwoBytes(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 1},
		{char: 121, count: 2}}
	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

	bytes := root.getByteMap()
	if len(bytes) != 2 {
		t.Error("Wrong number of bytes in byteCount map:", len(bytes))
	}

	if bytes[120] != "0" {
		t.Error("Wrong bit pattern for byte 120:", bytes[120])
	}

	if bytes[121] != "1" {
		t.Error("Wrong bit pattern for byte 121:", bytes[121])
	}
}

func TestGetByteMapSeveralBytes(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 1},
		{char: 121, count: 2},
		{char: 122, count: 3}}
	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

	bytes := root.getByteMap()
	if len(bytes) != 3 {
		t.Error("Wrong number of bytes in byteCount map:", len(bytes))
	}

	if bytes[122] != "0" {
		t.Error("Wrong bit pattern for byte 122. Got:", bytes[122])
	}

	if bytes[121] != "11" {
		t.Error("Wrong bit pattern for byte 121. Got:", bytes[121])
	}

	if bytes[120] != "10" {
		t.Error("Wrong bit pattern for byte 120. Got:", bytes[120])
	}
}

func TestGetByteMapManyBytes(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 1},
		{char: 121, count: 2},
		{char: 122, count: 4},
		{char: 123, count: 8},
		{char: 124, count: 16}}

	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

	bytes := root.getByteMap()
	if len(bytes) != 5 {
		t.Error("Wrong number of bytes in byteCount map:", len(bytes))
	}

	if bytes[120] != "0000" {
		t.Error("Wrong bit pattern for byte 122. Got:", bytes[120])
	}

	if bytes[121] != "0001" {
		t.Error("Wrong bit pattern for byte 121. Got:", bytes[121])
	}

	if bytes[122] != "001" {
		t.Error("Wrong bit pattern for byte 120. Got:", bytes[122])
	}

	if bytes[123] != "01" {
		t.Error("Wrong bit pattern for byte 121. Got:", bytes[123])
	}

	if bytes[124] != "1" {
		t.Error("Wrong bit pattern for byte 120. Got:", bytes[124])
	}
}

////////////////////////////////////////////////////////////////////////////////
// writeToFile and makeTreeFromTreeFile tests
////////////////////////////////////////////////////////////////////////////////

func TestReadAndWriteToFileOneNodeTree(t *testing.T) {
	filename := string(rand.Int63())
	nodes := []*huffNode{{char: 120, count: 1}}

	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

	file, err := os.Create(filename)
	defer os.Remove(filename)
	errorIfNecessary(t, err)

	err = root.writeToFile(file)
	errorIfNecessary(t, err)
	err = file.Close()
	errorIfNecessary(t, err)

	bytes, err := ioutil.ReadFile(filename)
	errorIfNecessary(t, err)
	if len(bytes) != 4 {
		t.Error("Wrong file size. Wanted 4 bytes, got", len(bytes))
	}
	if bytes[0]+1 != 1 {
		t.Error("Wrong number of bytes. Wanted 1, got", bytes[0]+1)
	}
	if bytes[1] != 120 {
		t.Error("Wrong byte written. Wanted 120, got", bytes[1])
	}
	if bytes[2] != 0 {
		t.Error("Wrong bitrep length. Wanted 0, got", bytes[2])
	}
	if bytes[3] != 0 {
		t.Error("Wrong bitrep. Should have been 0 (all padding), got", bytes[3])
	}

	file, err = os.Open(filename)
	errorIfNecessary(t, err)
	newRoot, err := makeTreeFromTreeFile(file)
	errorIfNecessary(t, err)

	if !equal(root, newRoot) {
		t.Error("Something went wrong creating the new tree.")
	}
}

func TestReadAndWriteToFileBasic(t *testing.T) {
	filename := string(rand.Int63())
	nodes := []*huffNode{{char: 120, count: 1},
		{char: 121, count: 2}}

	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

	file, err := os.Create(filename)
	defer os.Remove(filename)
	errorIfNecessary(t, err)

	err = root.writeToFile(file)
	errorIfNecessary(t, err)
	err = file.Close()
	errorIfNecessary(t, err)

	bytes, err := ioutil.ReadFile(filename)
	errorIfNecessary(t, err)
	if len(bytes) != 7 {
		t.Error("Wrong file size. Wanted 7 bytes, got", len(bytes))
	}
	if bytes[0]+1 != 2 {
		t.Error("Wrong number of bytes! Wanted 2, got", bytes[0]+1)
	}

	file, err = os.Open(filename)
	errorIfNecessary(t, err)
	newRoot, err := makeTreeFromTreeFile(file)
	errorIfNecessary(t, err)

	if !equal(root, newRoot) {
		t.Error("Something went wrong creating the new tree!")
	}
}

func TestReadAndWriteToFileMoreAdvanced(t *testing.T) {
	filename := string(rand.Int63())

	nodes := []*huffNode{{char: 120, count: 1},
		{char: 121, count: 2},
		{char: 122, count: 4},
		{char: 123, count: 8},
		{char: 124, count: 16},
		{char: 125, count: 32},
		{char: 126, count: 64},
		{char: 127, count: 128},
		{char: 128, count: 256}}

	root, err := makeTreeFromNodeSlice(nodes)
	errorIfNecessary(t, err)

	file, err := os.Create(filename)
	defer os.Remove(filename)
	errorIfNecessary(t, err)

	err = root.writeToFile(file)
	errorIfNecessary(t, err)
	err = file.Close()
	errorIfNecessary(t, err)

	bytes, err := ioutil.ReadFile(filename)
	errorIfNecessary(t, err)

	if len(bytes) != len(nodes)*3+1 {
		t.Error("Wrong file size. Wanted", len(nodes)*3+1, ", got", len(bytes))
	}
	if bytes[0]+1 != 9 {
		t.Error("Wrong size written. Wanted 9, got", bytes[0]+1)
	}

	file, err = os.Open(filename)
	errorIfNecessary(t, err)
	newRoot, err := makeTreeFromTreeFile(file)
	errorIfNecessary(t, err)

	if !equal(root, newRoot) {
		t.Error("Something went wrong creating the new tree!")
	}
}

////////////////////////////////////////////////////////////////////////////////
// writeEncodedText tests
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// writeDecodedText tests
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// Full-on tests. Encode, then decode several files of varying complexity
////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////
// helper functions
////////////////////////////////////////////////////////////////////////////////

// If err is non-nil, will call t.Error(err). Just for code clean-ness
func errorIfNecessary(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

// Checks if the passed two trees are equal--that is, if they were set on the same
// encoded text, they would produce the exact some decoded text.
func equal(t1 *huffNode, t2 *huffNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil && t2 != nil {
		return false
	} else if t2 == nil && t1 != nil {
		return false
	} else {
		return equal(t1.left, t2.left) &&
			equal(t1.right, t2.right) &&
			t1.char == t2.char &&
			t1.count == t2.count
	}
}
