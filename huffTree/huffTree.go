// Ben Eggers
// GNU GPL'd

package huffTree

import (
	"container/heap"
	"errors"
	"github.com/BenedictEggers/bitIO"
	"io/ioutil"
	"os"
)

// The actual Huffman Tree and all associated functions. Will build up a
// tree from either a file defining the text file to be encoded or a file
// defining the Huffman Tree (see makeTreeFromText(), makeTreeFromTreeFile()) and
// will write a tree out to disk (see tree.writeToFile()). Uses the node struct found
// in "./node.go".

// HuffTree is really just a pointer to the root node of the tree.
type HuffTree struct {
	root *huffNode
}

////////////////////////////////////////////////////////////////////////////////
//               Stuff to encode a file
////////////////////////////////////////////////////////////////////////////////

// EncodeText turns the bytes in fromFile into bytes in toFile, compressed under
// a tree created on the file. On success, returns a nil error and returns a
// non-nil error otherwise.
func EncodeText(fromFile, toFile string) (err error) {
	// Make a tree from the file
	tree, err := makeTreeFromText(fromFile)
	if err != nil {
		return err
	}

	// Open up our file to write
	openFile, err := os.Create(toFile)
	if err != nil {
		return err
	}

	// Write the tree to the file
	err = tree.writeToFile(openFile)
	if err != nil {
		return err
	}

	// Encode the actual stuff and write it out
	err = t.writeEncodedText(fromFile, openFile)
	if err != nil {
		return err
	}

	// Then return (writeEncodedText closes the file, so we won't deal with it)
	return nil
}

// makeTreeFromText takes in a text file and turns it into a HuffTree, which
// it then returns. Will be called by EncodeText to build the tree before
// trying to encode the text.
func makeTreeFromText(filename string) (t *HuffTree, err error) {
	// Read the text byte-by-byte, building up a map of byte counts
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Scan the byte slice "buf" and count how many times each byte shows up
	counts := map[byte]uint32{}
	for _, elem := range buf {
		counts[elem] += 1
	}

	// Turn the counts into huffNodes
	nodes := make([]*huffNode, 0)
	for currentByte, byteCount := range counts {
		node := &huffNode{char: currentByte, count: byteCount}
		nodes = append(nodes, node)
	}

	if len(nodes) == 0 {
		return nil, errors.New("Invalid node slice.")
	}

	return makeTreeFromNodeSlice(nodes)
}

// Makes a tree from a slice of huffNodes, with the lowest-count nodes
// going farthest from the root. Returns a non-nil error on failure, nil
// error otherwise. Returns the created tree. If nodes is empty, returns
// an error.
func makeTreeFromNodeSlice(nodes []*huffNode) (t *HuffTree, err error) {
	if len(nodes) == 0 {
		return nil, errors.New("Too few elements!")
	}
	// We're going to put the nodes in a heap, with low-ness determined
	// by the nodes' counts.
	nh := &nodeHeap{}
	heap.Init(nh)
	for _, node := range nodes {
		heap.Push(nh, node)
	}

	// Now, we're going to do the following:
	//
	// Until there's only one node in the heap:
	// 		Remove the lowest-count two nodes
	// 		Make a new node with those two as children, whose count is the
	//			sum of its childrens' counts
	//		Add that new node to the heap
	//
	// This will create an optimally-balanced tree, based on byte counts. For
	// more information, see http://en.wikipedia.org/wiki/Huffman_coding.
	for nh.Len() > 1 {
		nodeOne := heap.Pop(nh).(*huffNode)
		nodeTwo := heap.Pop(nh).(*huffNode)
		newNode := &huffNode{char: 255, // random char
			count: nodeOne.count + nodeTwo.count,
			left:  nodeOne,
			right: nodeTwo}
		heap.Push(nh, newNode)
	}

	// Great, now there's only one node and it's the root of the tree!
	return &HuffTree{heap.Pop(nh).(*huffNode)}, nil
}

// Write the tree out to a file at a file described by the passed string.
// Will be called by EncodeText to write the tree out to the beginning
// of the encoded file.
func (t *HuffTree) writeToFile(file *os.File) (err error) {
	return errors.New("Undefined method")
}

// writeEncodedTextToFile encodes the text in the passed file under the HuffTree
// it was called on, and writes out the encoded bits to the passed file. Is called
// by EncodeText. Returns a non-nil error on failure, nil otherwise.
func (t *HuffTree) writeEncodedText(fromFile string, toFile *os.File) (err error) {
	return errors.New("Not yet implemented")
}

////////////////////////////////////////////////////////////////////////////////
//               Stuff to decode a file
////////////////////////////////////////////////////////////////////////////////

// DecodeText turns the bytes in fromFile into bytes in toFile, decompressed under
// the tree it is called on. On success, returns a nil error and returns a
// non-nil error otherwise. If fromFile exists before the call, it is deleted
// and replaced with the decompressed file.
func DecodeText(fromFile, toFile string) (err error) {
	// Open up the encoded file
	encoded, err := os.Open(fromFile)
	if err != nil {
		return err
	}

	// Make the tree
	t, err := makeTreeFromTreeFile(encoded)
	if err != nil {
		return err
	}

	// And write it
	return t.writeDecodedText(encoded, toFile)
}

// makeTreeFromTreeFile takes in a filname of a file in the same format TREE.writeToFile()
// puts out, and remakes a HuffTree from it.
func makeTreeFromTreeFile(file *os.File) (t *HuffTree, err error) {
	return &HuffTree{}, errors.New("Undefined method")
}

func (t *HuffTree) writeDecodedText(fromFile *os.File, toFile string) (err error) {
	// Set up a BitReader on the file to decodes
	reader, err := bitIO.NewReader(toFile)
	if err != nil {
		return err
	}

	// And open up our file to write to
	outFile, err := os.Create(toFile)
	if err != nil {
		return err
	}

	// Decode our bits, writing them out to disk every 1,000 bytes so as not
	// to use up all of main memory
	toWrite := make([]byte, 1000)
	bytesWritten := 0
	current := t.root

	// until we reach the end of the file...
	for current.char != 0 {

		// Read a bit
		bit, err := reader.ReadBit()
		if err != nil {
			return err
		}

		if current.left == nil && current.right == nil {
			// We're at a leaf node, write out its character
			toWrite[bytesWritten] = current.char
			bytesWritten++
			current = t.root
		}

		if bytesWritten == len(toWrite) {
			// Time to flush our buffer and reset
			_, err = outFile.Write(toWrite)
			if err != nil {
				return nil
			}
		}

		if bit == 0 {
			current = current.left
		} else if bit == 1 {
			current = current.right
		} else {
			// Should never happen
			return errors.New("Got invalid bit")
		}
	}

	// We've terminated, but might still need to write some bytes
	if bytesWritten != len(toWrite) {
		_, err = outFile.Write(toWrite)
		if err != nil {
			return err
		}
	}

	// Great, we made it
	return nil
}
