package main

import "fmt"

func findJudge(n int, trust [][]int) int {

	return n
}

func main() {

	fmt.Println(findJudge(2, [][]int{{1, 2}}))                 // Output: 2
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))         // Output: 3
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}})) // Output: -1
}
