package data

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

var MANIFEST_FILE string = "ControlManifest.Input.xml"

type Manifest struct {
	FilePath *string
	Version  string
	Row      int
	Lines    []string
}

func ReadManifest(cache bool) (*Manifest, error) {
	fileContent, path, err := ReadFile(MANIFEST_FILE)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(*fileContent)))

	controlOpenAt := -1
	controlCloseAt := -1
	lineNum := 0

	manifest := &Manifest{
		FilePath: path,
		Version:  "",
		Row:      -1,
	}

	if cache {
		manifest.Lines = strings.Split(string(*fileContent), "\n")
	}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(strings.ToLower(line), "<control") {
			controlOpenAt = lineNum
		}

		if controlOpenAt != -1 && strings.Contains(line, "version=") {
			manifest.Row = lineNum
			startIdx := strings.Index(line, "version=")
			subStr := line[startIdx+9:]

			for _, c := range subStr {
				if c == '"' && manifest.Version != "" {
					break
				}

				manifest.Version += string(c)
			}
		}

		if controlOpenAt != -1 && strings.Contains(line, ">") {
			strOpen := false
			for _, c := range line {
				if c == '>' && !strOpen {
					controlCloseAt = lineNum
					break
				}

				if c == '"' {
					strOpen = !strOpen
				}
			}

			if controlCloseAt != -1 {
				break
			}
		}

		lineNum++
	}

	if manifest.Row == -1 {
		return nil, errors.New("unable to find version in manifest")
	}

	return manifest, nil
}

func (m *Manifest) Print(verbose bool, includeName bool) {
	if !verbose {
		if includeName {
			fmt.Printf("Manifest: %s", m.Version)
		} else {
			fmt.Println(m.Version)
		}

		return
	}

	fmt.Printf("File: %s\n", MANIFEST_FILE)
	fmt.Printf("Version: %s\n", m.Version)
	fmt.Printf("Path: %s\n", *m.FilePath)
}

func (m *Manifest) Update(v *Version) error {
	if m.Version == "" || m.Lines == nil {
		return errors.New("unable to set version in solution")
	}

	m.Lines[m.Row] = strings.Replace(m.Lines[m.Row], m.Version, v.String(), 1)

	newContents := []byte(strings.Join(m.Lines, "\n"))

	return WriteFile(m.FilePath, &newContents)
}
