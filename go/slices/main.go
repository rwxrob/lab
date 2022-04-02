package main

import "fmt"

type blah struct {
	names  []string
	tags   map[string]string
	pnames *blah
}

func testArgs(args []string) {
	if len(args) == 0 {
		fmt.Println("args are zero")
	}
	if args == nil {
		fmt.Println("args is nil")
	}
}

func testMapArgs(args map[string]string) {
	if len(args) == 0 {
		fmt.Println("args are zero")
	}
	if args == nil {
		fmt.Println("args is nil")
	}
}

func main() {
	testArgs(nil)
	testArgs([]string{})
	fmt.Println("=======")
	testMapArgs(nil)
	testMapArgs(map[string]string{})
	b := blah{}
	fmt.Println(b.names)
	fmt.Println(b.names == nil)
	fmt.Println(b.tags)
	fmt.Println(b.tags == nil)
	fmt.Println(b.pnames)
	fmt.Println(b.pnames == nil)
	fmt.Println(len(b.names))
	b.names = append(b.names, "rob")
	fmt.Println(b.names)
}
