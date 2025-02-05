package main

import "fmt"

// https://leetcode.com/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {
	set := make([]int, 0)
	seen := make(map[int]bool)

	for i := 0; i < len(nums1); i++ {
		curr := nums1[i]
		if seen[curr] {
			continue
		}
		seen[curr] = true
	}

	for i := 0; i < len(nums2); i++ {
		curr := nums2[i]
		if seen[curr] {
			set = append(set, curr)
			seen[curr] = false
		}
	}
	return set
}

func main() {
	fmt.Printf("%#v %#v\n", intersection([]int{1, 2, 3, 4}, []int{1, 2, 4}), "[1, 2, 4]")
	fmt.Printf("%#v %#v\n", intersection([]int{1, 1, 1, 1, 2, 3, 4}, []int{1, 2, 4}), "[1, 2, 4]")
}
