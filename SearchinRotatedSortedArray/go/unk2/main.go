package main

// https://leetcode.com/problems/search-in-rotated-sorted-array/description/
func search(nums []int, target int) int {
	//TODO find lesser value as it means the pivot position

	//TODO search value and consider the pivot
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
}
