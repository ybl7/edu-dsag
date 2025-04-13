package main

import "fmt"

func main() {
	s := new(StackLinkedList)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val := s.Top()
	fmt.Println("Peek() of a stack is:", val)
	fmt.Print("Stack consist following elements: ")
	for s.IsEmpty() == false {
		val = s.Pop()
		fmt.Print(val, " ")
	}
}

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int // not necessarily needed, but makes it easier to check if the stack is empty and return the length without having to traverse the LL
}

func (s *StackLinkedList) IsEmpty() bool {
	if s.size != 0 {
		return false
	}
	return true
}

func (s *StackLinkedList) Length() int {
	return s.size
}

func (s *StackLinkedList) Print() {
	node := s.head
	fmt.Println(s.head.value)
	for node.next != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (s *StackLinkedList) Push(val int) {
	n := &Node{
		value: val,
		next:  s.head,
	}
	s.head = n
	s.size++
}

func (s *StackLinkedList) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return 0
	}
	v := s.head.value
	s.head = s.head.next
	s.size--
	return v
}

func (s *StackLinkedList) Top() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return 0
	}
	v := s.head.value
	return v
}
