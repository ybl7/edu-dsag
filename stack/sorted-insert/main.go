package main

import (
	"fmt"
)

func main() {
	stack := new(StackInt)

	SortedInsert(stack, 6)
	SortedInsert(stack, 3)
	SortedInsert(stack, 2)
	SortedInsert(stack, 5)

	fmt.Println("Top() of the stack is: ", stack.Top())
	fmt.Print("Stack consist of following elements: ")
	for !stack.IsEmpty() {
		fmt.Print(stack.Pop(), " ")
	}
	fmt.Println()

	st1 := StackInt{
		s: []int{1, 2, 4, 9},
	}
	SortedInsert(&st1, 3)
	SortedInsert(&st1, 7)
	SortedInsert(&st1, 8)
	SortedInsert(&st1, 10)
	SortedInsert(&st1, 11)

	fmt.Print("st1 consist of following elements: ")
	for !st1.IsEmpty() {
		fmt.Print(st1.Pop(), " ")
	}
	fmt.Println()

	st2 := StackInt{
		s: []int{1, 2, 4, 9},
	}
	SortedInsertRec(&st2, 3)
	SortedInsertRec(&st2, 7)
	SortedInsertRec(&st2, 8)
	SortedInsertRec(&st2, 10)
	SortedInsertRec(&st2, 11)

	fmt.Print("st2 consist of following elements: ")
	for !st2.IsEmpty() {
		fmt.Print(st2.Pop(), " ")
	}
	fmt.Println()

}

// This is a very bad way that operates on the underlying data structure, you can accomplish this using just the stack ADT ooperations
func SortedInsert(s *StackInt, n int) {
	if s.IsEmpty() {
		s.s = append(s.s, n)
		return
	}

	i, j := 0, 0
	max := true
	for i < len(s.s) {
		if n < s.s[i] {
			j = i
			max = false
			break
		}
		i++
	}

	if max {
		s.s = append(s.s, n)
		return
	}

	ret := []int{}
	ret = append(ret, s.s[:j]...)
	ret = append(ret, n)
	ret = append(ret, s.s[j:]...)
	s.s = ret
}

// Method two is so much cleaner and easier to see
func SortedInsertRec(s *StackInt, n int) {
	var temp int
	if s.Length() == 0 || n > s.Top() {
		s.Push(n)
	} else {
		temp = s.Pop()
		SortedInsertRec(s, n)
		s.Push(temp)
	}
}
