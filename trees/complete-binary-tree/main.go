package tree

import (
	"fmt"
	"strings"
)

type BinTree struct {
	root *Node
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func CompBinTree(arr []int) *BinTree {
	b := &BinTree{
		root: GenTree(arr, 0),
	}
	return b
}

func GenTree(arr []int, i int) *Node {
	n := &Node{}
	n.val = arr[i]
	l := 2*i + 1
	r := 2*i + 2

	if l < len(arr) {
		n.left = GenTree(arr, l)
	}
	if r < len(arr) {
		n.right = GenTree(arr, r)
	}
	return n
}

func PrintCompBinTree(b *BinTree) {
	PrintNodeRec(b.root, 0)
	fmt.Println()
}

func PrintNodeRec(t *Node, i int) {
	s := strings.Repeat("\t", i)
	if t.left != nil {
		PrintNodeRec(t.left, i+1)
	}
	fmt.Println(s, t.val)
	if t.right != nil {
		PrintNodeRec(t.right, i+1)
	}
}

// All levels are filled in order, starting from index 0
// For any given node i left child is given by 2 * i + 1 and right child is given by 2 * i + 2
