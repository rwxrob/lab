package main

import "log"

func Add[V ~int](a, b V) V {
	return a + b
}

func main() {
	var a int64 = 23
	var b int64 = 35235
	log.Println(Add(a, b))
}
