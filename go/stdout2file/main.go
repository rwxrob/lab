package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	orig := os.Stdout
	defer func() { os.Stdout = orig }()
	file, _ := os.Create(`out`)
	defer os.Remove(`out`)
	os.Stdout = file
	fmt.Println("something")
	buf, _ := os.ReadFile(`out`)
	log.Print(string(buf))
}
