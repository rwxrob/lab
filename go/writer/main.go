package main

import (
	"io"
	"os"
)

// os os.Stdout DOES implement io.Writer

func main() {
	var b io.Writer = os.Stdout
	out := []byte{'s', 'o'}
	b.Write(out)
}
