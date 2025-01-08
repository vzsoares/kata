package main

func mySqrt(x int) int {
	var start = 0
	var end = x + 1

	var mid = 0

	for start < end {
		mid = ((end - start) / 2) + start
		if mid*mid > x {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return end - 1
}

func main() {
	println(mySqrt(4))
	println(mySqrt(9))
	println(mySqrt(5))
	println(mySqrt(8))
	println(mySqrt(13))
}
