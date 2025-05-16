package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func PostOrderTree(t *cbt.BinTree) {
	fmt.Println("PostOrder")
	PostOrder(t.Root)
	fmt.Println()
}

func PostOrder(n *cbt.Node) {
	if n.Left != nil {
		PostOrder(n.Left)
	}
	if n.Right != nil {
		PostOrder(n.Right)
	}
	fmt.Print(fmt.Sprint(n.Val) + " ")
}
