// Ben Eggers
// GNU GPL'd

// Tests the heap
package huffman

import (
	"container/heap"
	"testing"
)

func TestSimpleAdd(t *testing.T) {
	nh := &nodeHeap{}
	heap.Init(nh)
	toPush := &node{char: 'a', frequency: .004}
	heap.Push(nh, toPush)
	popped := heap.Pop(nh).(*node)
	if popped.char != toPush.char || popped.frequency != toPush.frequency {
		t.Error("Different frequencies or chars!")
	}
}
