package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show the current version",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("VERSION")
	},
}
