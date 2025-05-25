package tree

import (
	"fmt"
	cbt "trees/complete-binary-tree"
)

func InOrderDLLTree(t *cbt.BinTree) {
	fmt.Printf("InOrderDLLTree\n")
	d := &DLL{}
	InOrderDLLNode(t.Root, d)
	d.Print()
}

func InOrderDLLNode(n *cbt.Node, dll *DLL) {
	if n == nil {
		return
	}
	InOrderDLLNode(n.Left, dll)
	m := &NodeDLL{
		Val: n.Val,
	}
	dll.InsertBack(m)
	InOrderDLLNode(n.Right, dll)
}

type DLL struct {
	Head *NodeDLL
	Tail *NodeDLL
}

type NodeDLL struct {
	Next *NodeDLL
	Prev *NodeDLL
	Val  int
}

func (d *DLL) InsertFront(n *NodeDLL) {
	// Empty LL case
	if d.Head == nil {
		d.Head = n
		d.Tail = n
		return
	}
	n.Next = d.Head
	d.Head.Prev = n
	d.Head = n
}

func (d *DLL) InsertBack(n *NodeDLL) {
	// Empty LL case
	if d.Head == nil {
		d.Head = n
		d.Tail = n
		return
	}
	n.Prev = d.Tail
	d.Tail.Next = n
	d.Tail = n
}

func (d *DLL) Print() {
	PrintRec(d.Head)
}

func PrintRec(n *NodeDLL) {
	if n != nil {
		fmt.Print(fmt.Sprint(n.Val) + " ")
		if n.Next != nil {
			// fmt.Print(fmt.Sprint(n.Next.Val) + " ")
			PrintRec(n.Next)
			// if n.Next.Next != nil {
			// 	fmt.Print(fmt.Sprint(n.Next.Next.Val) + " ")
			// 	// PrintRec(n.Next)
			// }
		}
	}
}

// Given a binary tree, create a doubly linked list from the tree such that the elements are in the order of in-order traversal of the tree.
// Recap, in-order traversal means
// - visit n.Left
// - visit n
// - visit n.Right
// So the idea here is just to traverse the tree in-order and when we visit a node, at the time of visiting, add it to the doubly linked list
// Implemented just the insertion methods here, omitted deletion as not necessary for this queston, should a library be used, yes.
