package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    low := 1
		high := n

		for low <= high {
			mid := (high + low) / 2
			res := guess(mid)

			if res == 0 { 
				return mid
			} else if res == -1 {
				high = mid - 1
			} else if res == 1 {
				low = mid + 1
			}
		}
		return 0
}