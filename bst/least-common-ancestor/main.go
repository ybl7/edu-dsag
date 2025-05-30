package bst

import (
	"dsa/bst/bst-sorted-list"
	"sort"
)

func LCATree(t *bst.BST, m, n int) (int, bool) {
	p1 := GetPath(t.Root, m, &[]int{})
	p2 := GetPath(t.Root, m, &[]int{})
	if p1 == nil || p2 == nil {
		return 0, false
	}

	mP := *p1
	nP := *p2
	// if mP == nil || nP == nil {
	// 	return 0, false
	// }
	sort.Ints(mP)
	sort.Ints(nP)

	var min int
	var found bool
	i, j := 0, 0
	for i < len(mP) && j < len(nP) {
		if mP[i] == nP[j] {
			// First hit is what we want, since we sorted the slice, this will be the LCA value
			if !found {
				found = true
				min = mP[i]
				return min, found
			} else {
				if min > mP[i] {
					min = mP[i]
				}
			}
			i++
			j++
		} else if mP[i] < nP[j] {
			i++
		} else {
			j++
		}
	}
	// We don't need to iterate further since if we don't get a hit and we reach the end of either slice
	// since the slices are sorted, it means that whatever remains in the leftover slice is greater than anything in the exhausted slice

	return 0, found
}

// Find path to node if it exists
func GetPath(n *bst.Node, s int, path *[]int) *[]int {
	if n == nil {
		return nil
	}
	if n.Val == s {
		return path
	}
	// Update the path with the current node
	p := append(*path, n.Val)
	path = &p

	if n.Val < s {
		return GetPath(n.Right, s, path)
	}
	if n.Val > s {
		return GetPath(n.Left, s, path)
	}
	return nil
}

// For any given node n, it will have a chain back to the root, all nodes above it in the chain are its ancestors
// This is a binary tree so each child has at most one parent
// So really if we have m nodes and we want to find their least common ancestor
// First we must check if the tree contains the m nodes
// Then we need to get their ancestor chain
// Then we compare the chains and take the lowest common value

// Now we can use the fact that this is a BST to help us reduce our search space and we can search simulatenously for all m nodes in a single tree traversal
