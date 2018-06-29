package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// todo:
// how to figure out if a byte is an alphabetical character?
// if it is, how do we push it 13 places next in the alphabet?

func (r rot13Reader) Read(b []byte) (int, error) {
	// modify the stream by rotating it by 13 places
	for i := range b {
		var byteInQuestion = b[i]
		// we need to convert this byte
		// to the correct one... but only if it is an alphabetical character
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
