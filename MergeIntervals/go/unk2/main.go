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
		return arr
	}

	curr := arr[idx]
	next := arr[idx+1]

	//can merge
	if curr[1] >= next[0] {
		// merge and keep idx
		result := append(arr[:max(idx-1, 0)], []int{curr[0], next[1]})
		result = append(result, arr[min(idx+2, len(arr)-1):]...)
		recurse(result, idx)
	} else {
		// idx ++
		return recurse(arr, idx+1)
	}

	return arr
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
