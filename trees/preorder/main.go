package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func PreOrderTree(t *cbt.BinTree) {
	PreOrder(t.Root)
	fmt.Println()
}

func PreOrder(n *cbt.Node) {
	fmt.Print(fmt.Sprint(n.Val) + " ")
	if n.Left != nil {
		PreOrder(n.Left)
	}
	if n.Right != nil {
		PreOrder(n.Right)
	}
}
