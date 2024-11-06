package data

import (
	"encoding/xml"
	"fmt"
)

var SOLUTION_FILE string = "Solution.xml"

type Solution struct {
	XMLName  xml.Name         `xml:"ImportExportXml"`
	Manifest SolutionManifest `xml:"SolutionManifest"`
}

type SolutionManifest struct {
	XMLName xml.Name `xml:"SolutionManifest"`
	Name    string   `xml:"UniqueName"`
	Version string   `xml:"Version"`
}

func ReadSolution() (*Solution, error) {
	fileContent, err := ReadFile(SOLUTION_FILE)
	if err != nil {
		return nil, err
	}

	var solution Solution

	err = xml.Unmarshal(*fileContent, &solution)
	if err != nil {
		return nil, err
	}

	return &solution, nil
}

func (s *Solution) Print(verbose bool) {
	if !verbose {
		fmt.Printf("Solution: %s", s.Manifest.Version)
		return
	}

	fmt.Printf("Name: %s\n", s.Manifest.Name)
	fmt.Printf("File: %s\n", SOLUTION_FILE)
	fmt.Printf("Version: %s\n", s.Manifest.Version)
}
