package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func FindDepthTree(t *cbt.BinTree) {
	fmt.Println("FindDepth")
	// Start count from 0 since the root node has depth 1
	fmt.Println(FindDepth(t.Root, 1))
}

func FindDepth(n *cbt.Node, currDepth int) int {
	if n == nil {
		return currDepth
	}
	if n.Left == nil && n.Right == nil {
		return currDepth
	}
	l := 1 + FindDepth(n.Left, currDepth)
	r := 1 + FindDepth(n.Right, currDepth)

	if l > r {
		return l
	}
	return r
}

// Given a binary tree, find its depth. The depth of the root is 0. The depth of the tree is the number of edges in the longest branch from the root to the leaf
// We don't know all the paths from the root to the leaves of the tree from the outset, so we will need to visit every path. The only time we know that a path cannot possibly be the longest
// is if it ends while there is still an active path somewhere else that is longer. So we'll just traverse the tree while keeping track of the longest path yet.
