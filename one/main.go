package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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

type Rotation struct {
	Direction string
	Value     int
}

func parseRotation(s string) Rotation {
	c := string(s[0])
	numStr := s[1:]

	val, _ := strconv.Atoi(numStr)

	if c == "L" {
		val = -val
	}

	return Rotation {
		Direction: c,
		Value: val,
	}
}

func aggregateCyclic(current int, change int) (int, int) {
	rawSum := current + change	
	modResult := rawSum % 100
	result := (modResult + 100) % 100
	cycleAmount := rawSum - result
	cycles := cycleAmount / 100

	return result, int(math.Abs(float64(cycles)))
}

func main() {
	fileData, err := parseFile("one/input.txt")

	if err != nil {
		fmt.Printf("Fatal error processing file: %v\n", err)
		return
	}

	rotations := []Rotation{}
	
	// Parse rotations
	for _, word := range fileData {
		rotations = append(rotations, parseRotation(word))
	}
	
  currentPosition := 50
	diff := 0
	count := 0

	for _, rot := range rotations {
		currentPosition, diff =  aggregateCyclic(currentPosition, rot.Value)
		if currentPosition == 0 {
			count++
		}
		count += diff
	}

	fmt.Print(count)
}