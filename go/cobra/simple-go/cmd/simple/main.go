package main

import (
	"log"

	"github.com/rwxrob/lab/go/cobra/simple/cmd"
)

func init() {
	log.SetFlags(0)
}

func main() {
	cmd.Execute()
}
