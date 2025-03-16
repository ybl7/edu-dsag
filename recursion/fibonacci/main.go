package main

import "fmt"

func main() {
	fmt.Println(FibIte(10))
	fmt.Println(FibRec(10))
}

// Iterative is much faster since you don't need to provision a new level of stack with each function call
func FibIte(n int) int {
	last := 1
	curr := 1

	for i := 2; i < n; i++ {
		tmp := curr
		curr += last
		last = tmp
	}
	return curr
}

func FibRec(n int) int {
	if n == 2 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return FibRec(n-1) + FibRec(n-2)
}

// Assuming it starts from 1, 1 not 0, 1
