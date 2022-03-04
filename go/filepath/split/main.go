package main

import (
	"log"
	"path/filepath"
)

func main() {
	dir, file := filepath.Split("foo")
	log.Printf("%q %q", dir, file)
	dir, file = filepath.Split("foo/some")
	log.Printf("%q %q", dir, file)
	dir, file = filepath.Split("foo/some thing")
	log.Printf("%q %q", dir, file)
}
