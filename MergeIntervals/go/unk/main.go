//DOES NOT WORK

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

	result := [][]int{}

	for i := 0; i < len(intervals); i++ {
		for j := len(intervals) - 1; j >= i; j-- {
			firstEnd := intervals[i][1]
			lastStart := intervals[j][0]

			if lastStart <= firstEnd {
				// merge

				// get smallest interval point
				smallestStart := -1
				if intervals[i][0] < intervals[j][0] {
					smallestStart = intervals[i][0]
				} else {
					smallestStart = intervals[j][0]
				}

				// get largest interval point
				largestEnd := -1
				if intervals[j][1] < intervals[i][1] {
					largestEnd = intervals[i][1]
				} else {
					largestEnd = intervals[j][1]
				}

				result = append(result, []int{smallestStart, largestEnd})
				i = j
			} else {
				// keep
				if j == i+1 || j == i {
					result = append(result, intervals[i])
					break
				}
			}
		}
	}

	return result
}

func main() {
	//[[1,6],[8,10],[15,18]]
	// fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	//[[1,7]]
	// fmt.Println(merge([][]int{{4, 7}, {1, 4}}))
	//[[0,0],[1,4]]
	// fmt.Println(merge([][]int{{1, 4}, {0, 0}}))
	//[[0,5]]
	fmt.Println(merge([][]int{{1, 4}, {0, 2}, {3, 5}}))
}
