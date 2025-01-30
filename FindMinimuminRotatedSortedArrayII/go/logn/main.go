package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

func findMin(nums []int) int {
	l := len(nums)
	lo, hi := 0, l-1

	for lo < hi {
		mid := ((hi - lo) / 2) + lo

		if nums[mid] < nums[hi] {
			hi = mid
		} else if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi--
		}

	}

	return nums[hi]
}

func main() {
	println(findMin([]int{1, 3, 5}), "should 1")
	println("")
	println(findMin([]int{2, 2, 2, 0, 1}), "should 0")
	println("")
	println(findMin([]int{3, 3, 1, 3}), "should 1")
	println("")
	println(findMin([]int{1, 3, 3}), "should 1")
	println("")
	println(findMin([]int{-1, -1, -1, -1}), "should -1")
	println("")
	println(findMin([]int{3, 3, 3, 3, 3, 3, 3, 3, 1, 3}), "should 1")
	println("")
	println(findMin([]int{3, 3, 1, 3, 3, 3, 3}), "should 1")
	println("")
	println(findMin([]int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 2, 3}), "should 1")
	println("")
}
