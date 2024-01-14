package simple_test

import (
	"fmt"

	"github.com/rwxrob/lab/go/cobra/simple"
)

func ExampleFormatYAML() {

	yamlbyt := []byte(`
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

	yamlbyt := []byte(`
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
		fmt.Print(string(simple.ToYAML(m)))
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

func ExampleToYAML() {

	fmt.Print(string(simple.ToYAML(struct{ Key, Val string }{Key: `immakey`, Val: `immaval`})))

	// Output:
	// key: immakey
	// val: immaval
}
