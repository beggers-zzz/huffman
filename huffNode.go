// Ben Eggers
// GNU GPL'd

package huffman

// Node used in the Huffman Tree
type huffNode struct {
	char  byte
	freq  uint32
	left  *huffNode
	right *huffNode
}