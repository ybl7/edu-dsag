package main

import "fmt"

func main() {
	arr := []int{6, 8, 9, 1, 4, 12, 2, 7, 4, 11, 10}
	fmt.Println("Missing Number :", SmallestPositiveMissingNumber(arr))

}

func SmallestPositiveMissingNumber(arr []int) int {
	m := map[int]int{}

	for i := 0; i < len(arr); i++ {
		m[arr[i]] = 1
	}

	for i := 1; i < len(arr)+1; i++ {
		_, ok := m[i]
		if ok == false {
			return i
		}
	}

	return -1
}

// If we have a sorted array like the series above, this is trivial, but we aren't given a sorted array.
// And sorting an array would cost us at least nlog(n), so is there a way we can do this in O(n)?
// We can do it if we provision another array/hash table, making this O(n) in both time and space.

// This problem boils down to finding the first empty index in an array.
// let's iterate down this array and populate a hash table {6, 8, 9, 1, 12, 2, 7, 4, 11, 10}

// i = 0
// {
// 	6: 1
// }

// i = 2
// {
// 	6: 1
// 	8: 1
// }

// ...

// i = len(arr)
// {
// 	1: 1
// 	2: 1
// 	4: 1
// 	6: 1
// 	7: 1
// 	8: 1
// 	9: 1
// 	10: 1
// 	11: 1
// 	12: 1
// }

// Then iterate from i to the length of the array, checking if the corresponsing key in the map eists
