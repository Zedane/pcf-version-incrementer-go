package data

import (
	"encoding/json"
	"fmt"
)

var PACKAGE_FILE string = "package.json"

type Package struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func ReadPackage() (*Package, error) {
	fileContent, err := ReadFile(PACKAGE_FILE)
	if err != nil {
		return nil, err
	}

	var p Package

	err = json.Unmarshal(*fileContent, &p)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (p *Package) Print(verbose bool) {
	if !verbose {
		fmt.Printf("Package: %s", p.Version)
		return
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("File: %s\n", PACKAGE_FILE)
	fmt.Printf("Version: %s\n", p.Version)
}
