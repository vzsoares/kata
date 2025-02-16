package main

import "fmt"

// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

func twoSum(numbers []int, target int) []int {
	l := len(numbers)
	p1 := 0
	p2 := l - 1

	for (p1 < l) && (p2 < l) {
		if (numbers[p1] + numbers[p2]) == target {
			return []int{p1 + 1, p2 + 1}
		}

		// This only works for sorted arrays
		// [1,2,3,4,5,6,7,8,9] -> 12
		//  ^               ^  =  10
		//    ^             ^  =  11
		//      ^           ^  =  12
		//##########################
		// [1,2,3,4,5,6,7,8,9] -> 8
		//  ^               ^  =  10
		//  ^             ^    =  9
		//  ^           ^      =  8
		if numbers[p1]+numbers[p2] > target {
			p2 = p2 - 1
		} else {
			p1 = p1 + 1
		}
	}

	return []int{-1, -1}
}

func main() {

	fmt.Printf("%#v %#v\n", twoSum([]int{2, 7, 11, 15}, 9), []int{1, 2})
	fmt.Printf("%#v %#v\n", twoSum([]int{2, 3, 4}, 6), []int{1, 3})
	fmt.Printf("%#v %#v\n", twoSum([]int{-1, 0}, -1), []int{1, 2})
	fmt.Printf("%#v %#v\n", twoSum([]int{3, 24, 50, 79, 88, 150, 345}, 200), []int{3, 6})
}
