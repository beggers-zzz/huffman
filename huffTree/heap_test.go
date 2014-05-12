// Ben Eggers
// GNU GPL'd

// Tests the heap
package huffTree

import (
	"container/heap"
	"testing"
)

func TestSimpleAdd(t *testing.T) {
	nh := &nodeHeap{}
	heap.Init(nh)
	toPush := &huffNode{char: 'a', count: 4}
	heap.Push(nh, toPush)
	popped := heap.Pop(nh).(*huffNode)
	if popped.char != toPush.char || popped.count != toPush.count {
		t.Error("Different count or chars!")
	}
}

func TestAddTwoValues(t *testing.T) {
	nh := &nodeHeap{}
	heap.Init(nh)
	nodeOne := &huffNode{char: 'x', count: 1}
	nodeTwo := &huffNode{char: 'x', count: 2}
	heap.Push(nh, nodeTwo)
	heap.Push(nh, nodeOne)
	popped := heap.Pop(nh).(*huffNode)
	if popped.count != 1 {
		t.Error("Should be 1, got ", popped.count, ".")
	}
	popped = heap.Pop(nh).(*huffNode)
	if popped.count != 2 {
		t.Error("Should be 2, got ", popped.count, ".")
	}
}

func TestMultipleValuesAscending(t *testing.T) {
	counts := []uint32{1, 2, 3, 4, 5}
	nh := &nodeHeap{}
	heap.Init(nh)
	for _, f := range counts {
		n := &huffNode{char: 'a', count: f}
		heap.Push(nh, n)
	}
	for _, f := range counts {
		n := heap.Pop(nh).(*huffNode)
		if n.count != f {
			t.Error("Wanted node with count: ", f,
				", got node with count: ", n.count, ".")
		}
	}
}

func TestMultipleValuesDescending(t *testing.T) {
	counts := []uint32{1, 2, 3, 4, 5}
	nh := &nodeHeap{}
	heap.Init(nh)
	for i := range counts {
		n := &huffNode{char: 'a', count: counts[len(counts)-1-i]}
		heap.Push(nh, n)
	}
	for _, f := range counts {
		n := heap.Pop(nh).(*huffNode)
		if n.count != f {
			t.Error("Wanted node with count: ", f,
				", got node with count: ", n.count, ".")
		}
	}
}

func TestMultipleValuesRandomOrder(t *testing.T) {
	counts := []uint32{4, 2, 1, 5, 3}
	countsSorted := []uint32{1, 2, 3, 4, 5}
	nh := &nodeHeap{}
	heap.Init(nh)
	for i := range counts {
		n := &huffNode{char: 'a', count: counts[len(counts)-1-i]}
		heap.Push(nh, n)
	}
	for _, f := range countsSorted {
		n := heap.Pop(nh).(*huffNode)
		if n.count != f {
			t.Error("Wanted node with count: ", f,
				", got node with count: ", n.count, ".")
		}
	}
}
