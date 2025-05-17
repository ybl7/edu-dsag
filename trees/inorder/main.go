package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func InOrderTree(t *cbt.BinTree) {
	fmt.Println("InOrder")
	InOrder(t.Root)
	fmt.Println()
}

func InOrder(n *cbt.Node) {
	if n.Left != nil {
		InOrder(n.Left)
	}
	fmt.Print(fmt.Sprint(n.Val) + " ")
	if n.Right != nil {
		InOrder(n.Right)
	}
}
