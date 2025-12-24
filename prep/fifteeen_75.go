package main

import (
	"fmt"
)

func removeStars(s string) string {
	var stack []rune

	for _, r := range s {
		if r == '*' && len(stack) > 0{
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
		}
	return string(stack)
}

func main() {
	res := removeStars("leet*code")
	fmt.Println(res)
}