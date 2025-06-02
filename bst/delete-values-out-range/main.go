package bst

import (
	"dsa/bst/bst-sorted-list"
)

func DeleteRangeTree(t *bst.BST, min, max int) *bst.Node {
	return DeleteRangeNode(t.Root, min, max)
}

func DeleteRangeNode(n *bst.Node, min, max int) *bst.Node {
	if n == nil {
		return nil
	}
	n.Left = DeleteRangeNode(n.Left, min, max)
	n.Right = DeleteRangeNode(n.Right, min, max)
	if n.Val < min {
		return n.Right
	}
	if n.Val > max {
		return n.Left
	}
	return n
}

// The idea here is to once again make use of the BST structure, suppose we have some interval [min, max]
// At some node down the tree, all nodes to the left can be discarded since they are smaller than min, the converse applies with max
// So we can iterate down the tree, and when we discover a left child smaller than min, we discard the rest of the left tree and return it's right side
// because we know that the right must be larger than the node, so the right side could contain values in our range that we don't want to discard
// The converse applies for the right child and max. In the case where the node's value doesn't fall outside of the range, we just return it.
// We'll operate on the current node's children (via a recursive call) and return the current node
// Now suppose that our root node violates the range, we want to replace it with either it's left or right child, but we want to have operated on the rest of
// the tree first, otherwise if we immediately return at the root level, we would not have made a recursive call it, therefore we opt to put the recursive
// calls to delete down the tree BEFORE we do the final assignment to the node
