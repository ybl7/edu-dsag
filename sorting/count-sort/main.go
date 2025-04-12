// Count sort works well if you already know the range of data you are sorting.
// You create a counter array that is the same as the length as the range of data you are sorting, call it k.
package main

import "fmt"

func main() {
	CountSort([]int{2, 6, 4, 1, 5, 8, 1, 4, 6, 1}, 1, 8)
}

// k is the range of data
func CountSort(arr []int, low, high int) {
	ctr := make([]int, high-low+1)
	// We are going to try and represent a range low:high as a range 0:high-low, since our array is indexed from 0
	// Therefore when we populate ctr, we should subtract low from every value, and when we populate the final array we must add low back to get the original value

	for _, j := range arr {
		ctr[j-low]++
	}
	fmt.Println(ctr)

	l := 0
	for i, j := range ctr {
		for j > 0 {
			arr[l] = i + low
			j--
			l++
		}
	}
	fmt.Printf("Sorted slice: %v\n", arr)
}
