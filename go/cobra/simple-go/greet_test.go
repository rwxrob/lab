package simple_test

import (
	"fmt"

	"github.com/rwxrob/lab/go/cobra/simple"
)

func ExampleGreet() {

	fmt.Println(simple.Greet(`Rob`))

	// Output:
	// Hello, Rob!

}

/*
func ExampleGreet_noargs() {

	fmt.Println(simple.Greet(``))

	// Output:
	// Hello there.

}
*/
