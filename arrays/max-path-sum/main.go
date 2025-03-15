package main

import "fmt"

func main() {
	arr1 := []int{12, 13, 18, 20, 22, 26, 70}
	arr2 := []int{11, 15, 18, 19, 20, 26, 30, 31, 100}
	fmt.Println("MaxPathSum: ", MaxPathSum(arr1, arr2))

	// arr3 := []int{2, 3, 7, 10, 12, 15, 30, 34}
	// arr4 := []int{1, 5, 7, 8, 10, 15, 16, 19}
	// fmt.Println("MaxPathSum: ", MaxPathSum(arr3, arr4))
}

func MaxPathSum(arr1, arr2 []int) int {
	i, j := 0, 0
	sum_i, sum_j := 0, 0
	total := 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] && i != len(arr1) {
			sum_i += arr1[i]
			i++
		} else if arr1[i] > arr2[j] {
			sum_j += arr2[j]
			j++
		} else {
			sum_i += arr1[i]
			sum_j += arr2[j]
			if sum_i < sum_j {
				total += sum_j
			} else {
				total += sum_i
			}
			fmt.Println("total: ", total)
			sum_i = 0
			sum_j = 0
			i++
			j++
		}
	}

	for i < len(arr1) {
		sum_i += arr1[i]
		i++
	}
	for j < len(arr2) {
		sum_j += arr2[j]
		j++
	}
	if sum_i < sum_j {
		total += sum_j
	} else {
		total += sum_i
	}

	return total
}

// Suppose we have the two following arrays, and they are ORDERED
// [12, 13, 18, 20, 22, 26, 70]
// [11, 15, 18, 19, 20, 26, 30, 31]

// We are interested in taking the larger sum between two transition points, the number of elements doesn't matter
// The transition points in the above two array happen at values 18, 20, 26.
// We don't need to know all the transition points before traversing, we can discover tham and keep them in memory as we go down the arrays.

// Initiate two pointers, i and j, for each array respectively. Then we can take advantage of the arrays being ordered.
// We'll also have two sum, sum_i and sum_j, to track the current sum then reset it every time we hit a transition point.
// Set both i, j = 0, 0 then check the condition arr[i] < arr[j]
// Then we increment the counter of whichever is smaller, until we hit a transition point at which case they will be exactly equal
