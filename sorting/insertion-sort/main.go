package main

import "fmt"

func main() {
	InsertionSort([]int{10, 5, 1, 8, 2, 3, 9, 6, 4, 7})
}

func InsertionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
	fmt.Println(arr)
}

// The idea behind insertion sort is to keep two iterators
// The first tracks the size of a sub array on the left of the array, call it i
// The second tracks the position of the rightmost element of the subarray
// The idea is to ensure that the left subarray is always correctly ordered
// To do this we set i = 0 and iterate up, but we set j = i for each iteration of i and iterate down, propagating the rightmost element into it's rightful place
// We do this by swapping the j th element with the j-1 th element if j < j-1
// There is no optimnisation like bubble sort since we need to run all iterations over i to ensure we have evaluated all elements in the array
// The time complexity is O(n^2) i.e. sum (n)(n-1) + (n-1)(n-2)
