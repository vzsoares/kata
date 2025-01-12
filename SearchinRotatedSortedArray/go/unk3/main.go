package main

func search(nums []int, target int) int {
	l := len(nums)

	start := 0
	end := l - 1

	for start <= end {
		mid := ((end - start) / 2) + start

		if nums[end] <= nums[mid] {
			start = mid + 1
		} else {
			end = mid
		}

	}
	pivot := start - 1
	// println("pivot", pivot)
	lo := 0
	hi := l - 1
	for lo <= hi {
		m := (hi + lo) / 2
		rm := (m + pivot) % l

		if nums[rm] == target {
			return rm
		} else if nums[rm] > target {
			hi = m - 1
		} else {
			lo = m + 1
		}
	}

	return -1
}

func main() {
	println(search([]int{0, 1, 3, 5, 6, 8, 12, 15}, 3), "should 2, 0")
	println("")
	println(search([]int{0, 1, 3, 5, 6, 8, 12, 15}, 15), "should 7, 0")
	println("")
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), "should 4, 4")
	println("")
	println(search([]int{6, 7, 0, 1, 2, 4, 5}, 0), "should 2, 2")
	println("")
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3), "should -1, 4")
	println("")
	println(search([]int{1}, 0), "should -1, 0")
	println("")
	println(search([]int{5, 1, 3}, 5), "should 0, 1")
	println("")
	println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0), "should 4, 4")
	println("")
	println(search([]int{1, 3}, 1), "should 0, 0")
	println("")
	println(search([]int{1, 3, 5}, 3), "should 1, 0")
}
