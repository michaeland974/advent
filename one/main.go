package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(path string) ([]string, error) { // Note the changed signature
	file, err := os.Open(path)

	if err != nil {
		// Return a nil slice and the original error
		return nil, err 
	}
	defer file.Close()

	words := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)
	}

	if err := scanner.Err(); err != nil {
		// Return a nil slice and the original scanning error
		return nil, fmt.Errorf("error reading lines: %w", err) 
	}
	
	// If successful, return the slice and a nil error
	return words, nil 
}

func main() {
	fileData, err := parseFile("one/input.txt")

	if err != nil {
		fmt.Printf("Fatal error processing file: %v\n", err)
		return
	}

	// 7. Print the resulting slice
	fmt.Println("Total words:", len(fileData))
}