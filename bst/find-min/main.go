// Trivial since BST properties imply that this will be the leftmost node

package bst

import (
	"dsa/bst/bst-sorted-list"
)

func FindMin(t *bst.BST) int {
	return FindMinNode(t.Root)
}

func FindMinNode(n *bst.Node) int {
	v := n.Val
	curr := n.Left
	for curr != nil {
		// Set the value of v so that we save it, otherwise curr == nil is the condition for exiting this loop and we lose the value of the node
		v = curr.Val
		curr = curr.Left
	}
	return v
}
