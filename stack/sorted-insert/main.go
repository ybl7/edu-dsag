package main

import "fmt"

func main() {
	stack := new(StackInt)

	stack.SortedInsert(6)
	stack.SortedInsert(3)
	stack.SortedInsert(2)
	stack.SortedInsert(5)

	fmt.Println("Top() of the stack is: ", stack.Top())
	fmt.Print("Stack consist of following elements: ")
	for !stack.IsEmpty() {
		fmt.Print(stack.Pop(), " ")
	}
	fmt.Println()

	st1 := StackInt{
		s: []int{1, 2, 4, 9},
	}
	st1.SortedInsert(3)
	st1.SortedInsert(7)
	st1.SortedInsert(8)
	st1.SortedInsert(10)
	st1.SortedInsert(11)

	fmt.Print("st1 consist of following elements: ")
	for !st1.IsEmpty() {
		fmt.Print(st1.Pop(), " ")
	}
	fmt.Println()

	st2 := StackInt{
		s: []int{1, 2, 4, 9},
	}
	st2.SortedInsertRec(3)
	st2.SortedInsertRec(7)
	st2.SortedInsertRec(8)
	st2.SortedInsertRec(10)
	st2.SortedInsertRec(11)

	fmt.Print("st2 consist of following elements: ")
	for !st2.IsEmpty() {
		fmt.Print(st2.Pop(), " ")
	}
	fmt.Println()

}

// This is a very bad way that operates on the underlying data structure, you can accomplish this using just the stack ADT ooperations
func (s *StackInt) SortedInsert(n int) {
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
func (s *StackInt) SortedInsertRec(n int) {
	var temp int
	if s.Length() == 0 || n > s.Top() {
		s.Push(n)
	} else {
		temp = s.Pop()
		s.SortedInsertRec(n)
		s.Push(temp)
	}
}
