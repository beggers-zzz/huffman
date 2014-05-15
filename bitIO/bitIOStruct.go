// Ben Eggers
// GNU GPL'd

package bitIO

import "os"

// A struct useful for both bitReader and bitWriter
type bitIOStruct struct {
	bits    []byte // should ALWAYS have length 1
	numBits int8
	file    *os.File
}

// Make a bitIOStruct on the passed File descriptor
func makeBitIOStruct(file string) (b bitIOStruct, err error) {
	b.file, err = os.Open(file)
	if err != nil {
		return bitIOStruct{}, err
	}
	b.bits = make([]byte, 1)
	b.numBits = 0
	return b, nil
}
