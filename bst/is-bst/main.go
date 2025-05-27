// Trivial since BST properties imply that this will be the rightmost node
package bst

import (
	"dsa/bst/bst-sorted-list"
)

func IsBST(t *bst.BST) bool {
	return IsBSTNode(t.Root)
}

func IsBSTNode(n *bst.Node) bool {
	var l, r bool
	if n.Left != nil && n.Left.Val >= n.Val {
		return false
	} else {
		l = true
	}
	if n.Right != nil && n.Right.Val < n.Val {
		return false
	} else {
		r = true
	}

	// After we have dealt with the current node and established if it's valid, which happens when l && r == true, we check the subtrees too and combine
	if n.Left != nil {
		l = l && IsBSTNode(n.Left)
	}
	if n.Right != nil {
		r = r && IsBSTNode(n.Right)
	}

	return l && r
}

// The idea here is to traverse down the binary tree and look for the violation of the BST condition
