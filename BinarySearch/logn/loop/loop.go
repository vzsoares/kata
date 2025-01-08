package main

var v1 []int = []int{-1, 0, 3, 5, 9, 12}
var v2 []int = []int{-1, 0, 3, 5, 9, 12}
var v3 []int = []int{2, 5}
var v4 []int = []int{-1, 0, 3, 5, 9, 12}

func search(nums []int, target int) int {
	l := len(nums)

	start := 0
	end := l - 1

	return binary(nums, start, end, target)
}

func binary(nums []int, start, end, target int) int {

	if start > end {
		return -1
	}

	mid := ((end - start) / 2) + start
	el := nums[mid]

	if el == target {
		return mid
	}
	if el < target {
		return binary(nums, mid+1, end, target)
	}
	if el > target {
		return binary(nums, start, mid-1, target)
	}
	return -1
}

func main() {
	println(search(v1, 9))
	println(search(v2, 2))
	println(search(v3, 5))
	println(search(v4, 13))
}
