package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	incrementCmd.Flags().BoolVarP(&major, "major", "m", false, "Increment the major version by one")
	incrementCmd.Flags().BoolVarP(&patch, "patch", "p", false, "Increment the patch version by one")
	rootCmd.AddCommand(incrementCmd)
}

var major bool
var patch bool

var incrementCmd = &cobra.Command{
	Use:     "increment",
	Short:   "Increment the version by one",
	Aliases: []string{"i"},
	Run: func(cmd *cobra.Command, args []string) {
		if major {
			fmt.Println("Incrementing major version")
		} else if patch {
			fmt.Println("Incrementing patch version")
		} else {
			fmt.Println("Incrementing version")
		}
	},
}
