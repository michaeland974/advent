package main

import (
	"maps"
	"slices"
)

func isUniqueSlice(nums []int) bool {
	set := make(map[int]struct{})

	for _, n := range nums {
		set[n] = struct{}{}
	}

	return len(nums) == len(set)
}

func uniqueOccurrences(arr []int) bool {
	frqMap := make(map[int]int)

	for _, v := range arr {
		frqMap[v]++
	}

	frqMapKeys := slices.Collect(maps.Values(frqMap))
	return isUniqueSlice(frqMapKeys)
}