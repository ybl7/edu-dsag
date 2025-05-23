package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func SumNodesTree(t *cbt.BinTree) {
	fmt.Println("SumElements")
	fmt.Println(SumNodes(t.Root))
}

func SumNodes(n *cbt.Node) int {
	if n == nil {
		return 0
	}

	return (n.Val + SumNodes(n.Left) + SumNodes(n.Right))
}
