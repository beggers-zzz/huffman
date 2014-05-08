// Ben Eggers
// GNU GPL'd

// Node used in the Huffman Tree

package huffman

type huffNode struct {
	char  byte
	freq  float64
	left  *huffNode
	right *huffNode
}