package main

func twoSum(nums []int, target int) []int {
	count := len(nums)

	// value: index
	memory := make(map[int]int)

	for i := 0; i < count; i++ {
		num := nums[i]

		complement := target - num

		seen, ok := memory[complement]

		if ok {
			return []int{seen, i}
		}

		// add to memory after to dedupe
		memory[num] = i
	}
	return []int{}
}
