// Ben Eggers
// GNU GPL'd

// This package provides abstractions to read or write single bytes to and
// from disk. While a BitReader or BitWriter is doing stuff with a file,
// you shouldn't mess with it.
package bitIO

type BitReader struct {
	bitIOStruct
}

// Set up and return a BitReader on the passed file.
func MakeBitReader(file string) (b BitReader, err error) {
	str, err := makeBitIOStruct(file)
	return BitReader{str}, err
}

// Returns the next bit on the file stream. Will always be 0 or 1. Will
// return a non-nil err iff the read failed, or on EOF
func (b BitReader) ReadBit() (bit byte, err error) {
	bit = b.bits[0] % 2
	b.bits[0] = b.bits[0] >> 1
	b.numBits += 1
	if b.numBits == 8 {
		// we need the next byte!
		_, err = b.file.Read(b.bits)
		if err != nil {
			return 0, err
		}
		b.numBits = 0
	}
	return bit, nil
}

// Closes the reader, closing its associated file descriptor
func (b BitReader) Close() (err error) {
	return b.file.Close()
}
