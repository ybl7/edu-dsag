package main

import "fmt"

func main() {
	BubbleSort([]int{10, 5, 1, 8, 2, 3, 9, 6, 4, 7})
}

func BubbleSort(arr []int) {
	swapped := true
	for i := 0; i < len(arr)-1 && swapped; i++ {
		swapped = false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				swapped = true
			}
		}
	}
	fmt.Println(arr)
}

// We take the 2 element window arr[j] to arr[j+1] and move it down the array in the inner loop
// The value of j runs up to the length of the array minus the current value of i, bubble sort guarantees that the last n elements will be sorted after n runs
// The number of runs is controlled by i, and it can be shown that at max we need len(arr) runs
// So i is really just a counter of runs, and j is used to create the sliding window to perform the bubble sort, hence why j always starts from 0

// The only difference with modified bubble sort compared to regular bubble sort is that we stop iterating through the array once we have ordered everything
// This means we track when no swaps have happened during some iteration of i
