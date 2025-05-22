package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"

	"github.com/idsulik/go-collections/stack"
)

func PrintAllPaths(t *cbt.BinTree) {
	fmt.Println("PrintAllPaths")
	var s []int
	PrintCurrPath(t.Root, s)
	fmt.Println()
}

// We opt to copy the full stack, we don't want to accidentally mix up paths
func PrintCurrPath(n *cbt.Node, s []int) {
	s = append(s, n.Val)

	if n.Left == nil && n.Right == nil {
		// Use a stack to reverse order of printing as required by question
		t := stack.New[int](len(s))
		for _, e := range s {
			t.Push(e)
		}
		for !t.IsEmpty() {
			v, _ := t.Pop()
			fmt.Print(fmt.Sprint(v) + " ")
		}
		fmt.Print("; ")
	}
	if n.Left != nil {
		PrintCurrPath(n.Left, s)
	}
	if n.Right != nil {
		PrintCurrPath(n.Right, s)
	}
}

// The idea here is to keep track of the current path and pass it to the recursive call, then the condition to print the path is when we encounter a node with no left or right children
// The method of traversing the tree, pre/in/post order does not matter
