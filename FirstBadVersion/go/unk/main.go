package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int, check func(int) bool) int {
	//lowest insert point for a bad vesion

	lo, hi := 0, n

	m := 0
	for lo < hi {
		m = ((hi - lo) / 2) + lo

		if check(m) {
			hi = m
		} else {
			lo = m + 1
		}
	}

	return lo
}

func isBadVersion(n, v int) bool {
	return n >= v
}

func getBadVersionFn(v int) func(n int) bool {
	return func(n int) bool { return isBadVersion(n, v) }
}

func main() {
	println(firstBadVersion(5, getBadVersionFn(4)))
	println(firstBadVersion(1, getBadVersionFn(1)))
}
