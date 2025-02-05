package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

func intersect(nums1 []int, nums2 []int) []int {
	seen := make(map[int]int)
	set := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		curr := nums1[i]
		seen[curr] = seen[curr] + 1
	}

	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	for i := 0; i < len(nums2); i++ {
		curr := nums2[i]

		if seen[curr] > 0 {
			set = append(set, curr)
			seen[curr] = seen[curr] - 1
		}
	}
	return set
}

func main() {
	fmt.Printf("%#v %#v\n", intersect([]int{1, 2, 3, 4}, []int{1, 2, 4}), "[1, 2, 4]")
	fmt.Printf("%#v %#v\n", intersect([]int{1, 1, 1, 1, 2, 3, 4}, []int{1, 1, 2, 4}), "[1, 1, 2, 4]")
	fmt.Printf("%#v %#v\n", intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), "[4, 9]")
}
