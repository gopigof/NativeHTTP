package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readFile(directory string, filename string) ([]byte, error) {
	filePath := filepath.Join(directory, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", filename)
		return nil, err
	}
	return content, nil
}

func writeFile(directory string, filename string, content []byte) error {
	filePath := filepath.Join(directory, filename)
	err := os.WriteFile(filePath, content, 0644)
	if err != nil {
		fmt.Println("Error writing file", filePath)
	}
	return nil
}

func logRequest(req *Request) {
	fmt.Printf("[%s]: %s\n", req.method, req.uri)

	// Optional: Log headers if needed
	//fmt.Println("[HEADERS]")
	//for key, value := range req.headers {
	//	fmt.Printf("  %s: %s\n", key, value)
	//}
}
