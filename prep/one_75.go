package main

import (
	"strings"
)

// Input:
// 			word1 = "abc",
// 			word2 = "pqr"
// Output: "apbqcr"
// Explanation: The merged string will be merged as so:
// word1:  a   b   c
// word2:    p   q   r
// merged: a p b q c r

func mergeAlternately(word1 string, word2 string) string {
	longerWord := word1
	var sb strings.Builder

	if len(word2) > len(word1) {
		longerWord = word2
	}

	for i, _ := range longerWord {
		if i < len(word1) {
			word1Runes := []rune(word1)
			sb.WriteRune(word1Runes[i])
		}
		if i < len(word2) {
			word2Runes := []rune(word2)
			sb.WriteRune(word2Runes[i])
		}
	}
	return sb.String()
}

func main() {
	result := mergeAlternately("abcZZZZZZZZ","zxyEEE")
	print(result)
}