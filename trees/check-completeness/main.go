package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func CheckCompleteTree(t *cbt.BinTree) bool {
	p := GetCompletePaths(t.Root, 1, []int{})
	fmt.Printf("lengths of paths: %v\n", p)
	for i := 0; i < len(p)-1; i++ {
		// Every single path should differ by no more than 1 from every other path
		if p[i]-p[i+1] < -1 || p[i]-p[i+1] > 1 {
			return false
		}
	}
	return true
}

func GetCompletePaths(n *cbt.Node, curr int, lens []int) []int {
	// Return a list that we can merge later, the list stores the length of paths from root to leaf nodes
	if n.Left == nil && n.Right == nil {
		return []int{curr}
	}

	l, r := []int{}, []int{}
	if n.Left != nil {
		l = GetCompletePaths(n.Left, curr+1, lens)
	}
	if n.Right != nil {
		r = GetCompletePaths(n.Right, curr+1, lens)
	}

	l = append(l, r...)
	return l
}

// Given a binary tree, find if it is a complete tree using the queue.
// The tree is complete if it is filled at all possible levels except the last level. The last level is filled from left to right.

// So to solve this question we need to be able to identify the last level and ensure all levels above it are full.
// But this is equivalent to saying that all paths from the root node to leaves should differ by no more than one node in length.
// So to do this we will traverse the tree down all paths, counting the number of nodes in each path, and when we hit a leaf, we'll save this length.

// This question is a but pointless at this stage in the course since we've only worked with complete binary trees up until now
// But at least the solution correctly prints the lengths of all paths and all the trees in the test (see main.go) return true
