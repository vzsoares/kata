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
	start, end := 1, n

	for start <= end {
		m := ((end - start) / 2) + start

		g := guess(m)
		if g == 0 {
			return m
		} else if g == 1 {
			start = m + 1
		} else if g == -1 {
			end = m - 1
		}
	}
	return -1
}

// var pick = 6
var pick = 823161944

func guess(num int) int {
	if num == pick {
		return 0
	}
	if num < pick {
		return 1
	}
	if num > pick {
		return -1
	}
	panic("must return")
}

func main() {
	println(guessNumber(10))
	println(guessNumber(921239930))
}
