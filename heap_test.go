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
	toPush := &huffNode{char: 'a', freq: .004}
	heap.Push(nh, toPush)
	popped := heap.Pop(nh).(*huffNode)
	if popped.char != toPush.char || popped.freq != toPush.freq {
		t.Error("Different frequencies or chars!")
	}
}

func TestAddTwoValues(t *testing.T) {
	nh := &nodeHeap{}
	heap.Init(nh)
	nodeOne := &huffNode{char: 'x', freq: .1}
	nodeTwo := &huffNode{char: 'x', freq: .2}
	heap.Push(nh, nodeTwo)
	heap.Push(nh, nodeOne)
	popped := heap.Pop(nh).(*huffNode)
	if popped.freq != .1 {
		t.Error("Should be .1, got ", popped.freq, ".")
	}
	popped = heap.Pop(nh).(*huffNode)
	if popped.freq != .2 {
		t.Error("Should be .2, got ", popped.freq, ".")
	}
}

func TestMultipleValuesAscending(t *testing.T) {
	freqs := []float64{.1, .2, .3, .4, .5}
	nh := &nodeHeap{}
	heap.Init(nh)
	for _, f := range freqs {
		n := &huffNode{char: 'a', freq: f}
		heap.Push(nh, n)
	}
	for _, f := range freqs {
		n := heap.Pop(nh).(*huffNode)
		if n.freq != f {
			t.Error("Wanted node with freq: ", f,
				", got node with freq: ", n.freq, ".")
		}
	}
}

func TestMultipleValuesDescending(t *testing.T) {
	freqs := []float64{.1, .2, .3, .4, .5}
	nh := &nodeHeap{}
	heap.Init(nh)
	for i := range freqs {
		n := &huffNode{char: 'a', freq: freqs[len(freqs)-1-i]}
		heap.Push(nh, n)
	}
	for _, f := range freqs {
		n := heap.Pop(nh).(*huffNode)
		if n.freq != f {
			t.Error("Wanted node with freq: ", f,
				", got node with freq: ", n.freq, ".")
		}
	}
}

func TestMultipleValuesRandomOrder(t *testing.T) {
	freqs := []float64{.4, .2, .1, .5, .3}
	freqsSorted := []float64{.1, .2, .3, .4, .5}
	nh := &nodeHeap{}
	heap.Init(nh)
	for i := range freqs {
		n := &huffNode{char: 'a', freq: freqs[len(freqs)-1-i]}
		heap.Push(nh, n)
	}
	for _, f := range freqsSorted {
		n := heap.Pop(nh).(*huffNode)
		if n.freq != f {
			t.Error("Wanted node with freq: ", f,
				", got node with freq: ", n.freq, ".")
		}
	}
}
