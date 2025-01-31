package main

//https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	r := float64(x)
	absn := n
	if n < 0 {
		absn = n * -1
	}
	for i := absn; i > 1; i-- {
		r = r * x
	}
	if n < 0 {
		return 1 / r
	}
	return r
}

func main() {
	println(myPow(2, 3))
	println(myPow(2, -2))
	println(myPow(1, 2147483647))
}
