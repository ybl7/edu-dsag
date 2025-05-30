package bst

import (
	"dsa/bst/bst-sorted-list"
	"fmt"
)

func PrintValInRangeTree(t *bst.BST, min, max int) {
	fmt.Printf("Values between %v and %v (inclusive): ", min, max)
	PrintValInRageNode(t.Root, min, max)
	fmt.Println()
}

func PrintValInRageNode(n *bst.Node, min, max int) {
	if n == nil {
		return
	}

	if min <= n.Val && n.Val <= max {
		// In this case we want to do the in order traversal when we actually print something
		PrintValInRageNode(n.Left, min, n.Val)
		fmt.Printf("%v ", n.Val)
		PrintValInRageNode(n.Right, n.Val, max)
	}
	// In these cases we just discard one side of the tree and continue, there is nothing to print
	if n.Val > min && n.Val > max {
		PrintValInRageNode(n.Left, min, max)
	}
	if n.Val < min && n.Val < max {
		PrintValInRageNode(n.Right, min, max)
	}

}

// Print only those nodes of the tree whose value is in the given range.
// Input: Two node values.
// Output: An ordered list of node values

// Suppose we start at a node n, there are a couple of cases
// - min < n < max
// - min < max < n
// - n < min < max

// case: min < n < max
// (1) n is in the list that we return
// (2) there could be some values in both the left and right subtrees that fit our requirements
// (3) left subtree needs to be searched from min to n
// (4) right subtree needs to be searched from n to max

// case: min < max < n
// (1) n is not in the list that we return
// (2) left subtree needs to be searched from min to max

// case: n < min < max
// (1) n is not in the list that we return
// (2) right subtree needs to be searched from min to max

// Now here is the trick, we could store the values in a slice then return this and keep adding to it, then sort it at the end
// But what happens if we traverse the tree in pre-order? Well, since it's a binary search tree, if we traverse in-order we will visit the left child first,
// then the current node, then the right, and since we are guaranteer that n.Left < n < n.Right, we will get implicit ordering in our print statements

// The courses solution is simpler and easier to understand: just traverse the tree in order and print when a node is in the desired range
// func printDataInRange(root *Node, min int, max int) {
// 	if root == nil {
// 		return
// 	}
// 	printDataInRange(root.left, min, max)
// 	if root.value >= min && root.value <= max {
// 		fmt.Print(root.value, " ")
// 	}
// 	printDataInRange(root.right, min, max)
// }
// But I think(?) mine might be quicker for larger trees since we discard whole subtrees that we don't need instead of searching them
