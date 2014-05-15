// Ben Eggers
// GNU GPL'd

// This package provides abstractions to read or write single bytes to and
// from disk. While a ByteReader or ByteWriter is doing stuff with a file,
// you shouldn't mess with it.
package bitIO

import (
	"error"
)

type BitWriter struct {
	bitIOStruct
}

// Set up and return a BitWriter on the passed file.
func MakeBitReader(file *File) (b BitReader) {
	return makeBitIOStruct(file)
}

// Writes one bit. If the passed int8 is non-zero, writes a one. Else,
// writes a 0. Returns an error if something goes wrong.
func (b BitWriter) WriteBit(bit int8) err error {
	if b.numBits == 8 {
		b.flush()
	}
	b.bits
}

func (b BitReader) flush() err error {
	
}