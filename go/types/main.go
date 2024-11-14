package main

import "fmt"

type Fooable interface {
	Foo()
}

type CmdFooable interface {
	Foo(x string)
}

type Thing struct{}

func (Thing) Foo() {}

type ThingContainer struct {
	Thing
}

func (ThingContainer) Foo(x string) {}

func PrintType(an any) {
	switch v := an.(type) {
	case Fooable:
		fmt.Println("I'm a Fooable")
	case CmdFooable:
		fmt.Println("I'm a CmdFooable")
	case ThingContainer:
		fmt.Println("I'm a ThingContainer")
	case Thing:
		fmt.Println("I'm a Thing")
	default:
		fmt.Println("Unsupported type: %T", v)
	}
}

func main() {
	PrintType(ThingContainer{})
}
