package main

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)

	for range k {
		for i := len(nums) - 1; i >= 1; i-- {
			tmp := nums[i-1]
			nums[i-1] = nums[i]
			nums[i] = tmp
		}
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(arr, 3)
	fmt.Println("[5,6,7,1,2,3,4]")
	fmt.Println(arr)

	arr = []int{-1, -100, 3, 99}
	rotate(arr, 2)
	fmt.Println("[3,99,-1,-100]")
	fmt.Println(arr)

	arr = []int{1, 2, 3}
	rotate(arr, 1)
	fmt.Println("[3,1,2]")
	fmt.Println(arr)

	arr = []int{1, 2}
	rotate(arr, 1)
	fmt.Println("[2,1]")
	fmt.Println(arr)
}

