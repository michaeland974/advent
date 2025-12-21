package main

import (
	"maps"
	"slices"
)

func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for _, n := range nums1 {
		set1[n] = struct{}{}
	}
	for _, n := range nums2 {
		set2[n] = struct{}{}
	}

	for item := range set1 {
		if _, exists := set2[item]; exists {
			delete(set1, item)
		}
	}

	resultSet1

	for item := range set2 {
		if _, exists := set1[item]; exists {
			delete(set2, item)
		}
	}

	keys1 := slices.Collect(maps.Keys(set1))
	keys2 := slices.Collect(maps.Keys(set2))

	return [][]int{keys1, keys2}
}

func main() {

}