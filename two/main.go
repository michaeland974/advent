package main

import (
	"fmt"
	"os"
	"strconv"
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
		return count
	}	

	return 0
}

func isTwice(num int) bool {
	s := strconv.Itoa(num)
	l := len(s)
	if l % 2 == 0 {
		middleIndex := l / 2
		firstPart := s[:middleIndex]
		secondPart := s [middleIndex:]
		fmt.Println(firstPart)
		fmt.Println(secondPart)
		if firstPart == secondPart {
			return true
		}
	}
	return false
}

func main() {
	// c, err := parseFile("two/input.txt")

	// if err != nil {
	// 	fmt.Printf("Error fromm parseFile() %v\n", err)
	// }

	//nums := strings.Split(c, ",")

	// isTwice(88128812)
	// isTwice(112112)
	isTwice(54354)
}