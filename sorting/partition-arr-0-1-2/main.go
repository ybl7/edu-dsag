package main

import "fmt"

func main() {
	arr := PartitionArrZeroOneTwo([]int{0, 0, 2, 0, 2, 1, 0, 1})
	fmt.Println(arr)

	arr2 := PartitionArrZeroOneTwo([]int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1})
	fmt.Println(arr2)

	arr3 := PartitionArrZeroOneTwo([]int{0, 1, 2, 0, 1, 2})
	fmt.Println(arr3)
}

func PartitionArrZeroOneTwo(arr []int) []int {
	arr, idx := PartitionArrInt(arr, 0, 0, len(arr)-1)
	arr, _ = PartitionArrInt(arr, 1, idx, len(arr)-1)
	return arr
}

// Ensures slice between low and high is partitioned such that tgt shows up at the start of the array, returns the array and index at which tgt subarray ends
func PartitionArrInt(arr []int, tgt, low, high int) ([]int, int) {
	i, j := low, high

	for i < j {
		if arr[i] != tgt && arr[j] == tgt {
			arr[i], arr[j] = arr[j], arr[i]
		}
		if arr[i] == tgt {
			i++
		}
		if arr[j] != tgt {
			j--
		}
	}

	return arr, i
}

// This is just like partiioning an array into zeros and ones, except now we have three numbers, but we don't have to think of it that way.
// If we iterate through the array once, let's get just the 0s in place, not caring about whether the other value is 1 or 2.
// Then once we have all the zeros in place, we know that we just need to iterate over the remainder of the array and do the same procedure but for 1s and 2s.
// This is done a little differently in the solution but I prefer this solution, it's a lot more clear and can easily be extended to partitioning over n values.
