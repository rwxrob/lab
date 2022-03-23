package main

import "fmt"

type Foo struct {
	Some string
}

// nope, constants cannot be structs
// const MyFoo = Foo{"something"}
// const MyFoo = struct{ Foo string }{"something"}

func main() {
	fmt.Println("vim-go")
}
