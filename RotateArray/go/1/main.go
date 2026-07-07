package main

import "fmt"

func rotate(nums []int, k int) {

	if len(nums) <= 1 {
		return
	}

	start := k % len(nums)
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
