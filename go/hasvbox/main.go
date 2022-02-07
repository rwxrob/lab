package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	fname, err := exec.LookPath("VirtualBox.exe")
	if err != nil {
		log.Println("nope")
		os.Exit(1)
	}
	log.Println("yep: " + fname)
}
