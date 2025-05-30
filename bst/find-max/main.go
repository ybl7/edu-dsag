// Trivial since BST properties imply that this will be the rightmost node
package bst

import (
	"dsa/bst/bst-sorted-list"
)

func FindMax(t *bst.BST) int {
	return FindMaxVal(t.Root)
}

func FindMaxVal(n *bst.Node) int {
	v := n.Val
	curr := n.Right
	for curr != nil {
		// Set the value of v so that we save it, otherwise curr == nil is the condition for exiting this loop and we lose the value of the node
		v = curr.Val
		curr = curr.Right
	}
	return v
}

func FindMaxNode(n *bst.Node) *bst.Node {
	v := n
	curr := n.Right
	for curr != nil {
		// Set the value of v so that we save it, otherwise curr == nil is the condition for exiting this loop and we lose the value of the node
		v = curr
		curr = curr.Right
	}
	return v
}
