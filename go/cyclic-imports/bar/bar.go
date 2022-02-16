package bar

type Fooer interface {
	Foo()
}

func Bar(f Fooer) {
	f.Foo()
}
