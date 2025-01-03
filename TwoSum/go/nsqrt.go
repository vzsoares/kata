package main

func twoSum(nums []int, target int) []int {
	count := len(nums)

	for i := 0; i < count; i++ {
		a := nums[i]
		for j := i + 1; j < count; j++ {
			b := nums[j]
			if a+b == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
