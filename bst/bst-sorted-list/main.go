package bst

import (
	"fmt"
	"strings"
)

// We are operating on sorted arrays
func BSTFromSlc(arr []int) *BST {
	t := &BST{
		Root: &Node{},
	}
	PlaceChildNode(t.Root, arr)
	return t
}

func PlaceChildNode(n *Node, arr []int) {
	if len(arr) == 0 {
		return
	}

	mid := len(arr) / 2
	n.Val = arr[mid]

	// Left child
	if len(arr[:mid]) > 0 {
		n.Left = &Node{}
		PlaceChildNode(n.Left, arr[:mid])
	}

	// Right child
	if len(arr[mid+1:]) > 0 {
		n.Right = &Node{}
		PlaceChildNode(n.Right, arr[mid+1:])
	}
}

type BST struct {
	Root *Node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (b *BST) PrintCompBinTree() {
	PrintNodeRec(b.Root, 0, "")
	fmt.Println()
}

func PrintNodeRec(t *Node, i int, p string) {
	s := strings.Repeat("\t", i)
	if t.Right != nil {
		PrintNodeRec(t.Right, i+1, "r")
	}
	if i == 0 {
		fmt.Println("root", p, t.Val)
	} else {
		fmt.Println(s, p, t.Val)
	}
	if t.Left != nil {
		PrintNodeRec(t.Left, i+1, "l")
	}
}

// A BST or a binary search tree is a binary tree in which nodes are ordered in the following way:

// - The key in the left node is less than the key in its parent node.
// - The key in the right node is greater than the key in its parent node.
// - No duplicate key is allowed.

// The idea here is to use recursion to create the lowest depth binary tree, you could theoretically have 1 or 10 at the root with all other elements being left or right childre.
// That is to say, there are many ways to construct a binary search tree, we'll opt for the way that minimises the depth of the tree - which translates into the condition that a node is the median value of it's children.
