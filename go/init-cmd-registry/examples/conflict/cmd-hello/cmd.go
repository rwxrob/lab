package cmd

import (
	"fmt"
	"local/cmds"
)

func init() {
	x := cmds.Add("hello")
	x.Method = func(args ...string) error {
		fmt.Println("Hello!")
		return nil
	}
}
