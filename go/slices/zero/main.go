package main

import "fmt"

func testSlice(args []string) {
	if len(args) == 0 {
		fmt.Println("args length is zero")
	}
	if args == nil {
		fmt.Println("args value is nil")
	}
}

func testMap(args map[string]string) {
	if len(args) == 0 {
		fmt.Println("args length is zero")
	}
	if args == nil {
		fmt.Println("args value is nil")
	}
}

func main() {
	fmt.Println("======= testSlice(nil)")
	testSlice(nil)
	fmt.Println("======= testSlice([]string{})")
	testSlice([]string{})
	fmt.Println("======= testMap(nil)")
	testMap(nil)
	fmt.Println("======= testMap(map[string]string{})")
	testMap(map[string]string{})
	fmt.Println(`======= testMap(map[string]string{"foo":"bar"})`)
	testMap(map[string]string{"foo": "bar"})
}
