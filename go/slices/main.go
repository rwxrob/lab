package main

import "fmt"

type blah struct {
	names  []string
	tags   map[string]string
	pnames *blah
}

func main() {
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
