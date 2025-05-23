package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func SearchTree(t *cbt.BinTree, n int) {
	fmt.Printf("SearchTree for n=%v\n", n)
	fmt.Println(SearchNode(t.Root, n))
}

func SearchNode(n *cbt.Node, m int) bool {
	if n == nil {
		return false
	}
	if n.Val == m {
		return true
	}
	if SearchNode(n.Left, m) || SearchNode(n.Right, m) {
		return true
	}
	return false
}
