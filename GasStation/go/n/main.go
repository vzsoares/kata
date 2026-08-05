package main

func canCompleteCircuit(gas []int, cost []int) int {

	total, tank, candidate := 0, 0, 0

	for i := range gas {
		surplus := gas[i] - cost[i]

		total += surplus
		tank += surplus

		// can't reach i+1 from candidate, so no start in [candidate, i] works;
		// the next possible start is i+1
		if tank < 0 {
			candidate = i + 1
			tank = 0
		}
	}

	if total >= 0 {
		return candidate
	}

	return -1
}

func main() {

	println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) //expected output: 3
	println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))             //expected output: -1
	println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1})) //expected output: 4
	println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))             //expected output: 0
}
