package main

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func search(nums []int, target int) int {
	//find lesser value as it means the pivot position
	pivot := findLesserValueIndex(nums)
	// println("pivot", pivot)
	//search value and consider the pivot
	l := len(nums)

	hi := l - 1
	lo := 0

	mid := 0
	realMid := 0

	for lo <= hi {
		mid = lo + ((hi - lo) / 2)
		realMid = (mid + pivot) % l

		if nums[realMid] == target {
			return realMid
		} else if nums[realMid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	if nums[realMid] == target {
		return realMid
	}

	return -1
}

func findLesserValueIndex(nums []int) int {
	l := len(nums)

	hi := l - 1
	lo := 0

	mid := 0

	for lo <= hi {

		mid = ((hi - lo) / 2) + lo

		if nums[hi] <= nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return mid
}

func main() {
	println(search([]int{0, 1, 3, 5, 6, 8, 12, 15}, 3), "should 2, 0")
	println("")
	println(search([]int{0, 1, 3, 5, 6, 8, 12, 15}, 15), "should 7, 0")
	println("")
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), "should 4, 4")
	println("")
	println(search([]int{6, 7, 0, 1, 2, 4, 5}, 0), "should 2, 2")
	println("")
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), "should -1, 4")
	println("")
	println(search([]int{1}, 0), "should -1, 0")
	println("")
	println(search([]int{5, 1, 3}, 5), "should 0, 1")
	println("")
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), "should 4, 4")
	println("")
	println(search([]int{1, 3}, 1), "should 0, 0")
	println("")
	println(search([]int{1, 3, 5}, 3), "should 1, 0")
}
