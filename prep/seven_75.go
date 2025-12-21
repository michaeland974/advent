package main

import (
	"fmt"
	"strings"
)

func countVowels(s string) int {
	count := 0
	s = strings.ToLower(s)
	vowels := "aeiou"

	for _, c := range s {
		if strings.ContainsRune(vowels, c) {
			count++
		}
	}
	return count
}

func maxVowels(s string, k int) int {
	maxNum := 0

	for i := 0; i <= len(s) - k; i++ {
		window := s[i:k+i]
		n := countVowels(window)
		if n > maxNum {
			maxNum = n
		}
	}

	return maxNum
}

func main() {
	s := "leetcode"
	k := 3
	res := maxVowels(s, k)
	fmt.Println(res)
}