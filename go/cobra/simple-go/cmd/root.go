package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "simple",
	Short: "Simple example of Cobra command",
}

func init() {
	rootCmd.AddCommand(pingCmd)
	rootCmd.AddCommand(sayCmd)
}

// tool add-label-to-resource -n bar
// tool add-label-to-namespace -k Deployment -n bar

// tool transform add-label-to-resource -n bar
// tool transform add-label-to-namespace -k Deployment -n bar

// tool t
// tool trans
