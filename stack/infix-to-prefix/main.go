package main

import (
	"fmt"
	"stack/stack-over-array"
)

func main() {
	fmt.Println(InfixToPrefix("10+((3))*5/(16-4)")) // +10 * 3 / 5 - 16 4
	fmt.Println(InfixToPrefix("((5+6)+9)"))         // ++569
	fmt.Println(InfixToPrefix("((3+2)*(3-9)/(9))")) // /*+32-399
	fmt.Println(InfixToPrefix("13+((5-9)*(2*9))"))
	// fmt.Println(string('+'))
}

func InfixToPrefix(s string) string {
	// Reverse s first so we can iterate it from the back as if we are iterating from the front
	t := ""
	for _, e := range s {
		t = string(e) + t
	}
	out := ""
	st := stack.StackString{}

	for _, e := range t {
		if e == '(' {
			// Pop all operators in the parantheses and push
			for !st.IsEmpty() && st.Top() != ")" {
				out += string(st.Pop())
			}
			if !st.IsEmpty() && st.Top() == ")" {
				st.Pop()
			}
		} else if e == ')' {
			st.Push(string(e))
		} else if IsOperator(e) {
			if st.IsEmpty() {
				st.Push(string(e))
				continue
			}
			// If the incoming element has higher priority than the existing, then it can be added
			if Priority(string(e), st.Top()) {
				st.Push(string(e))
				continue
			}
			// Clear the stack of lower priority operators before pushing this one one
			for !Priority(string(e), st.Top()) {
				out += st.Pop()
			}
			st.Push(string(e))
		} else if IsDigit(e) {
			out += string(e)
		}
	}

	// Add whatever operators are left
	for !st.IsEmpty() {
		out += st.Pop()
	}

	ret := ""
	for _, e := range out {
		ret = string(e) + ret
	}
	return ret
}

func IsDigit(s rune) bool {
	if s <= '9' && s >= '0' {
		return true
	}
	return false
}

func IsOperator(s rune) bool {
	if s == '+' || s == '-' || s == '*' || s == '/' || s == '^' {
		return true
	}
	return false
}

func Priority(s, t string) bool {
	p := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"%": 2,
		"^": 3,
	}

	return p[s] > p[t]
}

// Let's first understand prefix notation with an example expression: (A + B * C) * (D - E) / F
// First, paranthesise everything according to BIDMAS: ((A + (B * C)) * ((D - E) / F))
// Next, shift all operators to the right of their operands: *(+(A*(BC))/(-(DE)F))
// Then remove brackets: *+A*BC/-DEF, how do we interpret this?

// Let's see this in action, iterating from the right our stack is [F E D], then we hit the first operator -
// We pop off E and D from the stack then perform D - E, then we get this result and continue
// After that we encounter /, so we pop the result of D - E and F off the stack and do, (D - E) / F and add this to the stack
// AT this point our stack just contains the single element [(D-E)/F]
// Continuing on, we add C and then B onto the stack so it is now [(D-E)/F C B] before encountering *, we pop B and C off an perform B * C then push this back
// So at this point our stack is [(D-E)/F B*C], continuing on we push A onto the stack [(D-E)/F B*C A] and before encountering the + operator
// So we pop off A and B*C from the stack and perform A+(B*C) and push this back onto the stack, our stack is now [(D-E)/F A+(B*C)]
// Finally we hit the leftmost * operator and pop off both of the remaining elements and perform A+(B*C) * (D-E)/F to get our final result
// So in essence, if we traverse the string from the right to the left, we can handle prefix expressions just like we handle postfix expressions

// A few things to notice: (1) the prefix order of operands is the same as the infix order
// (2) when two or more operators are in sequence, the highest priority operator is the RIGHTMOST one e.g. in /-DEF we do D-E first then (D-E)/F
// So it seems like we can use a similar logic to postfix, but instead traversing the string from the back and prepending instead of appending
// Let's see this in action

// idx					stk					out					comment
// 14					[]					F					prepend F to out
// 13					[/]					F					Push /, we don't know if there is a higher priority operation earlier in the string before it, so can't print it yet, remember from right to left, higher priority operations are first
// 12					[/ )]				F					Push )
// 11					[/ )]				EF					prepend E to out
// 10					[/ ) -]				EF					Push -, - is lower priority than ), but since Top is ) we push it anyway
// 9					[/ ) -]				DEF					prepend D to out
// 8					[/ ) -]			    DEF					we've hit an opening bracket, so we pop off the stack until we hit a closing bracket, in this case we prepend -
// 7					[*]					/-DEF				we reach * but it has the same priority as / so we pop and prepend / and push *
//																remember the whole idea is to decide for the value D-E. which should come first, / or *
//																i.e. which operator should be prepended, and the whole point of using a stack is to keep track of the operator before and after an operand
// 6					[* )]				/-DEF				Push ), brackets have the highest priority
// 5					[* )]				C/-DEF				prepend C to out
// 4					[* ) *]				C/-DEF				Push *, * is lower priority than ), but since Top is ) we push it anyway
// 3					[* ) *]				BC/-DEF				prepend B to out
// 2					[* ) +]				*BC/-DEF			+ is lower priority than Top so Pop and prepend, then Push +
// 1					[* ) +]				A*BC/-DEF			prepend A to out
// 0					[*]					+A*BC/-DEF			we've hit an opening bracket so prepend until we get the corresponding closing bracket
// 0					[]					*+A*BC/-DEF			We've iterated the whole string, so pop off the stack and prepend as necessary whatever is left

// For this example I used prepending, in practice adding to the start of a string is quite easy using string concatenation so we'll just stick to this.
// The alternative would be to append to the string and reverse it right at the end.

// Postfix thinking
// (A + (B + C)) --> ABC++
// ((A + B) + C) --> AB+C+
