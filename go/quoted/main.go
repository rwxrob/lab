package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Printf("%q%q\n", "=", ">") // "="">"
	buf, _ := json.Marshal("=>")   // "=\u003e"
	fmt.Println(string(buf))       // shitty escaping of > for no fucking good reason
}
