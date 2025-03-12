package main

import "fmt"

func main() {
	fmt.Println(MaxMinArr([]int{1, 2, 3, 5, 6, 8, 20}))
	fmt.Println(MaxMinArr([]int{1, 2, 3, 4, 5, 6, 8, 11, 12, 13, 15, 20}))
}

func MaxMinArr(arr []int) []int {
	// Make the slice in memory at the start since we already know how long it need to be
	out := make([]int, len(arr))

	i, j := 0, len(arr)-1
	for k := 0; k < len(arr); k++ {
		if k%2 == 0 {
			out[k] = arr[j]
			j--
		} else {
			out[k] = arr[i]
			i++
		}
	}
	return out
}

// We have an ordered array
// [1,2,3,5,6,8,20]

// We want to have even indexes increasing in value and odd indexes decreasing in value.
// So element 0 must be the largest element of all. Since the array is sorted, we know this is the final element.

// Let's think of the array as right and a left subarray, ensuring the right subarray is always larger or equal to the left in length.
// L = [1,2,3]
// R = [5,6,8,20]

// Then reverse the right array
// R = [20,8,6,5]

// Then create a new array and populate it by walking down each array, even indexes populated from R and odd from L.
// [20] -> [20,1] -> [20,1,8]
// In practice if we have two pointers, one walking up the array and one down it, this accomplishes the stask.
