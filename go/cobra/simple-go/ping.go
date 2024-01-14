package simple

import "fmt"

func Ping() {
	switch Lang {
	case `fr`:
		fmt.Println(`pong in french`)
	case `en`:
		fmt.Println(`pong`)
	}
}
