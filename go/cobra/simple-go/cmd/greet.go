package cmd

import (
	"fmt"
	"strings"

	"github.com/rwxrob/lab/go/cobra/simple"
	"github.com/spf13/cobra"
)

var greetCmd = &cobra.Command{
	Use:   `greet`,
	Short: `Print a nice greeting`,
	Run: func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(simple.Greet(``))
			return
		}
		glob := strings.Join(args, ` `)
		greeting := simple.Greet(glob)
		fmt.Println(greeting)
	},
}
