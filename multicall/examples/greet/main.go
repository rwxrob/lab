package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var MainCommand = "foo"

func init() {
	//fmt.Printf("MainCommand (start of init): %v\n", MainCommand)
	MainCommand = filepath.Base(os.Args[0])
	//fmt.Printf("MainCommand (end of init): %v\n", MainCommand)
}

func main() {
	//fmt.Printf("MainCommand (start of main()): %v\n", MainCommand)
	switch MainCommand {
	case "greet":
		fmt.Println("Greetings!")
	case "hi":
		fmt.Println("Hi there!")
	case "hello":
		fmt.Println("Hello.")
	case "salut":
		fmt.Println("Salut!")
	case "privet":
		fmt.Println("Привет!")
	}
}
