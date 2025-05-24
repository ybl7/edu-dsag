// The root nodes t.root and t2.root of two binary trees will be our input. We have to find if the trees are identical using the isEqual() function
// This means that we need to keep track of the position of the node in addition to all of it's values. In practise this works by a "proof by contradtion"
// The very first case we find that doesn't satify the condition; a single node with left child l and righ child r is identical to another node if:
// n1.Val == n2.Val
// Taking this condition recursively down the tree for left and right childre, we can ensure entire subtrees are identical

package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func IsEqualTree(t1, t2 *cbt.BinTree) {
	fmt.Println("IsEqualTree")
	fmt.Println(IsEqual(t1.Root, t2.Root))
}

func IsEqual(n1 *cbt.Node, n2 *cbt.Node) bool {
	if n1.Val != n2.Val {
		return false
	}
	var l, r bool
	// Check the case where children exist, are they equal?
	if n1.Left != nil && n2.Left != nil {
		l = IsEqual(n1.Left, n2.Left)
	}
	if n1.Right != nil && n2.Right != nil {
		r = IsEqual(n1.Right, n2.Right)
	}
	// Check the case where children don't exist, do both nodes lack the same children?
	if n1.Left == nil && n2.Left == nil {
		l = true
	}
	if n1.Right == nil && n2.Right == nil {
		r = true
	}

	// Only equal if both left and right nodes are equal
	return l && r
}
