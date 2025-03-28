package main

import "fmt"

func main() {
	SelectionSort([]int{10, 5, 1, 8, 2, 3, 9, 6, 4, 7})
}

func SelectionSort(arr []int) {
	j := 0
	for i := 0; i < len(arr); i++ {
		max := 0
		for j = 1; j < len(arr)-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[len(arr)-i-1], arr[max] = arr[max], arr[len(arr)-i-1]
	}
	fmt.Println(arr)
}

// Insertion sort iterates through the array and puts the greatest element at the end
// After the nth iteration the last n elements of the array should be sorted with the largest being the last
// We then iterate over the left subarray that is unordered and swap with the last element on this
