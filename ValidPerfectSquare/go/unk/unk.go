package main

// https://leetcode.com/problems/valid-perfect-square/

func isPerfectSquare(num int) bool {
	// num has perfect square
	l := num
	lo, hi := 0, l

	if num == 1 {
		return true
	}

	for lo < hi {
		m := ((hi - lo) / 2) + lo

		r := m * m
		if r == num {
			return true
		} else if r > num {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return false
}

func main() {
	println(isPerfectSquare(16), "true")
	println(isPerfectSquare(14), "false")
}
