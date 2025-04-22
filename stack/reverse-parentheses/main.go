package main

import (
	"fmt"
	stack "stack/stack-over-array"
)

func main() {
	fmt.Println(ReverseParenthese(")(())((("))
	fmt.Println(ReverseParenthese(")))(((("))
	fmt.Println(ReverseParenthese(")((("))
	fmt.Println(ReverseParenthese(")()()("))
	fmt.Println(ReverseParenthese(""))
}

// Returns the minimum number of swaps required to make a string consisting of parentheses balances
func ReverseParenthese(s string) int {
	st := stack.StackInt{}
	rev := 0

	for _, e := range s {
		if string(e) == "(" {
			st.Push(1)
		}
		if string(e) == ")" {
			if st.Length() != 0 {
				st.Pop()
			} else {
				st.Push(1)
				rev++
			}
		}
	}

	if st.Length()%2 == 0 {
		return rev + st.Length()/2
	}

	return -1
}

// Suppose we have some random ")))(((()((()(((())())()))))"
// As we iterate down the string If at any point we have more closing brackets than opening brackets then these need to be reversed
// If we push opening brackets onto the stack, then the condition (as discussed in other questions) where we pop off the empty stack is equivalent to having too many closing brackets
// Then we must reverse these brackets and push onto the stack, because we just added another opening bracket, so we add one to our reversal counter
// Next we have the case where we open too many brackets, in which case these need to be revered, at the end of the iteration we just check the length of the stack
// And since we need to reverse exactly half of the remaining openning brackets to close then, we increment our counter with stk.length/2
