package main

import "fmt"

const VersionConst = ``

var Version string

//var VersionInt int // nope

func main() {
	fmt.Println(Version)
	fmt.Println(VersionConst)
	// 	fmt.Println(VersionInt)  // nope
}
