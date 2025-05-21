package trees

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func BreadthFirstSpiralTree(t *cbt.BinTree) {
	fmt.Println("BreadthFirstSpiral")
	r := []*cbt.Node{
		t.Root,
	}
	BreadthFirstSpiral(r, 0)
	fmt.Println()
}

func BreadthFirstSpiral(n []*cbt.Node, end int) {
	if len(n) == 0 {
		return
	}
	m := []*cbt.Node{}

	if end%2 != 0 {
		for i := len(n) - 1; i >= 0; i-- {
			fmt.Print(fmt.Sprint(n[i].Val) + " ")
		}
		for _, e := range n {
			if e.Left != nil {
				m = append(m, e.Left)
			}
			if e.Right != nil {
				m = append(m, e.Right)
			}
		}
	} else {
		for _, e := range n {
			fmt.Print(fmt.Sprint(e.Val) + " ")
			if e.Left != nil {
				m = append(m, e.Left)
			}
			if e.Right != nil {
				m = append(m, e.Right)
			}
		}
	}
	fmt.Print("; ")

	BreadthFirstSpiral(m, end+1)
}
