package main

var v1 []int = []int{-1, 0, 3, 5, 9, 12}
var v2 []int = []int{-1, 0, 3, 5, 9, 12}
var v3 []int = []int{2, 5}
var v4 []int = []int{-1, 0, 3, 5, 9, 12}

func search(nums []int, target int) int {
	l := len(nums)

	start := 0
	end := l - 1

	for {
		if start > end {
			return -1
		}

		mid := ((end - start) / 2) + start

		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			start = mid + 1
		}
		if nums[mid] > target {
			end = mid - 1
		}
	}
}

func main() {
	println(search(v1, 9))
	println(search(v2, 2))
	println(search(v3, 5))
	println(search(v4, 13))
}
