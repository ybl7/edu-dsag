package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func FindMaxTree(t *cbt.BinTree) {
	fmt.Println("FindMax")
	// Going to assume -1 doesn't show up in any of our trees, ideally this should be a zero value that can't show up in the tree, but Go's zero value for ints is 0
	fmt.Println(FindMax(t.Root, -1))
}

func FindMax(n *cbt.Node, currMax int) int {
	if n == nil {
		return currMax
	}
	if n.Val > currMax {
		currMax = n.Val
	}
	l := FindMax(n.Left, currMax)
	r := FindMax(n.Right, currMax)

	if l > r {
		return l
	}
	return r
}

// Find the node with maximum value in a binary tree
