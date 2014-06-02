// Ben Eggers
// GNU GPL'd

package huffTree

// Node used in the Huffman Tree
type huffNode struct {
	char        byte
	count       uint32
	left, right *huffNode
}
