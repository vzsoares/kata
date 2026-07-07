package main

import "fmt"

func rotate(nums []int, k int) {

	start := k % len(nums)
	if start == 0 {
		return
	}
	cp := make([]int, len(nums))
	copy(cp, nums)

	for i := range nums {
		nums[start] = cp[i]
		start++
		if start >= len(nums) {
			start = 0
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	fmt.Println(arr)
}
