package cmd

import (
	"fmt"
	"pcfv/data"

	"github.com/spf13/cobra"
)

var showAll bool

func init() {
	statusCmd.Flags().BoolVarP(&showAll, "all", "a", false, "Show version of all files")
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Short:   "Show the current version",
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		manifest, err := data.ReadManifest()
		if err != nil {
			fmt.Println("Could not find manifest file.")
			if verbose {
				fmt.Println(err.Error())
			}

			return
		}

		manifest.Print(verbose, showAll)

		if !showAll {
			return
		}

		packageFile, err := data.ReadPackage()
		if err != nil {
			fmt.Println("Could not read package file")
			if verbose {
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Print("\n")
		packageFile.Print(verbose)

		solutionFile, err := data.ReadSolution()
		if err != nil {
			fmt.Println("Could not read solution file")
			if verbose {
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Print("\n")
		solutionFile.Print(verbose)
	},
}
