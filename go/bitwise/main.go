package main

import "fmt"

const (
	asleep = 1 << iota
	dream
	nightmare
	walking
	_
	_
	_
	lucid
)

func main() {
	flags := 0
	flags = asleep | dream | lucid
	fmt.Printf("%010b\n", flags)
	flags ^= dream
	fmt.Printf("%010b\n", flags)
	flags |= dream
	fmt.Printf("%010b\n", flags)

	/*
		flags |= asleep
		fmt.Printf("%b\n", flags)
		flags &= asleep
		fmt.Printf("%b\n", flags)
		flags |= lucid
		fmt.Printf("%b\n", flags)
		flags >>= 3
		fmt.Printf("%b\n", flags)
		flags <<= 3
		fmt.Printf("%b\n", flags)
	*/
}
