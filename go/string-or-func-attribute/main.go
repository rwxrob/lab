package main

import "fmt"

// ------------------------------- blah -------------------------------

type blah struct{}

func (b blah) String() string { return "something" }

// -------------------------------- Cmd -------------------------------
/*
type Stringer interface {
	String() string
}

type Cmd struct {
	Name Stringer
}

var Some = &Cmd{
	Name: blah{},
}
*/

// ------------------------------- Cmd2 -------------------------------
/*
type Stringer2 interface {
	string | []byte | func() string
}

type Cmd2 struct {
	Name Stringer2
}

var Some2 = &Cmd2{
	Name: func() string { return "something" },
}
*/

// ------------------------------ cmd3 -------------------------------

// this only works with my hack to fmt.Print*
/*

type Cmd3[T StringFunc | string] struct {
	Name T
}

var Some3 = &Cmd3[StringFunc]{
	Name: func() string { return "something" },
}
*/

type StringFunc func() string

type Cmd3 struct {
	Name interface{}
}

func Something() string { return "something" }

var Some3 = &Cmd3{
	//Name: "another thing", //func() string { return "something" },
	//Name: func() string { return "something" },
	Name: Something,
}

// ------------------------------- main -------------------------------

func main() {
	//fmt.Println(Some)
	//fmt.Println(Some2)
	fmt.Println(Some3.Name)
}
