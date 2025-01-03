package main

func twoSum(nums []int, target int) []int {
	count := len(nums)

	// value: index
	m := make(map[int]int)

	for i := 0; i < count; i++ {
		num := nums[i]

		goal := target - num

		saved, ok := m[goal]

		if ok {
			return []int{saved, i}
		}

		m[num] = i
	}
	return []int{}
}
