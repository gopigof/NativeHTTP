package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readFile(directory string, filename string) ([]byte, error) {
	filePath := filepath.Join(directory, filename)

	fmt.Println(filePath, directory, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", filename)
		return nil, err
	}
	return content, nil
}
