package simple

import "fmt"

func Greet(name string) string {
	greeting := `Hello there.`
	switch Lang {
	case `en`:
		greeting = fmt.Sprintf(`Hello, %v!`, name)
	case `fr`:
		greeting = fmt.Sprintf(`Salut, %v!`, name)
	}
	return greeting
}
