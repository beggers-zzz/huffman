// Ben Eggers
// GNU GPL'd

package main

import (
	"fmt"
	"github.com/BenedictEggers/huffman"
	"flag"
	"os"
)

// This will be the actual program to use the huffman tree to encode and decode
// files. It will accept a command-line argument giving it a file name, and
// flags to tell it what to with that file (encode, or decode).

func main() {
	var e bool
	var d bool
	flag.BoolVar(&e, "e", false, "Tells the program to encode")
	flag.BoolVar(&d, "d", false, "Tells the program to decode")
	flag.Parse()
	if !(e || d) || (e && d) {
		usage()
	}

	args := flag.Args()
	if len(args) != 2 {
		usage()
	}

	fromFile := args[0]
	toFile := args[1]
	var err error

	if e {
		err = huffman.EncodeText(fromFile, toFile)
	}
	if d {
		err = huffman.DecodeText(fromFile, toFile)
	}

	if err != nil {
		fmt.Println("Something went wrong:", err)
		os.Remove(toFile)
	}

	
}

func usage() {
	flag.Usage()
	fmt.Println("  FROM_FILE: File to read from (encoded if -e, plaintext if -d)")
	fmt.Println("  TO_FILE: File to write to (will be encoded if -e, decoded if -d)")
	fmt.Println("Either -e or -d should be flagged, not both. FROM_FILE and TO_FILE " +
		"must both be present.")
	os.Exit(1)
}