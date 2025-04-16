package main

import (
	stack "stack/stack-over-array"
)

func main() {
	s := stack.StackInt{}
	s.Push(1)
	s.Push(2)
	s.Push(4)
	s.Push(3)
	Sort(&s)
	s.Print()

	s2 := stack.StackInt{}
	s2.Push(4)
	s2.Push(8)
	s2.Push(2)
	s2.Push(7)
	s2.Push(1)
	Sort(&s2)
	s2.Print()

	s3 := stack.StackInt{}
	s3.Push(6)
	s3.Push(5)
	s3.Push(8)
	s3.Push(3)
	s3.Push(1)
	Sort(&s3)
	s3.Print()

	s4 := stack.StackInt{}
	s4.Push(0)
	s4.Push(3)
	s4.Push(2)
	s4.Push(4)
	s4.Push(5)
	Sort(&s4)
	s4.Print()
}

func Sort(s *stack.StackInt) {
	if s.Length() <= 1 {
		return
	}
	tmp := s.Pop()
	Sort(s)
	SortedInsertRec(s, tmp)
}

// Method two is so much cleaner and easier to see
func SortedInsertRec(s *stack.StackInt, n int) {
	var temp int
	if s.Length() == 0 || n > s.Top() {
		s.Push(n)
	} else {
		temp = s.Pop()
		SortedInsertRec(s, n)
		s.Push(temp)
	}
}

// Once we have the SortedInsert subroutine, this becomes a case of comparing the top element with those below it.
// We pop off the top element and use it in a call to SortedInsert. We do this recursively down the entire stack until we have re-inserted all elements.
// The recursive part is important because we want to guarantee that the stack we are re-inserting into is already sorted
