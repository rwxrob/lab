package main

import (
	_ "local/cmd-hello"
	_ "local/cmd-hi"
	_ "local/cmd-hiagain"
	"local/cmds"
	"os"
)

func init() { cmds.Rename("hi_", "hiagain") }

func main() {
	if len(os.Args) > 1 {
		cmds.Do(os.Args[1])
		return
	}
	cmds.Do("hi")
}
