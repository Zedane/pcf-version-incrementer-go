package data

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

var SOLUTION_FILE string = "Solution.xml"

type Solution struct {
	FilePath *string
	// Node     *Node
	Version string
	Row     int
	Lines   []string
}

func ReadSolution(cache bool) (*Solution, error) {
	fileContent, path, err := ReadFile(SOLUTION_FILE)
	if err != nil {
		return nil, err
	}

	solution := &Solution{
		FilePath: path,
		Version:  "",
		Row:      -1,
	}

	if cache {
		solution.Lines = strings.Split(string(*fileContent), "\n")
	}

	scanner := bufio.NewScanner(strings.NewReader(string(*fileContent)))

	lineNum := 0

	for scanner.Scan() {
		line := scanner.Text()
		startIdx := strings.Index(line, "<Version>")
		if startIdx != -1 {
			endIdx := strings.Index(line, "</Version>")
			if endIdx != -1 {
				solution.Row = lineNum
				solution.Version = line[startIdx+9 : endIdx]
				break
			}
		}

		lineNum++
	}

	if solution.Row == -1 {
		return nil, errors.New("unable to read solution file")
	}

	return solution, nil
}

func (s *Solution) Print(verbose bool) {
	if !verbose {
		fmt.Printf("Solution: %s", s.Version)
		return
	}

	fmt.Printf("File: %s\n", SOLUTION_FILE)
	fmt.Printf("Version: %s\n", s.Version)
	fmt.Printf("Path: %s\n", *s.FilePath)
}

func (s *Solution) Update(v *Version) error {
	if s.Version == "" || s.Lines == nil {
		return errors.New("unable to set version in solution")
	}

	s.Lines[s.Row] = strings.Replace(s.Lines[s.Row], s.Version, v.ShortString(), 1)

	newContents := []byte(strings.Join(s.Lines, "\n"))

	return WriteFile(s.FilePath, &newContents)
}
