package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func CountNodesTree(t *cbt.BinTree) {
	fmt.Println("CountNodes")
	count := 0
	DepthFirst(t.Root, &count)
	fmt.Println(count)
}

func DepthFirst(n *cbt.Node, count *int) {
	*count++
	if n.Left == nil && n.Right == nil {
		return
	}
	if n.Left != nil {
		DepthFirst(n.Left, count)
	}
	if n.Right != nil {
		DepthFirst(n.Right, count)
	}
}

// We'll use DFS with a counter
