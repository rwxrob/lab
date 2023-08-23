package main

import "fmt"

type Person struct {
	Name string
	Age  int

	printer *Printer
}

func (p Person) Printer() *Printer { return p.printer }

type Printer struct {
	Upper bool
}

var UpperPrinter = &Printer{true}

var me = Person{
	printer: UpperPrinter,
}

func main() {
	fmt.Printf("%p should be same as %p", me.printer, me.Printer())
}
