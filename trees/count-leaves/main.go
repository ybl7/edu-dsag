package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func CountLeavesTree(t *cbt.BinTree) {
	fmt.Println("CountLeaves")
	fmt.Println(CountLeaves(t.Root))
}

func CountLeaves(n *cbt.Node) int {
	if n == nil {
		return 0
	}
	if n.Left == nil && n.Right == nil {
		return 1 + CountLeaves(n.Left) + CountLeaves(n.Right)
	} else {
		return CountLeaves(n.Left) + CountLeaves(n.Right)
	}
}
