package main

import "fmt"

type Foo struct {
	Some string
}

func Bar() { fmt.Println("bar") }

var bbar = Bar

// nope, constants cannot be structs
// const MyFoo = Foo{"something"}
// const MyFoo = struct{ Foo string }{"something"}

func main() {
	bbar()
	fmt.Println("vim-go")
}
