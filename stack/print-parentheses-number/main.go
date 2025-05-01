package main

import (
	"fmt"
	"stack/stack-over-array"
)

func main() {
	fmt.Println(PrintParenthesisNumebr("(((a+(b))+(c+d)))"))
	fmt.Println(PrintParenthesisNumebr("(((a+b))+c)((("))
	fmt.Println(PrintParenthesisNumebr("(((a+b))+(c))"))
	fmt.Println(PrintParenthesisNumebr("(((a+b(())"))
}

func PrintParenthesisNumebr(s string) string {
	st := stack.StackInt{}
	ctr := 1
	out := ""

	for _, e := range s {
		if string(e) == "(" {
			st.Push(ctr)
			out += fmt.Sprintf("%d", ctr)
			ctr++
		}
		if string(e) == ")" {
			o := st.Pop()
			out += fmt.Sprintf("%d", o)
		}
	}

	return out
}

// Whevever you close a bracket it will awlays close the most recently opened bracket
// So we can just pop and push off the stack using a counter to keep track of which opening bracket the closing bracket belongs to
