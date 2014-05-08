// Ben Eggers
// GNU GPL'd

package huffman

import (
	"io/ioutil"
)

// The actual Huffman Tree and all associated functions. Will build up a
// tree from either a file defining the text file to be encoded or a file
// defining the Huffman Tree (see makeTreeFromText(), makeTreeFromTreeFile()) and
// will write a tree out to disk (see tree.writeToFile()). Uses the node struct found
// in "./node.go".

// huffTree is really just a pointer to the root node of the tree.
type huffTree *huffNode

// decode decodes the passed File using the huffTree it is called on. Returns
// a string representing the decoded File. If the decode fails (wrong tree, some
// file read error, etc) the value of the returned string is undefined but error
// will be non-nil.
func (t huffTree) decode(filename string) decoded string, err error {

}

// encode turns the bytes in fromFile into bytes in toFile, compressed under
// the tree it is called on. On success, returns a nil error and returns a 
// non-nill error otherwise.
func (t huffTree) encode(fromFile string, toFile string) err error {

}

// Write the tree out to a file at the point the file is currently seek-ed to.
// Necessary to keep around if you plan on decoding files :)
func (t huffTree) writeToFile(file *File) err error {

}

////////////////////////////////////////////////////////////////////////////////
//               Functions to help with making the tree.
////////////////////////////////////////////////////////////////////////////////


// makeTreeFromText takes in a text file and turns it into a huffTree, which
// it then returns.
func makeTreeFromText(filename string) (t huffTree, err error) {
	// Read the text byte-by-byte, building up a map of byte counts
	buf, err := readFile(filename)
	if err != nil {
		// Something went wrong!
		return nil, err
	}

	// Scan the byte slice "buf" and count how many times each byte shows up
	counts := map[byte]uint32 
	for _, elem := range buf {
		counts[elem] += 1
	}

	// Turn the counts into huffNodes
	nodes := make([]*huffNode, 0)
	for currentByte, byteCount := range counts {
		node := &huffNode{char: currentByte, count: byteCount}
		nodes = append(nodes, node)
	}

	return makeTreeFromNodeSlice(nodes)
}

// makeTreeFromTreeFile takes in a File of the same format TREE.writeToFile()
// puts out, and remakes a huffTree from it.  Only reads the few bytes required
// to make the tree, so you can put it anywhere in the File as long as the passed
// *File is pointing at the beginning of the tree. Returns a non-nil error iff
// the tree formation fails.
func makeTreeFromTreeFile(filename string) (t huffTree, err error) {

}

func makeTreeFromNodeSlice(nodes []*huffNode) (t huffTree, err error) {

}