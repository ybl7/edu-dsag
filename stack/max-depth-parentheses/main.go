package main

import (
	"fmt"
	stack "stack/stack-over-array"
)

func main() {
	fmt.Println(MaxDepth("((()))"))
	fmt.Println(MaxDepth("((((A)))((((BBB()))))()()()())"))
	fmt.Println(MaxDepth("(()(())())"))
	fmt.Println(MaxDepth("()"))
}

func MaxDepth(s string) int {
	st := stack.StackInt{}
	var max int

	for _, r := range s {
		if string(r) == "(" {
			st.Push(1)
		}
		if string(r) == ")" {
			st.Pop()
		}
		if st.Length() > max {
			max = st.Length()
		}
	}

	return max
}
