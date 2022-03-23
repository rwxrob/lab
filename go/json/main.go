package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// this works
type Foo struct {
	One int            `yaml:"two"`
	Two int            `yaml:"one"`
	O   map[string]any `yaml:",inline"`
}

func main() {
	/*
			buf := `
		ну <что> же
		😈💢©💖💘💥
		`
			byt, err := json.Marshal(buf)
			if err != nil {
				log.Print(err)
			}
			fmt.Println(string(byt))
	*/

	foo := new(Foo)
	buf, _ := os.ReadFile("foo.json")
	if err := yaml.Unmarshal(buf, foo); err != nil {
		log.Print(err)
	}
	fmt.Println(foo)

}
