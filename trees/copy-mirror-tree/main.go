package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func CopyMirrorTree(t *cbt.BinTree) *cbt.BinTree {
	fmt.Printf("CopyMirrorTree")
	t2 := &cbt.BinTree{}
	t2.Root = CopyMirrorNode(t.Root)
	return t2
}

func CopyMirrorNode(n *cbt.Node) *cbt.Node {
	if n == nil {
		return nil
	}
	m := &cbt.Node{
		Val:   n.Val,
		Left:  CopyMirrorNode(n.Right),
		Right: CopyMirrorNode(n.Left),
	}
	return m
}

// The solution is identical to copy-tree, the only difference is that we switch the Left and Right nodes
