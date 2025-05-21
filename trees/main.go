package main

import (
	bfs "trees/breadth-first"
	cbt "trees/complete-binary-tree"
	dfs "trees/depth-first"
	in "trees/inorder"
	pst "trees/postorder"
	pre "trees/preorder"
)

func main() {
	TestBinTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	TestBinTree([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})
	TestBinTree([]int{10, 100, 1000, 10000, 100000, 1000000, 10000000})
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
}
