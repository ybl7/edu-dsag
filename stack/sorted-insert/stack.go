package main

import "fmt"

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
