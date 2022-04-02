package main

import "fmt"

// THIS IS BORK
var EscThese = '\t' & '\n' & '|' & '&' & ';' & '(' & ')' & '<' & '>' & '!' & '[' & ']'

func main() {
	var buf []rune
	for _, r := range "a s\ntring" {
		if EscThese == EscThese|r {
			buf = append(buf, '\\', r)
		}
		buf = append(buf, r)
	}
	fmt.Println(string(buf))
}
