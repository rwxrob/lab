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

	fmt.Println(flags&(lucid|dream) == lucid|dream)        // true
	fmt.Println(flags&(lucid|walking) == lucid|walking)    // false
	fmt.Println(flags|lucid == flags)                      // true
	fmt.Println(flags|lucid|dream == flags)                // true
	fmt.Println(flags|lucid|dream|asleep == flags)         // true
	fmt.Println(flags|lucid|dream|asleep|walking == flags) // false
	fmt.Println(flags|walking == flags)                    // false

	/*
		fmt.Printf("%010b\n", flags)
		flags ^= dream
		fmt.Printf("%010b\n", flags)
		flags |= dream
		fmt.Printf("%010b\n", flags)


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
