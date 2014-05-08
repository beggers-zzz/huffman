// Ben Eggers
// GNU GPL'd

// Heap for use with the Huffman Tree
package huffman

type node struct {
	char  rune
	freq  float64
	left  *node
	right *node
}

type nodeHeap []*node

func (nh nodeHeap) Len() int { return len(nh) }

func (nh nodeHeap) Less(i, j int) bool {
	return nh[i].freq < nh[j].freq
}

func (nh nodeHeap) Swap(i, j int) {
	nh[i], nh[j] = nh[j], nh[i]
}

func (nh *nodeHeap) Push(x interface{}) {
	*nh = append(*nh, x.(*node))
}

func (nh *nodeHeap) Pop() interface{} {
	old := *nh
	n := len(old)
	node := old[n-1]
	*nh = old[0 : n-1]
	return node
}
