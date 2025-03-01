package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(binary_search([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5))
	fmt.Println(binary_search([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 50))
}

func binary_search(input []int, target int) bool {
	sort.Ints(input)

	low := 0
	high := len(input) - 1
	var mid int

	for low < high {
		mid = int(math.Floor(float64((low + high) / 2)))

		if input[mid] == target {
			return true
		}
		if input[mid] < target {
			low = mid + 1
		}
		if input[mid] > target {
			high = mid - 1
		}
	}
	return false
}

// By the Master Method
// T(n) = aT(n/b) + O(n^d), where
// a is the number of recursive calls spawned in each iteration,
// b is the factor by which input size decreases in each iteration, and
// O(n^d) is the Big-O of the work done outside of the recursive call

// a = 1
// b = 2
// d = 0
// So a = b^d and the Master Method reduces to log2(n)
