package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"time"
)

//go:embed files
var fsys embed.FS

func main() {

	info1, err := fs.Stat(fsys, "files/1gbfile")
	//	info1, err := fs.Stat(fsys, "files/busybox")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(info1)
	fmt.Printf("Name: %v\n", info1.Name())
	fmt.Printf("Size: %v\n", info1.Size())
	fmt.Printf("Mode: %v\n", info1.Mode())
	fmt.Printf("ModTime: %v\n", info1.ModTime())
	fmt.Printf("IsDir: %v\n", info1.IsDir())
	fmt.Printf("Sys: %v\n", info1.Sys())

	/*

	   	info2, err := fs.Stat(fsys, "files/1gbfileb")
	   	if err != nil {
	   		log.Print(err)
	   	//fmt.Println(info2)
	   	fmt.Println(info2.Name())
	   	fmt.Println(info2.Size())
	   }
	*/

	for {
		time.Sleep(1 * time.Second)
	}
}
