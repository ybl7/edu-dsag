package main

import (
	"fmt"
	bi "stack/bottom-insert"
	stack "stack/stack-over-array"
)

func main() {
	s := stack.StackInt{}
	t := stack.StackInt{}
	s.Push(1)
	s.Push(2)
	s.Push(4)
	s.Push(3)
	s.Print()
	ReverseK(&s, &t, 2)
	s.Print()
	fmt.Println()

	s2 := stack.StackInt{}
	t2 := stack.StackInt{}
	s2.Push(4)
	s2.Push(8)
	s2.Push(2)
	s2.Push(7)
	s2.Push(1)
	s2.Print()
	ReverseK(&s2, &t2, 3)
	s2.Print()
	fmt.Println()

	s3 := stack.StackInt{}
	t3 := stack.StackInt{}
	s3.Push(1)
	s3.Push(2)
	s3.Push(3)
	s3.Push(4)
	s3.Push(5)
	s3.Print()
	ReverseK(&s3, &t3, 4)
	s3.Print()
	fmt.Println()

	s4 := stack.StackInt{}
	t4 := stack.StackInt{}
	s4.Push(6)
	s4.Push(4)
	s4.Push(9)
	s4.Push(3)
	s4.Push(2)
	s4.Push(1)
	s4.Print()
	ReverseK(&s4, &t4, 6)
	s4.Print()
	fmt.Println()
}

func ReverseK(s, t *stack.StackInt, k int) {
	if k == 0 {
		return
	}
	tmp := s.Pop()
	bi.BottomInsert(t, tmp)
	ReverseK(s, t, k-1)
	tmp2 := t.Pop()
	s.Push(tmp2)
}
