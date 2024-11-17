package data

import (
	"encoding/json"
	"errors"
	"fmt"
)

var PACKAGE_FILE string = "package.json"

type Package struct {
	FilePath *string
	Version  string
	Row      int
	Cache    map[string]interface{}
}

func ReadPackage(cache bool) (*Package, error) {
	fileContent, path, err := ReadFile(PACKAGE_FILE)
	if err != nil {
		return nil, err
	}

	p := &Package{
		FilePath: path,
	}

	var rawContent interface{}
	if err := json.Unmarshal(*fileContent, &rawContent); err != nil {
		return nil, err
	}

	jsonVerison, ok := rawContent.(map[string]interface{})
	if !ok {
		return nil, errors.New("wrong package format")
	}

	if cache {
		p.Cache = rawContent.(map[string]interface{})
	}

	p.Version = jsonVerison["version"].(string)

	return p, nil
}

func (p *Package) Print(verbose bool) {
	if !verbose {
		fmt.Printf("Package: %s", p.Version)
		return
	}

	fmt.Printf("File: %s\n", PACKAGE_FILE)
	fmt.Printf("Version: %s\n", p.Version)
	fmt.Printf("Path: %s\n", *p.FilePath)
}

func (p *Package) Update(v *Version) error {

	p.Cache["version"] = v.String()

	bytes, err := json.MarshalIndent(p.Cache, "", "    ")
	if err != nil {
		return err
	}

	return WriteFile(p.FilePath, &bytes)
}
