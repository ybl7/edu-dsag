package main

import "fmt"

func main() {
	ArrayPermutation([]int{1, 2, 3, 4}, 0)
	ArrayPermutation([]int{1, 2, 3, 4, 5}, 0)
	ArrayPermutation([]int{1, 2, 3, 4, 5, 6}, 0)
}

func ArrayPermutation(arr []int, j int) {
	// arr stays the same length the entire time, j tracks which index we are permutating currently
	if len(arr) == j {
		fmt.Println(arr)
		return
	}
	for i := j; i < len(arr); i++ {
		arr2 := Swap(arr, j, i)
		ArrayPermutation(arr2, j+1)
		// This final swap back is crucial, because in the next iteration of the loop we want to operate on the original arr passed into this call of ArrayPermutation, then do a swap for the next index i
		arr = Swap(arr2, j, i)
	}
}

func Swap(arr []int, i, j int) []int {
	arr[i], arr[j] = arr[j], arr[i]
	return arr
}

// Given some array [a,b,c,d], how can we find all possible permutations?
// For any given index, every element has to occupy it at some point. For example for the 0th position.
// By swapping the 0th position with all others (including itself), we can ensure that we have cycled all elements through it

// [a,b,c,d]
// [b,a,c,d]
// [c,b,a,d]
// [d,b,c,a]

// Now for each of these arrays, we call our procedure again on the remaining elements, and we keep doing this

// 	[a,b,c,d]
// 		[a,b,c,d]
// 			[a,b,c,d]
// 			[a,b,d,c]
// 		[a,c,b,d]
// 			[a,c,b,d]
// 			[a,c,d,b]
// 		[a,d,c,b]
// 			[a,d,c,b]
// 			[a,d,b,x]

// Let's call n := len(arr)

// level 0
// Upper most call
// O(n) work done outside of recursive call since we iterate over the whole array, 0 recursive cals

// Level 1
// n recursive calls
// O(n-1) work done outside the recursive call due to the loop that excludes the first element

// Level 2
// n-1 recursive calls
// O(n-2) work done outside the recursive calls

// Level 3
// n-2 recursive calls
// O(n-3) work done outside the recurseive calls

// The recurrence relation is T(n) = n*T(n-1) + O(1)

// This algorithm is called Heaps Algorithm https://en.wikipedia.org/wiki/Heap%27s_algorithm#Proof
// See slide 16 https://sedgewick.io/wp-content/uploads/2022/03/2002PermGeneration.pdf for proof of the running time
