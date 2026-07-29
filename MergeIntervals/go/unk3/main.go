package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {

	slices.SortFunc(intervals, func(a, b []int) int {
		if a[0] == b[0] {
			return a[1] - b[1]
		}

		return a[0] - b[0]
	})

	acc := make([][]int, 0, len(intervals)-1)

	acc = append(acc, []int{intervals[0][0], intervals[0][1]})

	for i := range intervals {
		if i == 0 {
			continue
		}

		accVal := acc[len(acc)-1]
		currVal := intervals[i]

		if accVal[1] >= currVal[0] {
			accVal[1] = max(accVal[1], currVal[1])
		} else {
			acc = append(acc, []int{currVal[0], currVal[1]})
		}

	}

	return acc
}

func main() {

	// [[1,6],[8,10],[15,18]]
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	//[[1,3],[4,7]]
	fmt.Println(merge([][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}))
}
