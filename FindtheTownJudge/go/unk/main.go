package main

import (
	"fmt"
)

func findJudge(n int, trust [][]int) int {
	//[[1,3],[2,3]]
	//players 1,2,3
	//payers trusted by; [ [1,[]], [2,[]], [3,[1,2]] ]
	//judge == player trusted by all and not truster

	trusteds := make(map[int]map[int]bool)
	trusters := make(map[int]bool)

	for _, pair := range trust {
		if trusteds[pair[1]] == nil {
			trusteds[pair[1]] = make(map[int]bool)
		}
		trusteds[pair[1]][pair[0]] = true
		trusters[pair[0]] = true
	}

	for player := 1; player < n+1; player++ {
		trusted := trusteds[player]
		if len(trusted) == n-1 && trusters[player] == false {
			return player
		}
	}

	return -1
}

func main() {

	fmt.Println(findJudge(2, [][]int{{1, 2}}))                 // Output: 2
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))         // Output: 3
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})) // Output: -1
	fmt.Println(findJudge(3, [][]int{{1, 2}, {2, 3}}))         // Output: -1
}
