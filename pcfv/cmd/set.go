package cmd

import (
	"fmt"
	"pcfv/data"

	"github.com/spf13/cobra"
)

func init() {
	setCmd.Flags().Int16Var(&majorV, "major", -1, "Set the major version")
	setCmd.Flags().Int16Var(&minorV, "minor", -1, "Set the minor version")
	setCmd.Flags().Int16Var(&patchV, "patch", -1, "Set the patch version")

	rootCmd.AddCommand(setCmd)
}

var majorV int16
var minorV int16
var patchV int16

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a specific vesion",
	Run:   RunSetCmd,
}

func RunSetCmd(cmd *cobra.Command, args []string) {
	if majorV < 0 && minorV < 0 && patchV < 0 {
		fmt.Println("Please input valid version")
		return
	}

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

	if majorV > -1 {
		version.Major = int(majorV)
	}

	if minorV > -1 {
		version.Minor = int(minorV)
	}

	if patchV > -1 {
		version.Patch = int(patchV)
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
