package main

//https://leetcode.com/problems/powx-n/

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	absn := n
	if n < 0 {
		absn = n * -1
	}

	result := Pow(x, absn)

	if n < 0 {
		return 1 / result
	}
	return result
}

func Pow(base float64, expoent int) float64 {
	// 2^10 = 2^5 * 2^5
	//
	// 2^5 = 2 * 2^2 * 2^2
	if expoent == 1 {
		return base
	}
	if expoent%2 == 0 {
		power := Pow(base, expoent/2)
		return power * power
	} else {
		power := Pow(base, (expoent-1)/2)
		return base * power * power
	}
}

func main() {
	println(myPow(2, 3))
	println(myPow(2, -2))
	println(myPow(1, 2147483647))
}
