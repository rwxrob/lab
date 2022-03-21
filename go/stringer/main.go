package main

import "fmt"

type Blah struct{}

// force compile-time check for fulfillment of interface.
var _ fmt.Stringer = new(Blah)
var _ fmt.GoStringer = new(Blah)

func main() {
}
