// Ben Eggers
// GNU GPL'd

package huffman

// Tests the huffTree. Tests are (ostensibly) ordered in the order the function
// tested depend on each other, so failures are apparent at the lowest level.

import (
	"bytes"
	"io/ioutil"
	"math"
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
	if err != nil {
		t.Error(err)
	}

	if root.char != 120 || root.count != 10 {
		t.Error("Unexpected! Got node with count:", root.count, "and char:",
			root.char, "instead of count:", node.count, "and char:", node.char)
	}
}

func TestMakeTreeFromNodesBasicTree(t *testing.T) {
	nodes := []*huffNode{{char: 120, count: 2}, {char: 120, count: 2}}
	root, err := makeTreeFromNodeSlice(nodes)
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}
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
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	if err != nil {
		t.Error(err)
	}

	if root.count != 1 || root.char != 0 {
		t.Error("Tree was built improperly! Expected: { char: 0, count: 1 },",
			"got { char:", root.char, ", count:", root.count, "}")
	}
}

func TestMakeTreeFromTextTwoOfSameChar(t *testing.T) {
	filename := string(rand.Int63())
	b := []byte{0, 0}
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	if err != nil {
		t.Error(err)
	}

	if root.count != 2 || root.char != 0 {
		t.Error("Tree was built improperly! Expected: { char: 0, count: 2 },",
			"got { char:", root.char, ", count:", root.count, "}")
	}
}

func TestMakeTreeFromTextBasicTree(t *testing.T) {
	filename := string(rand.Int63())
	b := []byte{0, 0, 2}
	err := ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filename)

	root, err := makeTreeFromText(filename)
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}

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
	if err != nil {
		t.Error(err)
	}

	file, err := os.Create(filename)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	err = root.writeToFile(file)
	if err != nil {
		t.Error(err)
	}
	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	file, err = os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	newRoot, err := makeTreeFromTreeFile(file)
	if err != nil {
		t.Error(err)
	}

	if !equal(root, newRoot) {
		t.Error("Something went wrong creating the new tree.")
	}
}

func TestReadAndWriteToFileBasic(t *testing.T) {
	filename := ".test"
	nodes := []*huffNode{{char: 120, count: 1},
		{char: 121, count: 2}}

	root, err := makeTreeFromNodeSlice(nodes)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Create(filename)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	err = root.writeToFile(file)
	if err != nil {
		t.Error(err)
	}
	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	file, err = os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	newRoot, err := makeTreeFromTreeFile(file)
	if err != nil {
		t.Error(err)
	}

	if !equal(root, newRoot) {
		t.Error("Something went wrong creating the new tree!")
	}
}

func TestReadAndWriteToFileMoreAdvanced(t *testing.T) {
	filename := string(rand.Int63())

	nodes := []*huffNode{}
	for i := 0; i < 8; i++ {
		nodes = append(nodes, &huffNode{char: byte(i + 120),
			count: uint32(math.Pow(2.0, float64(i)))})
	}

	root, err := makeTreeFromNodeSlice(nodes)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Create(filename)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	err = root.writeToFile(file)
	if err != nil {
		t.Error(err)
	}
	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	file, err = os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	newRoot, err := makeTreeFromTreeFile(file)
	if err != nil {
		t.Error(err)
	}

	if !equal(root, newRoot) {
		t.Error("Something went wrong creating the new tree!")
	}
}

func TestReadAndWriteToFileExpertMode(t *testing.T) {
	filename := string(rand.Int63())

	nodes := []*huffNode{}
	for i := 0; i < 30; i++ {
		nodes = append(nodes, &huffNode{char: byte(i + 120),
			count: uint32(5 * i)})
	}

	root, err := makeTreeFromNodeSlice(nodes)
	if err != nil {
		t.Error(err)
	}

	file, err := os.Create(filename)
	defer os.Remove(filename)
	if err != nil {
		t.Error(err)
	}

	err = root.writeToFile(file)
	if err != nil {
		t.Error(err)
	}
	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	file, err = os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	newRoot, err := makeTreeFromTreeFile(file)
	if err != nil {
		t.Error(err)
	}

	if !equal(root, newRoot) {
		t.Error("Something went wrong creating the new tree!")
	}
}

////////////////////////////////////////////////////////////////////////////////
// Full-on tests. Encode, then decode several files of varying complexity
////////////////////////////////////////////////////////////////////////////////

func TestCompressLicense(t *testing.T) {
	compressed := ".license.com"
	decompressed := ".license.dec"
	err := EncodeText("./LICENSE", compressed)

	if err != nil {
		t.Error(err)
	}
	defer os.Remove(compressed)

	err = DecodeText(compressed, decompressed)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(decompressed)

	orig, err := ioutil.ReadFile("./LICENSE")
	if err != nil {
		t.Error(err)
	}

	other, err := ioutil.ReadFile(decompressed)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(orig, other) {
		t.Error("LICENSE compressed/decompressed improperly")
	}
}

