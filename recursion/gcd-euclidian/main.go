package main

import "fmt"

func main() {
	fmt.Println(GCD(35, 49))
}

func GCD(n, m int) int {
	if n == 0 {
		return m
	}
	if m == 0 {
		return n
	}
	return GCD(m, n%m)
}

// https://proofwiki.org/wiki/Euclidean_Algorithm
