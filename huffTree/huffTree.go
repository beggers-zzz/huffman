// Ben Eggers
// GNU GPL'd

package huffTree

import (
	"bytes"
	"container/heap"
	"encoding/binary"
	"errors"
	"github.com/BenedictEggers/bitIO"
	"io/ioutil"
	"os"
)

// The actual Huffman Tree and all associated functions. Will build up a
// tree from either a file defining the text file to be encoded or a file
// defining the Huffman Tree (see makeTreeFromText(), makeTreeFromTreeFile()) and
// will write a tree out to disk (see tree.writeToFile()). The behavior is undefined
// if a file only has one type of character (e.g. all 'c's).

// The tree is made of huffNodes. There is no actual tree type, since the tree
// is only available internally.
type huffNode struct {
	char        byte
	count       uint32
	left, right *huffNode
}

// Will be written at the beginning of every encoded file for integrity check
var magicBytes = [...]byte{'m', 'o', 'o', 'o', 's', 'e'}

// Endian-ness to encode/decode with
var endianness = binary.LittleEndian

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

// makeTreeFromText takes in a text file and creats a huffman tree over the characters
// in it, which it then returns. Will be called by EncodeText to build the tree before
// trying to encode the text.
func makeTreeFromText(filename string) (t *huffNode, err error) {
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

// makeTreeFromNodeSlice makes a tree from a slice of huffNodes, with the
// lowest-count nodes going farthest from the root. Returns a non-nil error
// on failure, nil error otherwise. Returns the created tree. If nodes is empty,
// returns an error.
func makeTreeFromNodeSlice(nodes []*huffNode) (t *huffNode, err error) {
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
	return heap.Pop(nh).(*huffNode), nil
}

// writeToFile writes the tree out to the passed os.File.
// Will be called by EncodeText to write the tree out to the beginning
// of the encoded file.
func (t *huffNode) writeToFile(f *os.File) (err error) {
	// First, get the map of byte->string of 0s and 1s (character->binary representation)
	var bytes map[byte]string = t.getByteMap()

	// Then, write the number of bytes we have in the tree
	err = binary.Write(f, endianness, int8(len(bytes)-1))
	if err != nil {
		return err
	}

	// Now for each byte, we write:
	//		- The byte
	//		- The length of its binary (as a uint16)
	//		- The binary
	for char, repString := range bytes {
		// First the character
		_, err = f.Write([]byte{char})
		if err != nil {
			return err
		}

		// Now the length of the binary
		err = binary.Write(f, endianness, int8(len(repString)))
		if err != nil {
			return err
		}

		// And the actual bits
		bw, err := bitIO.NewWriterOnFile(f)
		if err != nil {
			return err
		}

		for _, c := range repString {
			err = bw.WriteBit(byte(c - '0')) // WriteBit() wants an int, we have runes
			if err != nil {
				return err
			}
		}
		_, err = bw.CloseAndReturnFile()
		if err != nil {
			return err
		}
	}

	return nil
}

// writeEncodedTextToFile encodes the text in the passed file under the tree
// it was called on, and writes out the encoded bits to the passed file. Is called
// by EncodeText. Returns a non-nil error on failure, nil otherwise.
func (t *huffNode) writeEncodedText(fromFile string, toFile *os.File) (err error) {
	return errors.New("Not yet implemented")
}

// getByteMap returns a map from all the bytes in the tree onto strings, which will
// be entirely 0s and 1s (in string form). If a byte b maps to a string s, that
// means that the encoded representation of b will be s, but as bytes, not as
// a string.
func (t *huffNode) getByteMap() (characters map[byte]string) {
	characters = make(map[byte]string)
	getByteMapRecursiveHelper(t, "", characters)
	return characters
}

// Helper function for getByteMap(). See above/below
func getByteMapRecursiveHelper(cur *huffNode, soFarStr string, soFarMap map[byte]string) {
	// We're going to do a pre-order traversal of the tree, building up (and
	// sometimes, tearing down) a string--it will be 0 if we went left, 1 if
	// we went right. When we reach a leaf node, we'll add it's character to the
	// map, mapping to the current string
	if cur == nil {
		// nothing to see here
		return
	}

	// are we at a leaf node?
	if cur.right == nil && cur.left == nil {
		// yep
		soFarMap[cur.char] = soFarStr
		return
	}

	// Nope, need to keep recursing.
	getByteMapRecursiveHelper(cur.left, soFarStr+"0", soFarMap)
	getByteMapRecursiveHelper(cur.right, soFarStr+"1", soFarMap)
}

////////////////////////////////////////////////////////////////////////////////
//               Stuff to decode a file
////////////////////////////////////////////////////////////////////////////////

// DecodeText turns the bytes in fromFile into bytes in toFile, decompressed under
// the tree at the beginning of the file. On success, returns a nil error. Else,
// returns a non-nil error. If fromFile exists before the call, it is deleted
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
	err = t.writeDecodedText(encoded, toFile)
	if err != nil {
		return err
	}

	// close the file, and return
	return encoded.Close()
}

// makeTreeFromTreeFile takes in a file in the same format TREE.writeToFile()
// puts out, and remakes a HuffTree from it.
func makeTreeFromTreeFile(file *os.File) (t *huffNode, err error) {
	return nil, errors.New("Undefined method")
}

// writeDecodedText decompresses the bits in the passed file, and puts the decompressed
// text into a new file described by toFile. If toFile exists before this is called,
// it will be truncated. Returns a nil error on success, non-nil error otherwise.
func (t *huffNode) writeDecodedText(fromFile *os.File, toFile string) (err error) {
	// Set up a BitReader on the file to decodes
	reader, err := bitIO.NewReader(toFile)
	if err != nil {
		return err
	}

	// Decode our bits
	toWrite := make([]byte, 0)
	current := t

	// until we reach the end of the file...
	for current.char != 0 {

		// Read a bit
		bit, err := reader.ReadBit()
		if err != nil {
			return err
		}

		if current.left == nil && current.right == nil {
			// We're at a leaf node, write out its character
			toWrite = append(toWrite, current.char)
			current = t
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

	// We've terminated, write it all out
	return ioutil.WriteFile(toFile, toWrite, 0644)
}
