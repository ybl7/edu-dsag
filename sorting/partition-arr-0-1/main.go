package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0, 0, 1}
	fmt.Println(PartitionArrZeroOne(arr))

	arr2 := []int{1, 0, 1, 0, 1, 0, 1}
	fmt.Println(PartitionArrZeroOne(arr2))

	arr3 := []int{0, 1, 0, 0, 1, 0, 0}
	fmt.Println(PartitionArrZeroOne(arr3))

	arr4 := []int{0, 1, 1, 1}
	fmt.Println(PartitionArrZeroOne(arr4))
}

func PartitionArrZeroOne(arr []int) (int, []int) {
	i, j := 0, len(arr)-1
	swaps := 0

	for i < j {
		if arr[i] == 1 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
			swaps++
		}
		if arr[i] == 0 {
			i++
		}
		if arr[j] == 1 {
			j--
		}
	}

	return swaps, arr
}

// We want to partition the array such that all 0s come before all 1s, and we want to accomplish this in the minimum number of swaps.
// If we run two pointers, i from the start and j from the end of the array, we will just swap elements i and j when the condition i = 1 and j = 0 is met
// We increase the values of i and j when i = 0 and j = 1.
