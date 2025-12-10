package main

import (
	"fmt"
	"strings"
)


func reverseWords(s string) string {
	list := strings.Fields(strings.TrimSpace(s))
	result := []string{}

	for i := len(list) - 1; i >=0; i-- {
		el := list[i]
		result = append(result, el)
	}
	
	return strings.Join(result, " ")
}

func main() {
	s := "the sky is blue"
	result := reverseWords(s)
	fmt.Print(result)
}