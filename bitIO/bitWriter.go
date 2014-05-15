// Ben Eggers
// GNU GPL'd

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

// Writes one bit. If the passed int8 is 1, writes a one. If it's 0,
// writes a 0. Else, returns a non-nil error.
func (b BitWriter) WriteBit(bit int8) (err error) {
	if bit != 0 && bit != 1 {
		return errors.New("Invalid bit to write.")
	}

	if b.numBits == 8 {
		err = b.flush()
		if err != nil {
			return err
		}
	}

	b.bits[0] += bit << b.numBits
	b.numBits++
	return nil
}

// Flushes the current byte out to disk, padding with 0s if necessary.
func (b BitReader) flush() (err error) {
	for b.numBits != 8 {
		b.WriteBit(0)
	}
	_, err = file.Write(bits)
	return err
}

// Closes the BitReader, flushing final bits to disk if need be and closing
// the file descriptor.
func (b BitReader) Close() (err error) {
	err = b.flush()
	if err != nil {
		return err
	}
	return file.Close()
}
