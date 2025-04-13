package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr1 := []int{4, 5, 11, 6, 9, 12, 3, 7, 8, 2, 10, 15, 1}
	QuickSort(arr1, 100)
	fmt.Println(arr1)

	arr2 := []int{4, 5, 11, -4, 23, -6, 100, 24, 18, -84, 6, 9, 12, 3, 7, 8, 2, 10, 15, 1}
	QuickSort(arr2, 100)
	fmt.Println(arr2)
}

// n is not required, it's just a limit on the number of recursive calls that helps with debugging when hitting stack overflow errors
func QuickSort(arr []int, n int) {
	if n == 0 {
		return
	}
	n--

	if len(arr) <= 1 {
		return
	}
	piv := Partition(arr)
	QuickSort(arr[0:piv], n)
	QuickSort(arr[piv+1:], n)
}

func Partition(arr []int) int {
	piv := rand.Intn(len(arr))

	// Swap pivot and 0th element
	o := arr[0]
	arr[0] = arr[piv]
	arr[piv] = o

	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] <= arr[0] {
			tmp := arr[j]
			arr[j] = arr[i]
			arr[i] = tmp
			j++
		}
	}

	oo := arr[0]
	arr[0] = arr[j-1]
	arr[j-1] = oo

	return j - 1
}

// We'll skip the explanation of why quicksort works, the formal proof is a bit involved and there is an entire section in Roughgarden's book (Algorithms Illuminated) dedicated to choosing good pivots
// I chose to go with a randomised pivot for simplicity
