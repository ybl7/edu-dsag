// Trivial since BST properties imply that this will be the rightmost node
package bst

import (
	"dsa/bst/bst-sorted-list"
	max "dsa/bst/find-max"
)

func DeleteNodeTree(t *bst.BST, n int) *bst.Node {
	return DeleteNode(t.Root, n)
}

func DeleteNode(n *bst.Node, m int) *bst.Node {
	if n == nil {
		return nil
	}
	// Found the node and it has no children, so there is no replacement subtree to deal with
	if n.Val == m && n.Left == nil && n.Right == nil {
		return nil
	}
	// Found the node and it only has right child, therefore we return n.Right in THIS recursive call.
	// In the call that spawns this call, we are actually returning n.Right.Right relative to THAT n. (see line 40), so we have effectively removed the node n.Right.
	if n.Val == m && n.Left == nil {
		return n.Right
	}
	// Found the node and it only has left child, therefore we return n.Left in THIS recursive call.
	// In the call that spawns this call, we are actually returning n.Right.Right relative to THAT n. (see line 42)
	if n.Val == m && n.Right == nil {
		return n.Left
	}
	// If the node has both left and right subtrees, we can't just move a subtree up without breaking BST
	// Find the max od the left subtree, then set the value of the current node to it, then "delete" the node we have just copied into n.Val from the remaining subtree
	// finally we set the left subtree of the node to whatever we get after deleting the node we just promoted from the subtree
	if n.Val == m {
		max := max.FindMaxNode(n.Left)
		n.Val = max.Val
		n.Left = DeleteNode(n.Left, max.Val)
	} else if n.Val < m {
		n.Right = DeleteNode(n.Right, m)
	} else {
		n.Left = DeleteNode(n.Left, m)
	}
	return n
}

// The idea here is to search for the node we need, then if it's there we need to point the parent of the node to the child of the node while maintaining the BST property
// If we delete a node we can assume nothing about it's children, it may have 0, 1, or 2 children
// In the case when there are 0 children, the deletion is trivial
// In the case where there is one children, we just place that single child in place of the node just deleted
// In the case where there are two children, we need to make sure the BST property is maintained, so we see which of the two children is larger/smaller
// If we choose the larger of the two children and move it up, we must make the other node it's left child
// If we choose the smaller of the two children and move it up, me must make the other node it's right child

// In all cases, we will awlays start off from a node that isn't equal to m, unless the root node is equal to m
// What we will return is the subtree with the node that we don't want already deleted
// The logic above handles that in the 4 cases, if the node == m and has no children, the subtree we return is nil
// If the node == m only has left/right subtrees then we return a pointer directly to the left/right subtree and link this to the parent of n
// If the node has both trees, we need to follow the procedure outlined in paragrpah one to select an appropriate node to promote to the value at n, and we return this - but not before ensuring the rest of the tree maintins BST via a recursive call
