// Ben Eggers
// GNU GPL'd

package bitIO

import "os"

// A struct useful for both bitReader and bitWriter
type bitIOStruct struct {
	bits    []uint8 // should ALWAYS have length 1
	numBits uint8
	file    *File
}

// Make a bitIOStruct on the passed File descriptor
func makeBitIOStruct(file *File) (b bitIOStruct) {
	b.file = file
	b.bits = make([]int8, 1)
	b.numBits = 0
}
