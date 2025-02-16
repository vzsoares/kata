package main

import "fmt"

// https://leetcode.com/problems/find-the-duplicate-number/

func findDuplicate(nums []int) int {
	l := len(nums)

	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if nums[i] == nums[j] {
				return nums[i]
			}
		}
	}

	return -1
}

func main() {

	fmt.Printf("%#v %#v\n", findDuplicate([]int{1, 3, 4, 2, 2}), 2)
	fmt.Printf("%#v %#v\n", findDuplicate([]int{3, 1, 3, 4, 2}), 3)
	fmt.Printf("%#v %#v\n", findDuplicate([]int{3, 3, 3, 3, 3}), 3)
}
