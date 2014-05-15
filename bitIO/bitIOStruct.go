// Ben Eggers
// GNU GPL'd

package bitIO

// A struct useful for both bitReader and bitWriter

type bitIOStruct struct {
	bits    []uint8 // should ALWAYS have length 1
	numBits uint8
	file    *File
}

// Make a bitIOStruct with the passed parameters

func makeBitIOStruct(file *File) (b bitIOStruct) {
	b.file = file
	b.bits = make([]int8, 1)
	b.numBits = 8
}