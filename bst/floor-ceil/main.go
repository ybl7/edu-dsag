package bst

import (
	"dsa/bst/bst-sorted-list"
)

func FloorTree(t *bst.BST, v int) int {
	return FloorNode(t.Root, v)
}

func FloorNode(n *bst.Node, v int) int {
	var f int
	var found bool
	for n != nil {
		if n.Val >= v {
			n = n.Left
		} else {
			if !found {
				found = true
			}
			// Left subtree is smaller than current node, only worse values in it, right subtree may contain a better floor
			f = n.Val
			n = n.Right
		}
	}

	if !found {
		return -1
	}
	return f
}

func CeilTree(t *bst.BST, v int) int {
	return CeilNode(t.Root, v)
}

func CeilNode(n *bst.Node, v int) int {
	var f int
	var found bool

	for n != nil {
		if n.Val <= v {
			// This nose and left subtree are all smaller than v so can't possibly be ceil, we only need to look at the right subtree
			n = n.Right
		} else {
			if !found {
				found = true
			}
			// The current node is a viable candidate, everything to it's right will be larger than it, so they can't be a better ceil.
			// The left subtree may contain a smaller value that is a better ceil
			f = n.Val
			n = n.Left
		}
	}

	if !found {
		return -1
	}
	return f
}

// For a given value n, we want to find the smallest interval that constrains it
// The floor value is the largest number that is smaller than n, the ceiling is the smallest number that is larger than n
// The simplest way to think about is in terms of left to right, we know that a BST orders values from left to right (ignoring levels)
// So what we are really looking for is the closest nodes to the left/right of the node in question regardless of level
// The problem is the only way to know this is to traverse the tree

// You might think that we only need to search the left and right children of a particular node, but this isn't the case
// For example take the BST
//								6
//				4								8
// 		2				5				7				9
// 1		3													10

// Suppose we are at some node n, with some target value m for which we are trying to find the floor and ceil for
// If n is a floor candidate, that is n < m, then we can immediately discard n.Left because everything in this subtree is smaller than n
// and the floor value is the largest value that is smaller than m. If it is not a floor candidate, that is n >= m, then we must search
// the left subtree but can safely discard the right subtree. If we find a viable candidate, we can set our floor to that knowing that
// based on the logic above we will always get a more optimal value as we traverse down the BST.
// Since we discard the subtrees that would produce worse values, only better values are left to be disocered if they exist

// Conversly if n is a ceil candidate, n > m, then we can discard n.right, since everything in n.Right will be larger hence not the ceil
// However once we discard the left/right subtree in the two cases, the remaining subtree may still yet contain a better candidate.
