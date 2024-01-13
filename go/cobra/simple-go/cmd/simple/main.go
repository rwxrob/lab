package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rwxrob/lab/go/cobra/simple"
)

func usage() {
	log.SetFlags(0)
	log.Fatalf(`usage: %v greet [NAME]`, os.Args[0])
}

func main() {
	log.Print(os.Args)
	log.Print(len(os.Args))

	var greeting string

	switch len(os.Args) {
	case 0, 1:
		usage()
	case 2:
		greeting = simple.Greet(``)
	default:
		glob := strings.Join(os.Args[2:], ` `)
		greeting = simple.Greet(glob)
	}

	fmt.Println(greeting)
}