func TestCompressReadme(t *testing.T) {
	compressed := ".readme.com"
	decompressed := ".readme.dec"
	err := EncodeText("./README.md", compressed)

	if err != nil {
		t.Error(err)
	}
	defer os.Remove(compressed)

	err = DecodeText(compressed, decompressed)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(decompressed)

	orig, err := ioutil.ReadFile("./README.md")
	if err != nil {
		t.Error(err)
	}

	other, err := ioutil.ReadFile(decompressed)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(orig, other) {
		t.Error("README.md compressed/decompressed improperly")
	}
}

func TestCompressHeap(t *testing.T) {
	compressed := ".heap.com"
	decompressed := ".heap.dec"
	err := EncodeText("./heap.go", compressed)

	if err != nil {
		t.Error(err)
	}
	defer os.Remove(compressed)

	err = DecodeText(compressed, decompressed)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(decompressed)

	orig, err := ioutil.ReadFile("./heap.go")
	if err != nil {
		t.Error(err)
	}

	other, err := ioutil.ReadFile(decompressed)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(orig, other) {
		t.Error("heap.go compressed/decompressed improperly")
	}
}

func TestCompressHeapTest(t *testing.T) {
	compressed := ".heap_test.com"
	decompressed := ".heap_test.dec"
	err := EncodeText("./heap_test.go", compressed)

	if err != nil {
		t.Error(err)
	}
	defer os.Remove(compressed)

	err = DecodeText(compressed, decompressed)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(decompressed)

	orig, err := ioutil.ReadFile("./heap_test.go")
	if err != nil {
		t.Error(err)
	}

	other, err := ioutil.ReadFile(decompressed)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(orig, other) {
		t.Error("heap_test.go compressed/decompressed improperly")
	}
}

func TestCompressHuffTree(t *testing.T) {
	compressed := ".huffTree.com"
	decompressed := ".huffTree.dec"
	err := EncodeText("./huffTree.go", compressed)

	if err != nil {
		t.Error(err)
	}
	defer os.Remove(compressed)

	err = DecodeText(compressed, decompressed)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(decompressed)

	orig, err := ioutil.ReadFile("./huffTree.go")
	if err != nil {
		t.Error(err)
	}

	other, err := ioutil.ReadFile(decompressed)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(orig, other) {
		t.Error("Incorrect compressing of huffTree.go")
	}
}

func TestCompressHuffTreeTest(t *testing.T) {
	// The irony
	compressed := ".huffTree_test.com"
	decompressed := ".huffTree_test.dec"
	err := EncodeText("./huffTree_test.go", compressed)

	if err != nil {
		t.Error(err)
	}
	defer os.Remove(compressed)

	err = DecodeText(compressed, decompressed)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(decompressed)

	orig, err := ioutil.ReadFile("./huffTree_test.go")
	if err != nil {
		t.Error(err)
	}

	other, err := ioutil.ReadFile(decompressed)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(orig, other) {
		t.Error("huffTree_test.go compressed/decompressed improperly")
	}
}

func TestCompressDecompressJaneEyre(t *testing.T) {
	compressed := ".janeeyre.com"
	decompressed := ".janeeyre.dec"
	err := EncodeText("./.test_docs/janeeyre.txt", compressed)
	defer os.Remove(compressed)
	if err != nil {
		t.Error(err)
	}

	err = DecodeText(compressed, decompressed)
	defer os.Remove(decompressed)
	if err != nil {
		t.Error(err)
	}

	orig, err := ioutil.ReadFile("./.test_docs/janeeyre.txt")
	if err != nil {
		t.Error(err)
	}

	other, err := ioutil.ReadFile(decompressed)
	if !bytes.Equal(orig, other) {
		t.Error("janeeyre compressed/decompressed improperly")
	}
}

// The ringer
func TestCompressDecompressBooks(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}

	// Compress and decompress all the books in the .test_docs folder
	files, err := ioutil.ReadDir("./.test_docs/")
	if err != nil {
		t.Error(err)
	}

	for _, fileInfo := range files {
		name := fileInfo.Name()
		err := EncodeText("./.test_docs/"+name, name+".enc")
		if err != nil {
			t.Error(err)
		}

		err = DecodeText(name+".enc", name+".dec")
		if err != nil {
			t.Error(err)
		}

		orig, err := ioutil.ReadFile("./.test_docs/" + name)
		if err != nil {
			t.Error(err)
		}

		other, err := ioutil.ReadFile(name + ".dec")
		if err != nil {
			t.Error(err)
		}

		if !bytes.Equal(orig, other) {
			t.Error(name + " compressed/decompressed improperly.")
		}
		os.Remove(name + ".enc")
		os.Remove(name + ".dec")
	}
}

////////////////////////////////////////////////////////////////////////////////
// helper functions
////////////////////////////////////////////////////////////////////////////////

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
			t1.char == t2.char
	}
}
