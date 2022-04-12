package main

import "fmt"

func main() {

	var f any

	switch v := f.(type) {
	case nil:
		fmt.Printf("%T", v)
	}
	fmt.Println("humm")

}
