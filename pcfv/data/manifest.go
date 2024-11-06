package data

import (
	"encoding/xml"
	"fmt"
)

var MANIFEST_FILE string = "ControlManifest.Input.xml"

type Manifest struct {
	XMLName xml.Name `xml:"manifest"`
	Control Control  `xml:"control"`
}

type Control struct {
	XMLName xml.Name `xml:"control"`
	Version string   `xml:"version,attr"`
	Name    string   `xml:"display-name-key,attr"`
}

func ReadManifest() (*Manifest, error) {
	fileContent, err := ReadFile(MANIFEST_FILE)
	if err != nil {
		return nil, err
	}

	var manifest Manifest

	err = xml.Unmarshal(*fileContent, &manifest)
	if err != nil {
		return nil, err
	}

	return &manifest, nil
}

func (m *Manifest) Print(verbose bool, includeName bool) {
	if !verbose {
		if includeName {
			fmt.Printf("Manifest: %s", m.Control.Version)
		} else {
			fmt.Println(m.Control.Version)
		}

		return
	}

	fmt.Printf("Name: %s\n", m.Control.Name)
	fmt.Printf("File: %s\n", MANIFEST_FILE)
	fmt.Printf("Version: %s\n", m.Control.Version)
}
