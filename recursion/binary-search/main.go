package main

import "fmt"

func main() {
	BinarySearchRecursive([]int{0, 1, 2, 3, 4, 5, 6}, 6)
	BinarySearchRecursive([]int{0, 11, 20, 23, 34, 35, 67, 89}, 12)
}

func BinarySearchRecursive(arr []int, val int) bool {
	if len(arr) == 0 {
		return false
	}

	i := len(arr) / 2
	if val == arr[i] {
		fmt.Printf("Value %d found\n", val)
		return true
	}

	l := BinarySearchRecursive(arr[:i], val)
	r := BinarySearchRecursive(arr[i+1:], val)
	return r || l || false
}
