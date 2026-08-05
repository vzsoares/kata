package main

func canCompleteCircuit(gas []int, cost []int) int {
	l := len(gas)

	J := 0
	for i := 0; i < l; i++ {
		tank := 0
		J = i
		prevCost := cost[i]
		println("start")
		for j := 0; j < l; j++ {
			if j == 0 {
				tank = gas[J]
				prevCost = cost[J]
				continue
			}
			J = getNext(l, J)

			tank = tank + gas[J] - prevCost

			prevCost = cost[J]
			//
			println(J)
			//
			if j == l-1 {
				if tank-cost[getPrev(l, J)] >= 0 {
					return getPrev(l, J)
				}
			}
		}
	}
	return -1
}

func getNext(size int, curr int) int {
	next := curr + 1
	if next > size-1 {
		return 0
	} else {
		return next
	}
}
func getPrev(size int, curr int) int {
	prev := curr - 1
	if prev < 0 {
		return size - 1
	} else {
		return prev
	}
}

func main() {

	// println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) //expected output: 3
	// println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3})) //expected output: 0
	println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1})) //expected output: 4
	// println(canCompleteCircuit([]int{3, 1, 1}, []int{1, 2, 2}))             //expected output: 0
	// println(canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})) //expected output: 3
}
