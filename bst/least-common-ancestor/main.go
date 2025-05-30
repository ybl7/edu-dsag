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

// Ok my solution is very bad, I missed a piece of logic, for any node a, the nodes m and n can be either
// Both to the left of a, meaning they are both smaller
// Both to the right of a, meaning they are both larger
// Any other arrangement that splits m and n will lead them to be in different subtrees

// If they are in different subtrees then we already have our answer, the current node must be their ancestor
// In the cases where both nodes are in the same subtree, we just keep iterating down the tree until we hit the case where they are in different subtrees
// We can do this check to see if they are in different subtrees without having to traverse the tree at all, BECAUSE IT'S A BST, so we just compare the values
// of a, m, and n. This leads to a much shorter, efficient, elegant solution below:

// func (t *Tree) LcaBST(first int, second int) (int, bool) {
//     return LcaBST(t.root, first, second)
// }

// func LcaBST(curr *Node, first int, second int) (int, bool) {

//     if curr == nil {
//         fmt.Println("NotFoundException")
//         return 0, false
//     }
//     if curr.value > first && curr.value > second {
//         return LcaBST(curr.left, first, second)
//     }
//     if curr.value < first && curr.value < second {
//         return LcaBST(curr.right, first, second)
//     }
//     return curr.value, true
// }
