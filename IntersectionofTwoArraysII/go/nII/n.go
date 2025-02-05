package main

import (
	"fmt"
	"sort"
)

// https://leetcode.com/problems/intersection-of-two-arrays-ii/description/

func intersect(nums1 []int, nums2 []int) []int {
	set := make([]int, 0)

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	p1 := 0
	p2 := 0
	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] == nums2[p2] {
			// if they shouldnt repeat then use a map
			set = append(set, nums1[p1])
			p1 = p1 + 1
			p2 = p2 + 1
		} else if nums1[p1] < nums2[p2] {
			p1 = p1 + 1
		} else {
			p2 = p2 + 1
		}
	}

	return set
}

func main() {
	fmt.Printf("%#v %#v\n", intersect([]int{1, 2, 3, 4}, []int{1, 2, 4}), "[1, 2, 4]")
	fmt.Printf("%#v %#v\n", intersect([]int{1, 1, 1, 1, 2, 3, 4}, []int{1, 1, 2, 4}), "[1, 1, 2, 4]")
	fmt.Printf("%#v %#v\n", intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}), "[4, 9]")
}
