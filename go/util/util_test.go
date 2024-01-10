package util_test

import (
	"fmt"
	"log"
	"os"

	"github.com/rwxrob/lab/go/util"
)

func ExampleParseYAMLMaps() {
	buf, err := os.ReadFile(`testdata/multidoc-mapsonly.yaml`)
	if err != nil {
		log.Print(err)
	}
	maps, err := util.ParseYAMLMaps(buf)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(maps)

	// Output:
	// ---
	// foo: bar
	// title: YAML Ain't Markup Language™
	// ---
	// another: thing
	// one: 1
	// ---
	// some: random thing

}

func ExampleInlineYAMLMap_null() {
	var xxx util.InlineYAMLMap
	fmt.Println(xxx)
	// Output:
	// null
}
