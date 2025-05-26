package main

import (
	"dsa/bst/bst-sorted-list"
	sch "dsa/bst/find-node"
	add "dsa/bst/insertion"
	"fmt"
)

func main() {
	// TestBSTFromSlc([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// TestBSTFromSlc([]int{1, 2, 3, 4, 5, 6, 7})
	// TestBSTFromSlc([]int{1, 2, 3, 4})
	// TestBSTFromSlc([]int{0, 1, 2, 7, 8})
	// TestBSTFromSlc([]int{10})

	// TestAddNode([]int{6, 4, 9, 3, 7, 8, 2, 5, 1, 10})
	// TestAddNode([]int{4, 5, 1, 2, 6, 7, 3})
	// TestAddNode([]int{0, 3, 4, 2, 1})
	// TestAddNode([]int{7, 1, 0, 7, 8, 2})
	// TestAddNode([]int{10})
	// TestAddNode([]int{6, 4, 2, 5, 1, 3, 8, 7, 9, 10})

	TestSearch([]int{6, 4, 9, 3, 7, 8, 2, 5, 1, 10}, 10)
	TestSearch([]int{4, 5, 1, 2, 6, 7, 3}, 10)
	TestSearch([]int{0, 3, 4, 2, 1}, 3)
	TestSearch([]int{7, 1, 0, 7, 8, 2}, 7)
	TestSearch([]int{10}, 10)
	TestSearch([]int{6, 4, 2, 5, 1, 3, 8, 7, 9, 10}, 11)
}

func TestBSTFromSlc(arr []int) {
	t := bst.BSTFromSlc(arr)
	fmt.Println("-------------------------------------------------------------------------------")
	t.PrintCompBinTree()
	fmt.Println()
}

func TestAddNode(arr []int) {
	t := &bst.BST{}
	for _, e := range arr {
		add.AddNodeTree(t, e)
	}
	fmt.Println("-------------------------------------------------------------------------------")
	t.PrintCompBinTree()
	fmt.Println()
}

func TestSearch(arr []int, n int) {
	t := bst.BSTFromSlc(arr)
	fmt.Println("-------------------------------------------------------------------------------")
	t.PrintCompBinTree()
	fmt.Println()
	sch.SearchBST(t, n)
}
