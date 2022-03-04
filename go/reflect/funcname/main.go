package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

func Some() {}

func FuncName(i any) string {
	n := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return n[len(n)-1]
}

func main() {
	fmt.Println(FuncName(Some))
}
