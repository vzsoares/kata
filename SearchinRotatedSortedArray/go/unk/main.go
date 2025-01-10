package main

func search(nums []int, target int) int {
	l := len(nums)

	start := 0
	end := l - 1

	for start <= end {
		mid := ((end - start) / 2) + start

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return -1
}

func main() {
	println(search([]int{0, 1, 3, 5, 6, 8, 12, 15}, 3), "should 2")
	println(search([]int{0, 1, 3, 5, 6, 8, 12, 15}, 15), "should 7")
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), "should 4")
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), "should -1")
	println(search([]int{1}, 0), "should -1")
}
