package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func DepthFirstTree(t *cbt.BinTree) {
	fmt.Println("DepthFirst")
	DepthFirst(t.Root)
	fmt.Println()
}

func DepthFirst(n *cbt.Node) {
	fmt.Print(fmt.Sprint(n.Val) + " ")
	if n.Left == nil && n.Right == nil {
		return
	}
	if n.Left != nil {
		DepthFirst(n.Left)
	}
	if n.Right != nil {
		DepthFirst(n.Right)
	}
}
