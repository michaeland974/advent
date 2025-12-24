package main

func compareKeys(map1, map2 map[rune]int) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key := range map1 {
		if _, exists := map2[key]; !exists {
			return false
		}
	}
	return true
}

func compareValues(map1, map2 map[rune]int) bool {
	for key, val1 := range map1 {
		if val2 := map2[key]; val1 != val2 {
			return false
		}
	}
	return true
}

func closeStrings(word1 string, word2 string) bool {
	map1 := make(map[rune]int)
	map2 := make(map[rune]int)

	for _, r := range word1 {
		map1[r]++
	}
	for _, r := range word2 {
		map2[r]++
	}

	keysMatch := compareKeys(map1, map2)
	valuesMatch := compareValues(map1, map2)

	if keysMatch && valuesMatch {
		return true
	}
	return false
}