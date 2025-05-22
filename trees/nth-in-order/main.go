package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func NthInOrderTree(t *cbt.BinTree, count int) {
	fmt.Println("NthInOrder")
	NthInOrder(t.Root, &count)
}

func NthInOrder(n *cbt.Node, count *int) {
	if n.Left != nil {
		NthInOrder(n.Left, count)
	}
	// node has been visited in order, decrement the value that the pointer points to
	*count--

	if *count == 0 {
		fmt.Println(n.Val)
	}
	if n.Right != nil {
		NthInOrder(n.Right, count)
	}
}
