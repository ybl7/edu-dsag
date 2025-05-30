package main

import (
	"dsa/bst/bst-sorted-list"
	del "dsa/bst/delete-node"
	max "dsa/bst/find-max"
	min "dsa/bst/find-min"
	sch "dsa/bst/find-node"
	add "dsa/bst/insertion"
	is "dsa/bst/is-bst"
	lca "dsa/bst/least-common-ancestor"
	pvr "dsa/bst/print-values-in-range"
	"fmt"
	"slices"
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

	// TestSearch([]int{6, 4, 9, 3, 7, 8, 2, 5, 1, 10}, 10)
	// TestSearch([]int{4, 5, 1, 2, 6, 7, 3}, 10)
	// TestSearch([]int{0, 3, 4, 2, 1}, 3)
	// TestSearch([]int{7, 1, 0, 7, 8, 2}, 7)
	// TestSearch([]int{10}, 10)
	// TestSearch([]int{6, 4, 2, 5, 1, 3, 8, 7, 9, 10}, 11)

	// TestDelete([]int{6, 4, 9, 3, 7, 8, 2, 5, 1, 10}, 10)
	// TestDelete([]int{4, 5, 1, 2, 6, 7, 3}, 10)
	// TestDelete([]int{0, 3, 4, 2, 1}, 3)
	// TestDelete([]int{7, 1, 0, 7, 8, 2}, 7)
	// TestDelete([]int{10}, 10)
	// TestDelete([]int{6, 4, 2, 5, 1, 3, 8, 7, 9, 10}, 11)

	// TestLCA([]int{6, 4, 9, 3, 7, 8, 2, 5, 1, 10}, 10, 7)
	// TestLCA([]int{4, 5, 1, 2, 6, 7, 3}, 7, 3)
	// TestLCA([]int{0, 3, 4, 2, 1}, 3, 4)
	// TestLCA([]int{7, 1, 0, 7, 8, 2}, 7, 2)
	// TestLCA([]int{10}, 10, 0)
	// TestLCA([]int{6, 4, 2, 5, 1, 3, 8, 7, 9, 10}, 2, 4)

	TestPVR([]int{6, 4, 9, 3, 7, 8, 2, 5, 1, 10}, 4, 9)
	TestPVR([]int{4, 5, 1, 2, 6, 7, 3}, 3, 7)
	TestPVR([]int{0, 3, 4, 2, 1}, 0, 4)
	TestPVR([]int{7, 1, 0, 7, 8, 2}, 2, 7)
	TestPVR([]int{10}, 0, 10)
	TestPVR([]int{6, 4, 2, 5, 1, 3, 8, 7, 9, 10}, 2, 4)
}

func TestBSTFromSlc(arr []int) {
	// Our BSTFromSlc function assumes a sorted array
	slices.Sort(arr)
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
	// Use the copy method so that we don't use the same underlying array
	arr2 := make([]int, len(arr)) // grr, you must use make here, using a slice literal like []int{} doesn't work
	copy(arr2, arr)
	// Our BSTFromSlc function assumes a sorted array
	slices.Sort(arr)
	t := bst.BSTFromSlc(arr)
	// Alternatively we can construct the BST by adding iteratively to the tree
	t2 := &bst.BST{}
	for _, e := range arr2 {
		add.AddNodeTree(t2, e)
	}
	// The reason for showing both approaches, and printing both trees, is to demonstrate that the min/max values are always the leftmost/rightmost nodes respectively
	fmt.Println("-------------------------------------------------------------------------------")
	t.PrintCompBinTree()
	fmt.Println()
	t2.PrintCompBinTree()
	fmt.Println()
	sch.SearchBST(t, n)

	fmt.Printf("Min: %v\n", min.FindMin(t))
	fmt.Printf("Max: %v\n", max.FindMax(t))
	fmt.Printf("IsBST: %v\n", is.IsBST(t)) // true
	t3 := t
	var v int
	// Now let's break our BST
	if t3.Root.Left != nil {
		v = t3.Root.Left.Val
		t3.Root.Left.Val = 100
	}
	fmt.Println()
	t3.PrintCompBinTree()
	fmt.Println()
	fmt.Printf("IsBST: %v\n", is.IsBST(t3)) // false

	// Now let's make it a BST again
	if t3.Root.Left != nil {
		t3.Root.Left.Val = v
	}
	t3.PrintCompBinTree()
	fmt.Println()
	fmt.Printf("IsBST: %v\n", is.IsBST(t3)) // true

	// Now let's break our BST again
	if t3.Root.Right != nil {
		v = t3.Root.Right.Val
		t3.Root.Right.Val = -100
	}
	fmt.Println()
	t3.PrintCompBinTree()
	fmt.Println()
	fmt.Printf("IsBST: %v\n", is.IsBST(t3)) // false

	// Now let's make it a BST again
	if t3.Root.Right != nil {
		t3.Root.Right.Val = v
	}
	t3.PrintCompBinTree()
	fmt.Println()
	fmt.Printf("IsBST: %v\n", is.IsBST(t3)) // true
}

func TestDelete(arr []int, m int) {
	slices.Sort(arr)
	t := bst.BSTFromSlc(arr)
	fmt.Println("-------------------------------------------------------------------------------")
	t.PrintCompBinTree()
	fmt.Println()
	t.Root = del.DeleteNodeTree(t, m)
	fmt.Println("Delete node: ", m)
	if t.Root == nil {
		fmt.Println("root: nil")
	} else {
		t.PrintCompBinTree()
		fmt.Println()
	}
}

func TestLCA(arr []int, m, n int) {
	slices.Sort(arr)
	t := bst.BSTFromSlc(arr)
	fmt.Println("-------------------------------------------------------------------------------")
	t.PrintCompBinTree()
	fmt.Println()
	a, b := lca.LCATree(t, m, n)
	if b {
		fmt.Printf("LCM found for %v and %v: %v\n", m, n, a)
	} else {
		fmt.Printf("LCM not found for %v and %v\n", m, n)
	}
}

func TestPVR(arr []int, m, n int) {
	slices.Sort(arr)
	t := bst.BSTFromSlc(arr)
	fmt.Println("-------------------------------------------------------------------------------")
	t.PrintCompBinTree()
	fmt.Println()
	pvr.PrintValInRangeTree(t, m, n)
}
