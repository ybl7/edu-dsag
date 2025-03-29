package main

import "fmt"

func main() {
	fmt.Println(MergeSort([]int{10, 5, 1, 8, 2, 3, 9, 6, 4, 7}))
}

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			arr[0], arr[1] = arr[1], arr[0]
		}
	}
	mid := len(arr) / 2

	l := MergeSort(arr[:mid])
	r := MergeSort(arr[mid:])

	// We already know the length so we can just pre-set the slice capacity
	out := make([]int, 0, len(arr))

	i, j := 0, 0
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			out = append(out, l[i])
			i++
		} else {
			out = append(out, r[j])
			j++
		}
	}

	// Append whatever is left, one of the arrays will already be completely travered
	// So only the larger one will remain
	for i < len(l) {
		out = append(out, l[i])
		i++
	}
	for j < len(r) {
		out = append(out, r[j])
		j++
	}

	return out
}

// The idea behind merge sort is to split an array recursively into smaller arrays, then sort the smaller arrays
// Then we combine all the smaller arrarys in the merge step, see Roughgarden's book for details
// Here's an example to illustrate: suppose I have an array [3,2,1]
// Split it into [3] and [2,1], order these since these are our base cases of simplest arrays, so we get [3] and [1,2]
// Then create a temporary array and insert values into it, i.e. merge the two sorted subarrays
// We do this by running a pointer down each array, so it looks like:

// i = 0
// j = 0
// Compare arr1[i] and arr2[j] and insert whichever is smaller into the temp array
// Then incremenent whichever pointer was just used in the insertion, then repeat

// See Roughgarden's book for time complexity, but its O(nlog(n))
