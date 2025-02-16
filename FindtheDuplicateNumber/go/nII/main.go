package main

import "fmt"

// https://leetcode.com/problems/find-the-duplicate-number/
// https://www.youtube.com/watch?v=wjYnzkAhcNk

// ### nums.length == n + 1 ###
// ### nums[n] in [1, n]    ###
// this means each VALUE points to another AND theres only one duplicate
// so we CAN treat the values as pointers in a linked list
// conclusion use Floyd's Cycle Detection
func findDuplicate(nums []int) int {
	slow, fast := 0, 0 //nums[0] will always be HEAD
	// find first intersection
	for {
		slow = nums[slow]       //jumps one
		fast = nums[nums[fast]] //jumps two
		if slow == fast {
			break
		}
	}

	fast = 0 // back to start
	// find intersection again to get the start of the cylcle
	for {
		slow = nums[slow]
		fast = nums[fast]
		if slow == fast {
			break
		}
	}
	return fast
}

func main() {

	fmt.Printf("%#v %#v\n", findDuplicate([]int{1, 3, 4, 2, 2}), 2)
	fmt.Printf("%#v %#v\n", findDuplicate([]int{3, 1, 3, 4, 2}), 3)
	fmt.Printf("%#v %#v\n", findDuplicate([]int{3, 3, 3, 3, 3}), 3)
}
