package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func CountFullNodesTree(t *cbt.BinTree) {
	fmt.Println("CountFullNodes")
	fmt.Println(CountFullNodes(t.Root))
}

func CountFullNodes(n *cbt.Node) int {
	if n == nil {
		return 0
	}
	if n.Left != nil && n.Right != nil {
		return 1 + CountFullNodes(n.Left) + CountFullNodes(n.Right)
	} else {
		return CountFullNodes(n.Left) + CountFullNodes(n.Right)
	}
}
