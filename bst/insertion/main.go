package bst

import (
	"dsa/bst/bst-sorted-list"
)

func AddNodeTree(t *bst.BST, m int) {
	if t.Root == nil {
		t.Root = &bst.Node{
			Val: m,
		}
	} else {
		AddNode(t.Root, m)
	}
}

func AddNode(n *bst.Node, m int) {
	if m < n.Val && n.Left != nil {
		AddNode(n.Left, m)
	}
	if m < n.Val && n.Left == nil {
		n.Left = &bst.Node{
			Val: m,
		}
	}
	if m >= n.Val && n.Right != nil {
		AddNode(n.Right, m)
	}
	if m >= n.Val && n.Right == nil {
		n.Right = &bst.Node{
			Val: m,
		}
	}
}

// Insert n values into a BST while ensuring BST properties are preserved
// The idea here is to traverse the current tree and insert the node where appropriate by comparing the value we are inserting with the current node
// then we recurse down the tree as necessary until we find the correct place for the value
