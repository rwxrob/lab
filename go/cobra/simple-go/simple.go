package simple

import (
	"fmt"
)

func Greet(name string) string {
	if len(name) == 0 {
		return `Hello there`
	}
	return fmt.Sprintf(`Hello, %v!`, name)
}
