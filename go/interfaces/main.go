package main

import "fmt"

type Printer interface {
	Print()
}

// ------------------------- original printer -------------------------

type printer struct {
	Buf string
}

func NewOriginal() *printer { return new(printer) }

func (p *printer) Print() {
	fmt.Println(p.Buf)
}

// -------------------------- another printer -------------------------

type another struct {
	Buffer string
}

func NewAnother() *another { return new(another) }

func (p *another) Print() {
	fmt.Println(p.Buffer)
}

// ------------ interface that depends on another interface -----------

type UsesPrinter interface {
	Use(Printer)
}

// ------------------- something that uses a Printer ------------------

type User struct {
}

func Use(p Printer) {
	p.Print()
}

// ------------------------------- main -------------------------------

func main() {
	po := NewOriginal()
	pa := NewAnother()
	po.Buf = "Original Buffer"
	pa.Buffer = "Another Buffer"
	Use(po)
	Use(pa)

}
