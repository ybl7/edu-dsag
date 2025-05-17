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
