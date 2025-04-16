package BottomInsert

import (
	stack "stack/stack-over-array"
)

func main() {
	s := stack.StackInt{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	BottomInsert(&s, 5)
	BottomInsert(&s, 8)
	BottomInsert(&s, 0)
	BottomInsert(&s, 6)
	s.Print()
}

func BottomInsert(s *stack.StackInt, n int) {
	if s.Length() == 0 {
		s.Push(n)
		return
	}
	tmp := s.Pop()
	BottomInsert(s, n)
	s.Push(tmp)
}
