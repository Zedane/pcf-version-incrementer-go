package cmd

import (
	"fmt"
	"pcfv/data"

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
	Run:     RunIncrementCmd,
}

func RunIncrementCmd(cmd *cobra.Command, args []string) {
	manifest, err := data.ReadManifest(true)
	if err != nil {
		fmt.Println("Could not find manifest file.")
		if verbose {
			fmt.Println(err.Error())
		}

		return
	}

	version := data.ParseVersion(&manifest.Version)
	if version == nil {
		fmt.Println("Version is an invalid format")
		return
	}

	if major {
		version.IncrementMajor()
	} else if patch {
		version.IncrementPatch()
	} else {
		version.IncrementMinor()
	}

	fmt.Printf("%s -> %s\n", manifest.Version, version)

	err = manifest.Update(version)
	if err != nil {
		fmt.Println(err.Error())
	}

	packageFile, err := data.ReadPackage(true)
	if err != nil {
		fmt.Println("Could not read package file")
		if verbose {
			fmt.Println(err.Error())
		}
		return
	}

	err = packageFile.Update(version)
	if err != nil {
		fmt.Println(err.Error())
	}

	solutionFile, err := data.ReadSolution(true)
	if err != nil {
		fmt.Println("Could not read solution file")
		if verbose {
			fmt.Println(err.Error())
		}
		return
	}

	err = solutionFile.Update(version)
	if err != nil {
		fmt.Println(err.Error())
	}
}
