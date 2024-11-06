package data

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func GetFilePath(name string) (*string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var filePath string

	err = filepath.Walk(wd, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() && info.Name() == "node_modules" {
			return filepath.SkipDir
		}

		if err == nil && info.Name() == name {
			filePath = path
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &filePath, nil
}

func ReadFile(name string) (*[]byte, error) {
	filePath, err := GetFilePath(name)
	if err != nil {
		return nil, err
	}

	file, err := os.Open(*filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return &fileContent, nil
}