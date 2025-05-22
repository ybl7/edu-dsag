package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func NthPostOrderTree(t *cbt.BinTree, count int) {
	fmt.Println("NthPostOrder")
	NthPostOrder(t.Root, &count)
}

func NthPostOrder(n *cbt.Node, count *int) {
	if n.Left != nil {
		NthPostOrder(n.Left, count)
	}
	if n.Right != nil {
		NthPostOrder(n.Right, count)
	}

	// node has been visited last, decrement the value that the pointer points to
	*count--

	if *count == 0 {
		fmt.Println(n.Val)
	}
}
