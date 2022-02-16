package foo

import (
	"cyclic-example/foo/bar"
	"log"
)

type Fooer interface {
	Foo()
}

type Thing struct{}

func (t *Thing) Foo() {
	bar.Bar(t)
	log.Println("I'm foo-ing")
}
