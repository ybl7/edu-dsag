package main

import (
	"dsa/bst/bst-sorted-list"
	"fmt"
)

func main() {
	TestBSTFromSlc([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	TestBSTFromSlc([]int{1, 2, 3, 4, 5, 6, 7})
	TestBSTFromSlc([]int{1, 2, 3, 4})
	TestBSTFromSlc([]int{0, 1, 2, 7, 8})
	TestBSTFromSlc([]int{10})
}

func TestBSTFromSlc(arr []int) {
	t := bst.BSTFromSlc(arr)
	fmt.Println("-------------------------------------------------------------------------------")
	t.PrintCompBinTree()
	fmt.Println()
}
