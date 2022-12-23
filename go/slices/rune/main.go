package main

import (
	"fmt"
	"os"
)

func main() {
	buf, _ := os.ReadFile(`README.md`)
	runes := []rune(string(buf))
	for i, r := range runes {
		fmt.Printf("%v %q\n", i, r)
	}
}
