package tree

import (
	"fmt"
	"strings"
)

type BinTree struct {
	Root *Node
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func CompBinTree(arr []int) *BinTree {
	b := &BinTree{
		Root: GenTree(arr, 0),
	}
	return b
}

func GenTree(arr []int, i int) *Node {
	n := &Node{}
	n.Val = arr[i]
	l := 2*i + 1
	r := 2*i + 2

	if l < len(arr) {
		n.Left = GenTree(arr, l)
	}
	if r < len(arr) {
		n.Right = GenTree(arr, r)
	}
	return n
}

func PrintCompBinTree(b *BinTree) {
	PrintNodeRec(b.Root, 0, "")
	fmt.Println()
}

func PrintNodeRec(t *Node, i int, p string) {
	s := strings.Repeat("\t", i)
	if t.Right != nil {
		PrintNodeRec(t.Right, i+1, "r")
	}
	fmt.Println(s, p, t.Val)
	if t.Left != nil {
		PrintNodeRec(t.Left, i+1, "l")
	}
}

// All levels are filled in order, starting from index 0
// For any given node i left child is given by 2 * i + 1 and right child is given by 2 * i + 2
