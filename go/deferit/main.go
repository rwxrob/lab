package main

import (
	"errors"
	"fmt"
)

func main() {
	err := readFile("")
	fmt.Print(err)
}

func readFile(name string) (err error) {

	defer func() {
		closeerr := errors.New(`happened within defer`)
		//	olderr := err
		if closeerr != nil {
			err = errors.Join(err, closeerr)
		}
		//fmt.Printf(`close %p old %p final %p`, closeerr, olderr, err)
		fmt.Printf("close %p final %p\n", closeerr, err)
	}()

	// do something unhappy
	err = errors.New(`regular`)
	return
}
