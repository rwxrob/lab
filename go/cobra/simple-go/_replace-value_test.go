package simple_test

import "fmt"

func ExampleReplaceValue() {

	yamlbyt := []byte(`
---
some: thing
---
other: thing
another:
  some: thing
---
alist:
  - some: thing
    wicked: comes
---
configish: |
  some: thing
  imma: string
  alist:
    - some: thing
      laugh: muahaha
---
some: other thing
---
other: other thing
another:
  some: other thing
---
alist:
  - some: other thing
    wicked: comes
---
configish: |
  some: other thing
  imma: string
  alist:
    - some: other thing
      laugh: muahaha
`)

	maps, err := simple.ParseYAMLMaps(yamlbyt)
	if err != nil {
		fmt.Println(err)
	}

	byt, err := simple.ReplaceValue(somex, valx, newval, maps)

	// Output:
	// blah
}
