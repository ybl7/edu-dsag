package main

import (
	"fmt"
	bfs "trees/breadth-first"
	spl "trees/breadth-first-spiral"
	cbt "trees/complete-binary-tree"
	cfl "trees/count-full-nodes"
	ctl "trees/count-leaves"
	nel "trees/count-nodes"
	dfs "trees/depth-first"
	in "trees/inorder"
	npi "trees/nth-in-order"
	npt "trees/nth-post-order"
	npo "trees/nth-pre-order"
	pst "trees/postorder"
	pre "trees/preorder"
	pap "trees/print-all-paths"
	sum "trees/sum-nodes"
)

func main() {
	// TestBinTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// TestBinTree([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	// TestBinTree([]int{10, 100, 1000, 10000, 100000, 1000000, 10000000})
	// TestBinTree([]int{0, 7, 2, 8, 1})

	TestGetNth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7)
	TestGetNth([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	TestGetNth([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	TestGetNth([]int{1, 2, 3, 4}, 2)
	TestGetNth([]int{0, 7, 1, 8, 2}, 3)
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
	cbt.PrintCompBinTree(t)

	npo.NthPreOrderTree(t, n)
	npt.NthPostOrderTree(t, n)
	npi.NthInOrderTree(t, n)
	pap.PrintAllPaths(t)
	nel.CountNodesTree(t)
	sum.SumNodesTree(t)
	ctl.CountLeavesTree(t)
	cfl.CountFullNodesTree(t)
	fmt.Println()
}
