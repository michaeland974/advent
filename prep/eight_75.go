package main

func largestAltitude(gain []int) int {
	highest := 0
	current := 0

	for _, n := range gain {
		current += n
		if current > highest {
			highest = current
		}
	}

	return highest
}