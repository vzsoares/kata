package main

// https://leetcode.com/problems/find-peak-element/

func findPeakElement(nums []int) int {
	l := len(nums)
	lo, hi := 0, l-1

	for lo < hi {
		mid := ((hi - lo) / 2) + lo

		if nums[mid] < nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	println("hi lo mid", hi, lo)

	return lo
}

func main() {
	println("res", findPeakElement([]int{1, 2, 3, 1}), "expected", 2)
	println("res", findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}), "expected", 5, 1)
	println("res", findPeakElement([]int{1, 2}), "expected", 1)
	println("res", findPeakElement([]int{1, 3, 2, 1}), "expected", 1)
}
