package cmd

import (
	"fmt"
	"local/cmds"
)

func init() {
	x := cmds.Add("hi")
	x.Method = func(args ...string) error {
		fmt.Println("Hi!")
		return nil
	}
}
