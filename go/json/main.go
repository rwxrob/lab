package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	buf := `
ну <что> же
😈💢©💖💘💥
`
	byt, err := json.Marshal(buf)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(byt))

}
