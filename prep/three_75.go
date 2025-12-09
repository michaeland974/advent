package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) == 0 {
		return []bool{}
	}

	max := 0

	for _, i := range candies {
		if i > max {
			max = i
		}
	}

	result := make([]bool, len(candies))

	for index, val := range candies {
		if val+extraCandies >= max {
			result[index] = true
		} else {
			result[index] = false
		}
	}

	return result
}

func main() {
	res := kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)
	print(res)
}