package main

import (
	bi "stack/bottom-insert"
	stack "stack/stack-over-array"
)

func main() {
	s := stack.StackInt{}
	s.Print()
	s.Push(1)
	s.Push(2)
	s.Push(4)
	s.Push(3)
	Reverse(&s)
	s.Print()

	s2 := stack.StackInt{}
	s2.Push(4)
	s2.Push(8)
	s2.Push(2)
	s2.Push(7)
	s2.Push(1)
	Reverse(&s2)
	s2.Print()

	s3 := stack.StackInt{}
	s3.Push(6)
	s3.Push(5)
	s3.Push(8)
	s3.Push(3)
	s3.Push(1)
	Reverse(&s3)
	s3.Print()

	s4 := stack.StackInt{}
	s4.Push(0)
	s4.Push(3)
	s4.Push(2)
	s4.Push(4)
	s4.Push(5)
	Reverse(&s4)
	s4.Print()
}

func Reverse(s *stack.StackInt) {
	if s.Length() == 0 {
		return
	}
	tmp := s.Pop()
	Reverse(s)
	bi.BottomInsert(s, tmp)
}
