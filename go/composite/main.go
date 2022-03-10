package main

import "fmt"

type P struct {
	T string
	V []string
}

func main() {
	p := P{"T", "some", "thing"...}
	fmt.Println(p)
}
