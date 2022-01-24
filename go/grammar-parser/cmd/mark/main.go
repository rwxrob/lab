package main

import (
	"os"
)

func main() {
	defer lib.trapPanic()
	if len(os.Args) == 1 {
		lib.usage()
		return
	}
	mark(os.Args[1])
}
