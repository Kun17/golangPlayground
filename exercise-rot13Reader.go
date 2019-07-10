package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte{
	var l, h byte
	switch {
		case 'a' <= b && b <= 'z':
			l, h = 'a', 'z'
		case 'A' <= b && b <= 'Z':
			l, h = 'A', 'Z'
		default:
			return b
	}
	return (b - l + 13)%(h-l+1) + l
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i:= range b{
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

