package bst

import (
	"dsa/bst/bst-sorted-list"
	"fmt"
)

func SearchBST(t *bst.BST, m int) {
	b := SearchNode(t.Root, m)
	if b {
		fmt.Printf("value %v found\n", m)
	} else {
		fmt.Printf("value %v not found\n", m)
	}
}

func SearchNode(n *bst.Node, m int) bool {
	if n == nil {
		return false
	}
	if m < n.Val {
		return SearchNode(n.Left, m)
	}
	if m > n.Val {
		return SearchNode(n.Right, m)
	}
	if m == n.Val {
		return true
	}
	return false
}

// Return true if a node of value n is found in a BST
// Since thie is a BST searching will be easy, we just go down the tree and compare the value of of the node toe the value we are inserting
// The iterative solution from the course is cleaner
// func (t *Tree) Find(value int) bool {
//     var curr *Node = t.root
//     for curr != nil {
//         if curr.value == value {
//             return true
//         } else if curr.value > value {
//             curr = curr.left
//         } else {
//             curr = curr.right
//         }
//     }
//     return false
// }
