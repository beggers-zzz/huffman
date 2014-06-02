// Ben Eggers
// GNU GPL'd

package huffTree

import (
	"bytes"
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

// Will be written at the beginning of every encoded file for integrity check
var magicBytes = [...]byte{'m', 'o', 'o', 'o', 's', 'e'}

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

	// Set up the write cursor to be in the correct position
	_, err = openFile.Seek(int64(len(magicBytes)), 0)
	if err != nil {
		return err
	}

	// Write the tree to the file
	err = tree.writeToFile(openFile)
	if err != nil {
		return err
	}

	// Encode the actual stuff and write it out
	err = tree.writeEncodedText(fromFile, openFile)
	if err != nil {
		return err
	}

	// Now write our "magic bytes" at the beginning so we can check
	// when Decode() is called if the tree is valid
	_, err = openFile.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = openFile.Write(magicBytes[:])
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

// getByteMap returns a map from all the bytes in the tree onto strings, which will
// be entirely 0s and 1s (in string form). If a byte b maps to a string s, that
// means that the encoded representation of b will be s, but as bytes, not as
// a string.
func (t *HuffTree) getByteMap() (characters map[byte]string, err error) {
	err = getByteMapRecursiveHelper(t.root, "", characters)
	return characters, err
}

// Helper function for getByteMap(). See above/below
func getByteMapRecursiveHelper(cur *huffNode, soFarStr string,
	soFarMap map[byte]string) (err error) {
	// We're going to do a pre-order traversal of the tree, building up (and
	// sometimes, tearing down) a string--it will be 0 if we went left, 1 if
	// we went right. When we reach a leaf node, we'll add it's character to the
	// map, mapping to the current string

	// are we at a leaf node?
	if cur.right == nil && cur.left == nil {
		// yep
		soFarMap[cur.char] = soFarStr
		return nil
	}

	// Nope, need to keep recursing. First set up our map
	if cur.left != nil {
		err = getByteMapRecursiveHelper(cur.left, soFarStr+"0", soFarMap)
		if err != nil {
			// oh no!
			return err
		}
	}

	if cur.right != nil {
		err = getByteMapRecursiveHelper(cur.right, soFarStr+"1", soFarMap)
		if err != nil {
			// oh no!
			return err
		}
	}

	return nil
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

	integrityCheck := make([]byte, len(magicBytes)) // to store the first few bytes
	_, err = encoded.Read(integrityCheck)
	if err != nil {
		return err
	}

	if !bytes.Equal(integrityCheck, magicBytes[:]) {
		// File is corrupted
		return errors.New("Corrupted file")
	}

	// Make the tree
	t, err := makeTreeFromTreeFile(encoded)
	if err != nil {
		return err
	}

	// And decode the rest of the file
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
