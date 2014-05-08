// Ben Eggers
// GNU GPL'd

// Heap for use with the Huffman Tree
package huffman

type nodeHeap []*huffNode

func (nh nodeHeap) Len() int { return len(nh) }

func (nh nodeHeap) Less(i, j int) bool {
	return nh[i].count < nh[j].count
}

func (nh nodeHeap) Swap(i, j int) {
	nh[i], nh[j] = nh[j], nh[i]
}

func (nh *nodeHeap) Push(x interface{}) {
	*nh = append(*nh, x.(*huffNode))
}

func (nh *nodeHeap) Pop() interface{} {
	old := *nh
	n := len(old)
	huffNode := old[n-1]
	*nh = old[0 : n-1]
	return huffNode
}
