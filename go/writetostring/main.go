package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	buf := new(bytes.Buffer)
	log.SetFlags(0)
	log.SetOutput(buf)
	log.Print("foo")
	fmt.Print(buf.String())
}
