package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func NthPreOrderTree(t *cbt.BinTree, count *int) {
	fmt.Println("NthPreOrder")
	NthPreOrder(t.Root, count)
}

func NthPreOrder(n *cbt.Node, count *int) {
	// node has been visited, decrement the value that the pointer points to
	*count--

	if *count == 0 {
		fmt.Println(n.Val)
	}

	if n.Left != nil {
		NthPreOrder(n.Left, count)
	}
	if n.Right != nil {
		NthPreOrder(n.Right, count)
	}
}

// The idea here is to traverse the tree in pre-order while keeping a global counter and decrementing it every time we visit a node#
// We'll use a pointer for the count variable so that we can modify it globally regardless of recursion level
