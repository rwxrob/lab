package cmds

import (
	"fmt"
	"os"
)

// registry of commands (happens before init())
var _x = map[string]*Command{}

func Has(c string) bool {
	_, has := _x[c]
	return has
}

func Do(c string) error {
	if cmd, ok := _x[c]; ok {
		return cmd.Method(os.Args[1:]...)
	}
	return fmt.Errorf("command not found: %v", c)
}

func Rename(from, to string) {
	_x[to] = _x[from]
	delete(_x, from)
}
