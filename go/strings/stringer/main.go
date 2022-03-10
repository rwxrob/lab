package main

import "fmt"

type Foo struct{}

func (f *Foo) String() string { return "FOOO" }

// does not care if pointer or not

type Bar struct{}

func (f Bar) String() string { return "BAAR" }

func main() {
	fmt.Printf("%v\n", new(Foo))
	fmt.Printf("%v\n", new(Bar))
}
