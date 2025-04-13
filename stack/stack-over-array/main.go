package main

import "fmt"

func main() {
	stack := new(StackInt)

	stack.Push(6)
	stack.Push(3)
	stack.Push(2)
	stack.Push(5)

	fmt.Println("Top() of the stack is: ", stack.Top())
	fmt.Print("Stack consist of following elements: ")
	for stack.IsEmpty() == false {
		fmt.Print(stack.Pop(), " ")
	}
}

type StackInt struct {
	s []int
}

func (s *StackInt) IsEmpty() bool {
	if len(s.s) == 0 {
		return true
	}
	return false
}

func (s *StackInt) Length() int {
	l := len(s.s)
	return l
}

func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

func (s *StackInt) Push(value int) {
	s.s = append(s.s, value)
}

func (s *StackInt) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return 0
	}
	l := len(s.s)
	val := s.s[l-1]
	s.s = s.s[:l-1]
	return val
}

func (s *StackInt) Top() int {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return 0
	}
	l := len(s.s)
	val := s.s[l-1]
	return val
}
