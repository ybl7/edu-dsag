package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func IsHeapTree(t *cbt.BinTree) bool {
	fmt.Println("IsHeapTree")

	return IsHeapNode(t.Root)
}

func IsHeapNode(n *cbt.Node) bool {
	// For when we hit the leaves and attempt to run IsHeapNode on their children, having no children doesn't violate the minHeap property
	if n == nil {
		return true
	}

	// Set to true initially since if we don't have a left or right child, this doesn't violate the minHeap property
	l, r := true, true
	if n.Left != nil && n.Val < n.Left.Val {
		// Condition for minHeap met, continue to recurse
		l = true
	}
	if n.Left != nil && n.Val >= n.Left.Val {
		return false
	}
	if n.Right != nil && n.Val < n.Right.Val {
		// Condition for minHeap met, continue to recurse
		r = true
	}
	if n.Right != nil && n.Val >= n.Right.Val {
		return false
	}

	// If heap property is satisfied for this node, then check left and right children
	// Otherwise return false
	if l && r {
		return IsHeapNode(n.Left) && IsHeapNode(n.Right)
	} else {
		return false
	}
}

// Given a binary tree, find if it represents a min heap.
// Note: A min heap is a complete binary tree where the data stored in the parent node is smaller than the data in both of the child nodes.
// We'll just recurse down the tree and if we find a node that violates the heap property return false, if no node violates the property we return true
