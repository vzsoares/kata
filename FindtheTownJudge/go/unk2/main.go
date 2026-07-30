package main

import (
	"fmt"
)

func findJudge(n int, trust [][]int) int {
	// the player score needs to be n-1 for him to be the judge
	// arr of scores, idx being the player pos
	// sum scores for winner and if he trusts someone then he CANNOT be the judge
	// for [1,3],[2,3],[3,1] then [-1, 0, 0 ]; no score 1 then -1
	// for [1,3],[2,3] then [-1, -1, 2]; score 2(n-1) then 3(player)
	// sum score and find judge

	if n == 1 {
		return 1
	}

	scores := make([]int, n+1)

	for _, pair := range trust {
		scores[pair[0]] = -67 // this is just to disqualify that player
		scores[pair[1]]++
	}

	for i, el := range scores {
		if el == n-1 {
			return i
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
