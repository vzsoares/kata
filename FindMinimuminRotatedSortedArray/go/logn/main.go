package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

func findMin(nums []int) int {
	l := len(nums)
	lo, hi := 0, l-1

	for lo < hi {
		mid := ((hi - lo) / 2) + lo

		if nums[hi] < nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}

	}

	return nums[lo]
}

func main() {
	println(findMin([]int{3, 4, 5, 1, 2}), "should 1")
	println("")
	println(findMin([]int{4, 5, 6, 7, 0, 1, 2}), "should 0")
	println("")
	println(findMin([]int{11, 13, 15, 17}), "should 11")
	println("")
	println(findMin([]int{2, 3, 4, 5, 1}), "should 1")
	println("")
}
