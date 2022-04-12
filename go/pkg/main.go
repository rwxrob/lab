package main

// ----------------------------- the fooer ----------------------------

type fooer struct{}

func (fooer) Foo() {}

func NewFooer() *fooer { return new(fooer) }

type Fooer interface {
	Foo()
}

// ------------------------------- thing ------------------------------

type thing struct {
	Pub   string
	Fooer Fooer
}

// ------------------------- one of the thing -------------------------

var some = &thing{
	Fooer: NewFooer(),
}

// ------------------------------- main  ------------------------------

func main() {
	some.Fooer.Foo()
}
