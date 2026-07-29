package main

import (
	"cmp"
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	result := recurse(intervals, 0)

	return result
}

func recurse(arr [][]int, idx int) [][]int {

	if idx+1 > len(arr)-1 {
		return arr[:]
	}

	curr := arr[idx]
	next := arr[idx+1]

	//can merge
	if curr[1] >= next[0] {
		// merge and keep idx
		first := max(idx-1, 0)

		result := arr
		if first == 0 {
			result = [][]int{{curr[0], next[1]}}
		} else {
			result = append(arr[:first], []int{curr[0], next[1]})
		}
		last := min(idx+2, len(arr)-1)
		if len(arr) > 2 {
			result = append(result, arr[last:]...)
		}
		return recurse(result, idx)
	} else {
		// idx ++
		return recurse(arr, idx+1)
	}
}

func main() {
	//[[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	//[[1,7]]
	// fmt.Println(merge([][]int{{4, 7}, {1, 4}}))
	//[[0,0],[1,4]]
	// fmt.Println(merge([][]int{{1, 4}, {0, 0}}))
	//[[0,5]]
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
}
