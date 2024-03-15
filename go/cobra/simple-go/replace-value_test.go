package simple_test

/*
func ExampleReplaceValue() {

	yamlbyt := strings.NewReader(`
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
unrelated: |
  I am a random string paragraph that just happens to have
  some: thing on a line by itself
  and more stuff after that.
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

	for _, m := range maps {
		fmt.Println(`---`)
		fmt.Print(string(simple.ToYAML(m)))
	}

	somex := regexp.MustCompile(`^some$`)
	valx := regexp.MustCompile(`^thing$`)

	err = simple.ReplaceValue(somex, valx, `new thing`, maps)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(simple.ToYAMLDocs(maps)))

	// Output:
	// ---
	// some: new thing
	// ---
	// another:
	//     some: new thing
	// other: thing
	// ---
	// alist:
	//     - some: new thing
	//       wicked: comes
	// ---
	// configish: |
	//     some: new thing
	//     imma: string
	//     alist:
	//       - some: new thing
	//         laugh: muahaha
	// unrelated: |
	//     I am a random string paragraph that just happens to have
	//     some: thing on a line by itself
	//     and more stuff after that.
	// ---
	// some: other thing
	// ---
	// another:
	//     some: other thing
	// other: other thing
	// ---
	// alist:
	//     - some: other thing
	//       wicked: comes
	// ---
	// configish: |
	//     some: other thing
	//     imma: string
	//     alist:
	//       - some: other thing
	//         laugh: muahaha

}
*/
