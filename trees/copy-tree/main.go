package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func CopyTree(t *cbt.BinTree) *cbt.BinTree {
	fmt.Printf("CopyTree")
	t2 := &cbt.BinTree{}
	t2.Root = CopyNode(t.Root)
	return t2
}

func CopyNode(n *cbt.Node) *cbt.Node {
	if n == nil {
		return nil
	}
	m := &cbt.Node{
		Val:   n.Val,
		Left:  CopyNode(n.Left),
		Right: CopyNode(n.Right),
	}
	return m
}

// The idea here is to just iterate down the tree and instantiate a node for every node that we visit
