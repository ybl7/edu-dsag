package trees

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func BreadthFirstTree(t *cbt.BinTree) {
	fmt.Println("BreadthFirst")
	r := []*cbt.Node{
		t.Root,
	}
	BreadthFirst(r, 0)
	fmt.Println()
}

func BreadthFirst(n []*cbt.Node, level int) {
	if len(n) == 0 {
		return
	}
	m := []*cbt.Node{}

	// fmt.Print("level: ", level, " | ")
	for _, e := range n {
		fmt.Print(fmt.Sprint(e.Val) + " ")
		if e.Left != nil {
			m = append(m, e.Left)
		}
		if e.Right != nil {
			m = append(m, e.Right)
		}
	}
	// fmt.Print("\n")

	BreadthFirst(m, level+1)
}

func BreadthFirstLevelOrderTree(t *cbt.BinTree) {
	fmt.Println("BreadthFirstLeverOrder")
	r := []*cbt.Node{
		t.Root,
	}
	BreadthFirstLevelOrder(r, 0)
	fmt.Println()
}

func BreadthFirstLevelOrder(n []*cbt.Node, level int) {
	if len(n) == 0 {
		return
	}
	m := []*cbt.Node{}

	for _, e := range n {
		fmt.Print(fmt.Sprint(e.Val) + " ")
		if e.Left != nil {
			m = append(m, e.Left)
		}
		if e.Right != nil {
			m = append(m, e.Right)
		}
	}
	fmt.Print("; ")

	BreadthFirstLevelOrder(m, level+1)
}

// This is the same as the question Print Level Order by Line Using Two Queues, and even the saem as Print Level Order by Line Using a Queue, if you uncomment the print statements you get it
