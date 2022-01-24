package cmds

import (
	"encoding/json"
	"fmt"
	"log"
)

type Command struct {
	Name    string
	Summary string
	Method  func(args ...string) error `json:"-"`
}

// fulfills fmt.Stringer
func (c Command) String() string {
	byt, err := json.MarshalIndent(c, "  ", "  ")
	if err != nil {
		return fmt.Sprintf(`{"ERROR": %q}`, err)
	}
	return string(byt)
}

func Add(a interface{}) *Command {
	var x *Command
	switch s := a.(type) {
	case string:
		x = new(Command)
		x.Name = s
	case *Command:
		x = s
	default:
		log.Printf("unsupported type: %T", a)
	}
	if Has(x.Name) {
		x.Name += "_"
	}
	_x[x.Name] = x
	return x
}
