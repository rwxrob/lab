package simple_test

import (
	"fmt"
	"strings"

	"github.com/rwxrob/lab/go/cobra/simple"
	"gopkg.in/yaml.v3"
)

func ExampleFormatYAML() {

	yamlbyt := strings.NewReader(`
%YAML 1.1
---
alist:
  - some: thing
    wicked: comes
---
configish: |
  alist:
    - some: thing
      laugh: muahaha
`)

	byt, err := simple.FormatYAML(yamlbyt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(byt))

	// Output:
	// ---
	// alist:
	//     - some: thing
	//       wicked: comes
	// ---
	// configish: |
	//     alist:
	//       - some: thing
	//         laugh: muahaha

}

func ExampleParseYAMLMaps() {

	yamlbyt := strings.NewReader(`
%YAML 1.1
---
alist:
  - some: thing
    wicked: comes
---
configish: |
  alist:
    - some: thing
      laugh: muahaha
`)

	maps, err := simple.ParseYAMLMaps(yamlbyt)
	if err != nil {
		fmt.Println(err)
	}

	for _, m := range maps {
		fmt.Println("---")
		byt, err := yaml.Marshal(m)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(string(byt))
	}

	// Output:
	// ---
	// alist:
	//     - some: thing
	//       wicked: comes
	// ---
	// configish: |
	//     alist:
	//       - some: thing
	//         laugh: muahaha

}

/*
func ExampleParseYAML() {

	yamlbyt := strings.NewReader(`

---
1: true
some: thing
another: thing
---
a scalar here
--- !!str
imma: map in reality despite the tag
`)

	it, err := simple.ParseYAML(yamlbyt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Kind: %v\n", it.Kind)
	fmt.Printf("Has: %v\n", len(it.Content))

	//for _, child := range node.Content {
	//fmt.Println(child.Kind)
	//}

	// Output:
	// something

}
*/

func ExampleToYAML() {

	byt, _ := yaml.Marshal(struct{ Key, Val string }{Key: `immakey`, Val: `immaval`})
	fmt.Print(string(byt))

	// Output:
	// key: immakey
	// val: immaval
}

func ExampleToYAML_node() {

	r := strings.NewReader(`
some: thing
another: one
1: one
`)

	d := yaml.NewDecoder(r)
	node := new(yaml.Node)
	err := d.Decode(node)
	if err != nil {
		fmt.Println(err)
	}
	byt, _ := yaml.Marshal(node)
	fmt.Print(string(byt))

	// Output:
	// some: thing
	// another: one
	// 1: one
}

func ExampleToYAML_yaml_Docs() {

	r := strings.NewReader(`
---
some: thing
#another: original
another: one
1: one

# I'm a footer comment on 1 I think

# I think I should be ignored

# I hope I am a header comment

final: thing

---
what: is this
`)

	d := yaml.NewDecoder(r)

	node1 := new(yaml.Node)
	err := d.Decode(node1)
	if err != nil {
		fmt.Println(err)
	}

	node2 := new(yaml.Node)
	err = d.Decode(node2)
	if err != nil {
		fmt.Println(err)
	}

	nodes := []*yaml.Node{node1, node2}
	//byt, err := json.MarshalIndent(nodes, "  ", "  ")
	byt, err := simple.ToYAML(nodes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(byt))

	// Output:
	// ---
	// some: thing
	// #another: original
	// another: one
	// 1: one
	// # I'm a footer comment on 1 I think
	//
	// # I think I should be ignored
	//
	// # I hope I am a header comment
	// final: thing
	// ---
	// what: is this

}
