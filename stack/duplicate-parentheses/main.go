package main

import (
	"fmt"
	stack "stack/stack-over-array"
)

func main() {
	fmt.Println(DuplicateParentheses("(((a+(b))+(c+d)))"))
	fmt.Println(DuplicateParentheses("(((a+b))+c)"))
	fmt.Println(DuplicateParentheses("(b)"))
	fmt.Println(DuplicateParentheses("(a+b)"))
	fmt.Println(DuplicateParentheses("(a+((c+d)+d))"))
	fmt.Println(DuplicateParentheses("(((a+b)+c)+d)"))
	fmt.Println(DuplicateParentheses("(((a+b)+c))"))
}

// Assumes strings that have balanced brackets
func DuplicateParentheses(s string) bool {
	si := stack.StackInt{}
	sj := stack.StackInt{}
	i, j := 0, len(s)-1

	for i <= j {
		if string(s[i]) == "(" {
			si.Push(i)
		}
		if string(s[j]) == ")" {
			sj.Push(j)
		}
		i++
		j--
	}

	for si.Length() > 0 && sj.Length() != 0 {
		i := si.Pop()
		j := sj.Pop()

		// Case where there is no top to compare to
		if si.Length() == 0 && sj.Length() == 0 {
			if j-i <= 2 {
				return true
			} else {
				return false
			}
		}

		// The case where brackets are either empty () or have a single expression inside (a)
		if j-i <= 2 {
			return true
		}

		// expression exists so we can go to next loop iteration
		if i-si.Top() > 1 || sj.Top()-j > 1 {
			continue
		} else {
			// we have found a case where both i and si.Top and j AND sj.Top differ by only 1, so these are duplicate nested brackets with no expressions on both the left or right side
			return true
		}
	}
	return false
}

// Given a string "(((a+b))+c)" we need to determine if there are any redundant parentheses. This is obvious to humans looking at the string.
// Effectively what we require is an expression in between brackets.
// For any pair of brackets, as long as either the left side or the right side has an expression in between, then these are valid
// E.g. (((a+b)+c)+d) or (a+((c+d)+d))
// So effectively this problem can be reduced to keeping track of all the left and the right brackets using two stacks
// For the example (((a+b)+c)+d) our stacks would look like
// l_stack = [0,  1, 2]
// r_stack = [12, 9, 6]
// For any pair of brackets, we need to make sure that there is an expression contained within either the left or right side
// This translates to a condition that checks the difference betweent the current bracket and the NEXT bracket of that time, if the difference > 1 then we have an expression
