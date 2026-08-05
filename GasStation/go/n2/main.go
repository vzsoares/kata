package main

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)

	for i := range l {
		tank := 0
		for j := i; j < l+i; j++ {
			J := j % l

			tank = tank + gas[J] - cost[J]

			if tank < 0 {
				break
			} else if j == l+i-1 {
				return i
			}
		}
	}
	return -1
}

func main() {

	println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) //expected output: 3
	// println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})) //expected output: -1
	// println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1})) //expected output: 4
	// println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))             //expected output: 0
}
