package lib_test

import (
	"fmt"

	lib "github.com/rwxrob/lab-grammar-parser-in-go"
)

func ExampleState_String() {
	fmt.Println(new(lib.State))
	// Output:
	// pos: 0
	// strongem: false
	// strong: false
	// emph: false
	// code: false
	// dquote: false
	// quote: false
	// link: false
}

func ExampleMark() {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println(p)
		}
	}()
	lib.Mark(`*`)
	// Output:
	// *
}
