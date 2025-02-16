package main

import "fmt"

// https://leetcode.com/problems/find-the-duplicate-number/

func findDuplicate(nums []int) int {
	mem := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		_, ok := mem[curr]
		if ok {
			return curr
		}
		mem[curr] = true
	}

	return -1
}

func main() {

	fmt.Printf("%#v %#v\n", findDuplicate([]int{1, 3, 4, 2, 2}), 2)
	fmt.Printf("%#v %#v\n", findDuplicate([]int{3, 1, 3, 4, 2}), 3)
	fmt.Printf("%#v %#v\n", findDuplicate([]int{3, 3, 3, 3, 3}), 3)
}
