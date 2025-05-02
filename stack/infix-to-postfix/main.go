package main

import (
	"fmt"
	"stack/stack-over-array"
)

func main() {
	fmt.Println(InfixToPostfix("10+((3))*5/(16-4)"))
	fmt.Println(InfixToPostfix("((5+6)+9)"))
	fmt.Println(InfixToPostfix("((3+2)*(3-9)/(9))"))
	fmt.Println(InfixToPostfix("13+((5-9)*(2*9))"))
	// fmt.Println(string('+'))
}

func InfixToPostfix(s string) string {
	out := ""
	st := stack.StackString{}

	for _, e := range s {
		// If opening bracket always push onto the stack
		if e == '(' {
			st.Push("(")
		}
		// If closing bracket, keep popping until you get to the first opening bracket, add the operators you found to the output
		if e == ')' {
			for st.Top() != "(" && !st.IsEmpty() {
				v := st.Pop()
				out += v
			}
			if st.Top() == "(" && !st.IsEmpty() {
				st.Pop()
			}
		}
		if IsDigit(e) {
			out += string(e)
		}
		if IsOperator(e) {
			// If the current element is an operator of higher priority than the top of the stack, we push it it to the stack
			if st.IsEmpty() {
				st.Push(string(e))
				continue
			}
			if !st.IsEmpty() && Priority(string(e), st.Top()) {
				// We use a continue since the else part of this block is the for loop
				st.Push(string(e))
				continue
			}
			// Once we reach this stage we know that we have a same or lower priority operator, in which case we pop off all the equal or higher priority operators already on the stack
			// Remembering to add the the current operator at the end
			for !st.IsEmpty() && !Priority(string(e), st.Top()) {
				v := st.Pop()
				out += v
			}
			st.Push(string(e))
		}
	}

	// Add whatever operators are left
	for !st.IsEmpty() {
		v := st.Pop()
		out += v
	}
	return out
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

// If you don't know what postfix is in detail, you might struggle with this question as I did
// You might even attempt a recursive solution, save yourself a few hours and go read this first https://www.andrew.cmu.edu/course/15-200/s06/applications/ln/junk.html

// Let's take a few illustritive examples of turning infix to postfix
// 3 + 4 * 5 - 12 * (6 - 4)
// Then as detailed in the blog post above, we need to fully paranthesise, that is around every operator and it's operands, put brackets ensuring BIDMAS
// ((3 + (4 * 5) - (12 * (6 - 4)))
// Then move every operator the the right of it's operands and remove the brackets
// 3 4 5 * + 12 6 4 - * -
// Now we also need to know how to evaluate this, in fact postfix notation is very computer friendly
// To evaluate the scan our string, and populate a stack with the numbers as we see them, when we hit an operator we pop that last two numbers off the stack
// then we perform the operation against these two numbers and push the result onto the stack, let's see how this works
// [3 4 5] -> [3] and 5 * 4 -> [3 20] -> [] and 20 + 3 -> [23]
// [23 12 6 4] -> [23 12] and 6 - 4 -> [23 12 2] -> [23] and 12 * 2 -> [23 24] -> [] and 23 - 24 = -1
// Once the stack is empty we have our answer, and we can see that the correct value in indeed -1

// So now that we have a decent understanding of what postfix notation is an how to evaluate it, how can we go about translating from infix to postfix
// There are a few key things to note:
// 1) The order of operands stays the SAME regardless of notation
// 2) The order of operators CHANGES, in doing so we eliminate the need for brackets
// 3) When operators are consecutive, higher priority operators come FIRST
// 4) We only ever need to consider how to order two operators, since every operation can be expressed as some compostion of two operands (as complex as they may be)
//    this means we just need to decide if the next operator we see should come before or after the current one which depends on two factors, the implicit
//    order according to BIDMAS and whether the operator is in brackets

// So let's do an example of an algorithmic approach, suppose we are traversing 3 + 4 * 5 - 12 * (6 - 4) again
// Since we know the order of the operands will not change, whenever we see an operand, we can just output it to the output string immediately
// Let's use a table to store the current index (starting from 1), output string, and crucially a stack to store the operators (more on this later)

// idx			stack			out
// 1			[]				3
// 2			[+]				3
// 3			[+]				3 4
// 4			[+ *]			3 4
// 5			[+ *]			3 4 5

// At this point we should stop and reason a bit, why did we not print + to out yet? Why did we wait for *?
// Well the + operates on the 3 and 4, that's obvious, but the 4 is also operated on by *, so 4 is kind of like a decision point
// It's being operated on from both sides (+ to the left and * to the right), so we need to consider both sides before making a decision
// As it turns out the * has higher priority than the +, so we should output the *, right? Wrong!!!
// Why? Well let's think about what using a stack means, it means that we can pop, so it means the next operator we can access is the top most
// So if my topmost operator is the highest priority, this is fine, because it SHOULD be coming first in my postfix expression
// But you may say, * is higher priority than +, so why can't we output it? Well, the * is just one operator that operates on the 5
// We have no idea what the operator to the right of the 5 is, it could well be higher priority than the *, think exponeniation
// So the only time we know we can output an operator is when we know the next operator for sure is lower priority
// So on the sixth iteration above we will see that we want to push - onto the stack and that this is lower than *, so we pop off the stack and add to out
// (see table below for summary) at this point we are left with just [+] in the stack, should we now push our - to the stack?

// The answer is once again no, notice that in our examples when two operators are adjacent, the higher priority one comes first
// We enforce this by using the stack which will reverse the order of operations by virtue of being FILO
// But since + and - have the same priority we don't need to do any such reversal and can just have the + output immediately
// I'll repeat, the only time we need to kick an operator to the back is when there is a higher priority one coming after it, in which case postfix requires
// that it comes first. Finally we are in a position to push - onto the stack, so with that, let's go back to our table:

// idx			stack			out
// 1			[]				3
// 2			[+]				3
// 3			[+]				3 4
// 4			[+ *]			3 4
// 5			[+ *]			3 4 5
// 6			[+]				3 4 5 *
// 6			[]              3 4 5 * +
// 6			[-]				3 4 5 * +

// A lot happened in the sixth index, but let's continue

// idx			stack			out
// 1			[]				3
// 2			[+]				3
// 3			[+]				3 4
// 4			[+ *]			3 4
// 5			[+ *]			3 4 5
// 6			[+]				3 4 5 *
// 6			[]              3 4 5 * +
// 6			[-]				3 4 5 * +
// 7			[-]				3 4 5 * + 12
// 8			[- *]			3 4 5 * + 12
// 8			[- * (]			3 4 5 * + 12

// Ok this needs a little explaining, why did we push the bracket onto the stack? Well because it's an operator!
// Not in the traditional sense of being a unary/binary operator to produce an output, but one to enforce ordering.
// The cool part is that our logic doesn't have the change, the B in Bidmas stands for brackets, and they are the highest priority operator.
// So there is no chance of anything coming on later thta will force us to pop off the bracket, not until we get another closing bracket.

// idx			stack			out
// 1			[]				3
// 2			[+]				3
// 3			[+]				3 4
// 4			[+ *]			3 4
// 5			[+ *]			3 4 5
// 6			[+]				3 4 5 *
// 6			[]              3 4 5 * +
// 6			[-]				3 4 5 * +
// 7			[-]				3 4 5 * + 12
// 8			[- *]			3 4 5 * + 12
// 9			[- * (]			3 4 5 * + 12
// 10			[- * (]		    3 4 5 * + 12 6
// 11			[- * ( -]       3 4 5 * + 12 6
// 12			[- * ( -]       3 4 5 * + 12 6 4
// 13			[- * (]			3 4 5 * + 12 6 4 -
// 13			[- *]			3 4 5 * + 12 6 4 -

// At idx 11 we are trying to push - onto the stack, however we already have a bracket, usually we'd start popping off the stack
// since we are pushing something of lower priority after the bracket, but brackets are special and don't get popped off until we encounter a closing bracket
// Effectively after an opening bracket, we treat the stack as if it is a new substack that has no reference to the operators that came before it
// At idx 12 we add the operand 4 to the output, then at idx 13 we hit the closing bracket, at which point we pop off the stack until we reach the opening bracket
// and we pop off the opening bracket too. Finally we reach the end of our string so we just pop off whatever remaining symbols we have left

// What we havent discussed is how nested brackets will be handled, suppose we have some example A O1 ((B O2 C) O3 D), and assume 01 > 02 > 03 in precedence
// let's run through our algorithm quickly to see if this really works

// idx			stack			out					comment
// 1			[]				A
// 2			[O1]			A
// 3			[01 (]			A
// 4			[O1 ( (]		A
// 5			[O1 ( (]		A B
// 6			[O1 ( ( O2]		A B
// 7			[O1 ( ( O2]		A B C
// 8			[01 (]		    A B C O2			hit closing bracket, pop off O2 and one (
// 9		    [O1 ( O3]		A B C O2
// 10			[O1 ( O3]		A B C O2 D
// 11			[O1]			A B C O2 D O3		hit closing bracket, pop off O3 and one (
// 12			[]				A B C O2 D O3 O1	reached end of string, put all remaining symbols at the end

// Which in english reads to B O2 C, then do (B O2 C) O3 D, then do A O1 ((B O2 C) O3 D) as expected.
