package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(path string) (string, error) {
	contentBytes, err := os.ReadFile(path)

	if err != nil {
		fmt.Printf("Error reading file %v\n", err)
	}

	return string(contentBytes), err
}

type EvalFunc func(int) bool

func applyToRange(start, end int, operation EvalFunc) int {
	count := 0
	for i := start; i <= end; i++ {
		if operation(i) {
			count++
		}
	}	
	return count
}

func isTwice(num int) bool {
	s := strconv.Itoa(num)
	l := len(s)
	if l % 2 == 0 {
		middleIndex := l / 2
		firstPart := s[:middleIndex]
		secondPart := s [middleIndex:]
		if firstPart == secondPart {
			return true
		}
	}
	return false
}

func main() {
	c, err := parseFile("two/input.txt")

	if err != nil {
		fmt.Printf("Error fromm parseFile() %v\n", err)
	}

	// count := 0
	limits := strings.Split(c, ",")
	finalCount := 0

	for _, limit := range limits {
		fmt.Printf("limit is %s\n", limit)
	
		parts := strings.Split(limit, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])
		result := applyToRange(start, end, isTwice)
		finalCount += result
	}

	// testing := "11-22"

	// parts := strings.Split(testing, "-")
	// start, _ := strconv.Atoi(parts[0])
	// end, _ := strconv.Atoi(parts[1])

	// result := applyToRange(start, end, isTwice)
	fmt.Printf("Result is %d", finalCount)
}