package main

import (
	"fmt"
	bfs "trees/breadth-first"
	spl "trees/breadth-first-spiral"
	cbt "trees/complete-binary-tree"

	// cpm "trees/copy-mirror-tree"

	// cp "trees/copy-tree"
	// eq "trees/identical-trees"
	cmp "trees/check-completeness"
	cfl "trees/count-full-nodes"
	ctl "trees/count-leaves"
	nel "trees/count-nodes"
	dfs "trees/depth-first"
	dpt "trees/find-depth"
	max "trees/find-max"
	in "trees/inorder"
	mh "trees/is-min-heap"
	npi "trees/nth-in-order"
	npt "trees/nth-post-order"
	npo "trees/nth-pre-order"
	pst "trees/postorder"
	pre "trees/preorder"
	pap "trees/print-all-paths"
	sch "trees/search-value"
	sum "trees/sum-nodes"
	ttl "trees/tree-to-list"
)

func main() {
	// TestBinTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// TestBinTree([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	// TestBinTree([]int{10, 100, 1000, 10000, 100000, 1000000, 10000000})
	// TestBinTree([]int{0, 7, 2, 8, 1})

	// TestGetNth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7)
	// TestGetNth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	// TestGetNth([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	// TestGetNth([]int{1, 2, 3, 4}, 2)
	// TestGetNth([]int{0, 7, 1, 8, 2}, 3)
	// TestGetNth([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, 3)
	// TestGetNth([]int{10}, 0)

	// TestTreeComp([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) // true
	// TestTreeComp([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7})           // false
	// TestTreeComp([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2})    // false
	// TestTreeComp([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 2, 3, 4, 5, 6, 7})                     // true

	// IsMinHeap([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// IsMinHeap([]int{1, 2, 3, 4, 5, 6, 7})
	// IsMinHeap([]int{1, 2, 3, 4})
	// IsMinHeap([]int{0, 7, 1, 8, 2})
	// IsMinHeap([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	// IsMinHeap([]int{10})

	TestTreeToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	TestTreeToList([]int{1, 2, 3, 4, 5, 6, 7})
	TestTreeToList([]int{1, 2, 3, 4})
	TestTreeToList([]int{0, 7, 1, 8, 2})
	TestTreeToList([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	TestTreeToList([]int{10})
}

func TestBinTree(arr []int) {
	t := cbt.CompBinTree(arr)
	cbt.PrintCompBinTree(t)

	pre.PreOrderTree(t)
	pst.PostOrderTree(t)
	in.InOrderTree(t)
	bfs.BreadthFirstTree(t)
	bfs.BreadthFirstLevelOrderTree(t)
	dfs.DepthFirstTree(t)
	spl.BreadthFirstSpiralTree(t)
}

func TestGetNth(arr []int, n int) {
	t := cbt.CompBinTree(arr)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()
	cbt.PrintCompBinTree(t)

	npo.NthPreOrderTree(t, n)
	npt.NthPostOrderTree(t, n)
	npi.NthInOrderTree(t, n)
	pap.PrintAllPaths(t)
	nel.CountNodesTree(t)
	sum.SumNodesTree(t)
	ctl.CountLeavesTree(t)
	cfl.CountFullNodesTree(t)
	sch.SearchTree(t, n)
	max.FindMaxTree(t)
	dpt.FindDepthTree(t)
	fmt.Println()
}

func TestTreeComp(arr1, arr2 []int) {
	t1 := cbt.CompBinTree(arr1)
	// t2 := cbt.CompBinTree(arr2)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()
	// cbt.PrintCompBinTree(t1)
	// cbt.PrintCompBinTree(t2)
	// eq.IsEqualTree(t1, t2)

	// t3 := cp.CopyTree(t1)
	// eq.IsEqualTree(t1, t3)
	// fmt.Println("t1")
	// cbt.PrintCompBinTree(t1)
	// fmt.Println("t3")
	// cbt.PrintCompBinTree(t3)

	// t3 := cpm.CopyMirrorTree(t1)
	// fmt.Println("t1")
	// cbt.PrintCompBinTree(t1)
	// fmt.Println("t3")
	// cbt.PrintCompBinTree(t3)

	cbt.PrintCompBinTree(t1)
	fmt.Println("CheckCompleteness")
	fmt.Println(cmp.CheckCompleteTree(t1))
}

func IsMinHeap(arr []int) {
	t := cbt.CompBinTree(arr)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()
	cbt.PrintCompBinTree(t)
	fmt.Println(mh.IsHeapTree(t))
}

func TestTreeToList(arr []int) {
	t := cbt.CompBinTree(arr)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()
	cbt.PrintCompBinTree(t)
	ttl.InOrderDLLTree(t)
	fmt.Println()
}
