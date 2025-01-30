package main

// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/

func findMin(nums []int) int {
	l := len(nums)
	lo, hi := 0, l-1

	inOnce := false
	for lo < hi {
		mid := ((hi - lo) / 2) + lo

		if nums[hi] == nums[mid] && ((hi - mid) > 1) && !inOnce {
			//find if are all equal
			loI := mid
			hiI := hi
			println("IN", loI, hiI)
			// panic("")
			for loI < hiI {
				midI := ((hiI - loI) / 2) + loI
				if nums[midI] < nums[hiI] {
					//if found lesser value set lo
					loI = loI + 1
					lo = midI
				} else {
					hiI = hiI - 1
				}
				println("IN", loI, hiI, midI)
				inOnce = true
			}
		} else if nums[hi] < nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
		println("OUT")

	}

	return nums[lo]
}

func main() {
	// println(findMin([]int{1, 3, 5}), "should 1")
	// println("")
	// println(findMin([]int{2, 2, 2, 0, 1}), "should 0")
	// println("")
	// println(findMin([]int{3, 3, 1, 3}), "should 1")
	// println("")
	// println(findMin([]int{1, 3, 3}), "should 1")
	// println("")
	// println(findMin([]int{-1, -1, -1, -1}), "should -1")
	// println("")
	println(findMin([]int{3, 3, 3, 3, 3, 3, 3, 3, 1, 3}), "should 1")
	println("")
}
