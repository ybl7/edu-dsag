package main

import "fmt"

func main() {
	fmt.Println(RotateByK([]int{1, 2, 3, 4, 5, 6, 7, 8}, 4))
	fmt.Println(RotateByK([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6))
}

func RotateByK(arr []int, k int) []int {
	ReverseInPlace(arr[0:k])
	ReverseInPlace(arr[k:])
	ReverseInPlace(arr)
	return arr
}

// O(n) in time since it loops over arr n/2 times performing 3 constant operations in each loop
// O(1) in space since a constant number of variables are created then cleared each loop iteration independently of the length of n
func ReverseInPlace(arr []int) {
	l := len(arr) / 2

	for i := 0; i <= l; i++ {
		tmp := arr[i]
		arr[i] = arr[len(arr)-i-1]
		arr[len(arr)-i-1] = tmp
	}
}

// Doing this in O(n) time and space is easy, just make a second array and copy elements into it at the desired index
// Can we do this in constant space? If we are to attempt this then we'll need to rely on swapping elements in the array.
// We also know that it is possible to reverse an array through swapping only i.e. to go from [a,b,c,d] to [d,b,c,a] to [d,c,b,a].

// Suppose we have an array [a,b,c,d,e,f,g,h,i] and we want to rotate it by 3 producing output array [d,e,f,g,h,i,a,b,c]
// If we split the array about the rotation value 3, we can imagine two sub arrays of [a,b,c] and [d,e,f,g,h,i]
// Now if we reverse each array we end up with [c,b,a] and [i,h,g,f,e,d]
// Now if we reverse the whole array we end up with [d,e,f,g,h,i] and [a,b,c]
// This will be our algorithm

// Note that integer division rounds down in Go
// Go passes everything by value, what's actually being passed here is a slice descriptor, the underlying array is not being recreated at any point so we are O(1) in memory
