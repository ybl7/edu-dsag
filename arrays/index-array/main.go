package main

import "fmt"

func main() {
	arr1 := []int{8, -1, 6, 1, 9, 3, 2, 7, 4, -1}
	fmt.Println("before: ", arr1)
	IndexArray(arr1)
	fmt.Println("after: ", arr1)
}

func IndexArray(arr []int) {
	for i := 0; i < len(arr)-1; {
		if arr[i] != i && arr[i] != -1 {
			tmp := arr[i]
			tmp2 := arr[arr[i]]

			arr[i] = tmp2
			arr[tmp] = tmp
		}
		if arr[i] == -1 {
			i++
		}
		if arr[i] == i {
			i++
		}
	}
}

// We want to output a sorted integer array with -1 if the element A[i]=i doesn't exist
// No elements are greater than len(array-1)

// Input
// array = [8, -1, 6, 1, 9, 3, 2, 7, 4, -1]
// Output
// array = [-1, 1, 2, 3, 4, -1, 6, 7, 8, 9]

// -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Approach 1
// Let's try to do this in O(1) space.
// We'll iterate through the array and swap the value at A[i] with it's position, i.e. if A[0]=8, we'll swap A[0] <> A[8].
// Then we'll check if the new value at A[i]=i, if not, we wil continue swapping until this condition is true or until we hit -1.
// At which point we will move down the array by 1 index. This works because with each swap we guarantee at least one element is in it's correct index.
// So it saves us work later down the line since by the time we get to later indices in the array, they are likely already in the correct place.

// i=0
// [8, -1, 6, 1, 9, 3, 2, 7, 4, -1]

// A[0] <> A[8]
// [4, -1, 6, 1, 9, 3, 2, 7, 8, -1]

// A[0] == 4 != 0 so continue swapping
// A[0] <> A[4]
// [9, -1, 6, 1, 4, 3, 2, 7, 8, -1]

// A[0] == 9 != 0 so continue swapping
// A[0] <> A[9]
// [-1, -1, 6, 1, 4, 3, 2, 7, 8, 9]

// A[0] = -1 so i=1
// A[1] = -1 so i=2

// A[2] == 6 != 2
// A[2] <> A[6]
// [-1, -1, 2, 1, 4, 3, 6, 7, 8, 9]

// A[2] == 2 so i=3

// A[3] == 1 != 3
// A[3] <> A[1]
// [-1, 1, 2, -1, 4, 3, 6, 7, 8, 9]

// A[3] == -1 so i=4
// A[4] == 4 so i=5 (we placed this correctly already when i=0)

// A[5] == 3 != 5
// A[5] <> A[3]
// [-1, 1, 2, 3, 4, -1, 6, 7, 8, 9]

// A[5] != 5 so i=6
// A[6] == 6 so i=7
// A[7] == 7 so i=8
// A[6] == 6 so i=7
// A[7] == 7 so i=8
// A[8] == 8 so i=9
// A[9] == 9 so but we exit when i = n-1

// So what is the time complexity?
// Every time we swap, we are guaranteeing correct placement of 1 element.
// Which means when we do eventually revist it later we do O(1) on it (just the comparison).
// So (hand waving here a bit), we do max O(n) work traversing the array.
// And we also know we do a max of O(n) swaps distributed across all values of i.
// Which leads me to think this is a O(n) algorithm overall.
