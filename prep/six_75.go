package main

import (
	"fmt"
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	maxAverage := float64(math.MinInt64)

	if len(nums) == 1 {
		return float64(nums[0])
	}

	for i := 0; i <= len(nums) - k; i++ {
		current := nums[i:i+k]
	  currentSum := 0.0
		for _, val := range current {
			currentSum += float64(val)
		}
		if currentSum / float64(k) > maxAverage {
			maxAverage = currentSum / float64(k)
		}
	}

	return float64(maxAverage)
}

func main() {
	l := []int{1,12,-5,-6,50,3}
	res := findMaxAverage(l, 4)
	fmt.Print(res)
}