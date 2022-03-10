package main

import "fmt"

type Loc struct {
	X int
	Y int
}

func main() {

	m := map[Loc]string{}
	l := Loc{1, 2}
	m[l] = "something"
	fmt.Println(m)

}
