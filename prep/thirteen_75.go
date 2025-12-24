package main

import (
	"fmt"
)

func equalPairs(grid [][]int) int {
	// slices are equal length
	n := len(grid[0])
	count := 0

	frqMap := make(map[string]int, n*2)

	for _, row := range grid {
		key := fmt.Sprint(row)
		frqMap[key]++
	}

	for i := 0; i < n; i++ {

		currentCol := make([]int, n)

		for j := 0; j < n; j++ {
			currentCol[j] = grid[j][i]
		}

		colKey := fmt.Sprint(currentCol)
		if val, exists := frqMap[colKey]; exists {
			count += val
		}
	}
	return count
}

func main() {
	grid := [][]int{
			{3, 2, 1}, // Row 0
			{1, 7, 6}, // Row 1
			{2, 7, 7}, // Row 2
	}
	res := equalPairs(grid)
	fmt.Print(res)
}