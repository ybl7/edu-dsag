package main

import (
	"fmt"
	stack "stack/stack-over-array"
)

func main() {
	fmt.Println(BalancedParenthesis("((}))"))
	fmt.Println(BalancedParenthesis("{[}]"))
	fmt.Println(BalancedParenthesis("{()}["))
	fmt.Println(BalancedParenthesis("{()}[]"))
}

func BalancedParenthesis(s string) bool {
	op := stack.StackInt{}

	for _, r := range s {
		if string(r) == "{" {
			op.Push(1)
		}
		if string(r) == "(" {
			op.Push(2)
		}
		if string(r) == "[" {
			op.Push(3)
		}

		if string(r) == "}" {
			o := op.Pop()
			if o != 1 {
				return false
			}
		}
		if string(r) == ")" {
			o := op.Pop()
			if o != 2 {
				return false
			}
		}
		if string(r) == "]" {
			o := op.Pop()
			if o != 3 {
				return false
			}
		}
	}

	return op.IsEmpty()
}
