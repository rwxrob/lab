package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Foo struct {
	O map[string]any `yaml:",inline" json:""`
}

func main() {
	s := new(Foo)
	buf, _ := os.ReadFile("sample.yaml")
	yaml.Unmarshal(buf, s)
	fmt.Println(s)
	buff, err := json.Marshal(s.O) // BOOM
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(buff))
}
