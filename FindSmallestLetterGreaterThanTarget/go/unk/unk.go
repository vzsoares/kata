package main

import "fmt"

// https://leetcode.com/problems/find-smallest-letter-greater-than-target/

func nextGreatestLetter(letters []byte, target byte) byte {
	l := len(letters)
	lo, hi := 0, l-1

	for lo < hi {
		m := ((hi - lo) / 2) + lo

		if letters[m] > target {
			hi = m
		} else {
			lo = m + 1
		}
	}

	if letters[lo] > target {
		return letters[lo]
	}

	return letters[0]
}

func main() {
	fmt.Printf("%c %s\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'), "c")
	fmt.Printf("%c %s\n", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c'), "f")
	fmt.Printf("%c %s\n", nextGreatestLetter([]byte{'x', 'x', 'y', 'y'}, 'z'), "x")
}
