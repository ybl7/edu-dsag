package main

import (
	"fmt"
	stack "stack/stack-over-array"
)

func MaxContigBalancedParen(str string, size int) int {
	stk := stack.StackInt{}
	stk.Push(-1)
	length := 0
	for i := 0; i < size; i++ {
		if str[i] == '(' {
			stk.Push(i)
		} else {
			stk.Pop()
			if stk.Length() != 0 {
				if length < i-stk.Top() {
					length = i - stk.Top()
				}
			} else {
				stk.Push(i)
			}
		}
	}
	return 0
}

// Testing code
func main() {
	expn := "())((()))(())()(()"
	size := len(expn)
	value := MaxContigBalancedParen(expn, size)
	fmt.Println("MaxContigBalancedParenren :: ", value)
}

// This is a simple solution that hides a lot of clever reasoning.
// The easiest way to see it is to run through an intuitive example. The conclusion up front is that we are looking for the indexes that violate the conditions for having a valid substring of balanced paranthesis.
// Keep in mind the condition for a valid substring, that there are an equal number of closing brackets left of the opening brackets.

// Suppose our string is (added spaces for clarity)			     ")  )  (  (  )  )  )  (  )  )  (  )  (  (  )  )  )  )  (  (  ) "
// Let's write the index underneath the character                 0  1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16 17 18 19 20
// Visually it's easy to see invalid substrings                   _  _              _        _                    _  _
// The valid substrings are the substrings are
// (inclusive of first index, exclusive of last)
// [2:6] [7:9] [10:16] [18:20]

// So now let's suppose we start walking down this string, because we want to track indexes, we can use a stack for this (more on why in a bit).
// So let's walk down the string, and update our stack accordingly - it will become clear why this works. For now let's use two stacks for clarity, we'll come back at the end and explain why one stack is enough.

// Index | Operation | State          | Saved 		    | Comment
// 0       pop         []			    [0]		          We try to pop from an empty stack, we know we've got more closing brackets than opening brackets, so we've hit a partition in the string. Save this index to another stack.
// 1       pop         []               [0,1]		      We try to pop from an empty stack, we know we've got more closing brackets than opening brackets, so we've hit a partition in the string. Save this index to another stack.
// 2       push(2)     [2]
// 3       push(3)     [2,3]
// 4       pop         [2]
// 5       pop         []
// 6       pop         []               [0,1,6]           We try to pop from an empty stack, we know we've got more closing brackets than opening brackets, so we've hit a partition in the string. Save this index to another stack.
// 7       push(7)     [7]
// 8	   pop         []
// 9       pop         []               [0,1,6,9]         We try to pop from an empty stack, we know we've got more closing brackets than opening brackets, so we've hit a partition in the string. Save this index to another stack.
// 10      push(10)    [10]
// 11      pop         []
// 12      push(12)    [12]
// 13      push(13)    [12,13]
// 14      pop         [12]
// 15      pop         []
// 16      pop         []               [0,1,6,9,16]      We try to pop from an empty stack, we know we've got more closing brackets than opening brackets, so we've hit a partition in the string. Save this index to another stack.
// 17      pop         []               [0,1,6,9,16,17]   We try to pop from an empty stack, we know we've got more closing brackets than opening brackets, so we've hit a partition in the string. Save this index to another stack.
// 18      push(18)    [18]
// 19      pop         []

// So hopefully you see now that whenever we have a valid substring, the net effect on the stack is nothing, we pop whatever we push.
// Once we have the final Saved stack we can just pop each element off, then pop the next one off, calculate their difference, and update a variable, we take away 1 to get the correct value and to not include adjacent values that are one index apart.
// So this would look like: 17-16-1=0, 16-9-1=6, 9-6-1=2, 6-1-1=4, 1-0-1=0.

// Now is there a way we can do this without using two stacks? We only need to calculate a running max substring length. What if we use our stack to only store the best interval seen up until now?
// Suppose we always pop when we see closing brackets, why? Because if we end up popping more than we push, we'll end up at the empty stack, and we know for sure that we've reached the end of a subset.

// Index | Operation | State   			| Comment
// 0       pop         []			      Popping from empty
// 1       pop         []                 Popping from empty
// 2       push(2)     [2]
// 3       push(3)     [2,3]
// 4       pop         [2]
// 5       pop         []
// 6       pop         []                 Popping from empty
// 7       push(7)     [7]
// 8	   pop         []
// 9       pop         []                 Popping from empty
// 10      push(10)    [10]
// 11      pop         []
// 12      push(12)    [12]
// 13      push(13)    [12,13]
// 14      pop         [12]
// 15      pop         []
// 16      pop         []                 Popping from empty
// 17      pop         []                 Popping from empty
// 18      push(18)    [18]
// 19      pop         []

// Once again we see that popping from an empty stack is the indication that we've hit an invalid closing bracket index.
// Right now e don't actually store any useful information about such indexes in our stack. Can we use the stack to store the last invalid closing bracket?
// Yes we can, after every pop event against an empty stack, we push the current index into the stack, i.e. the index of the current invalid closing bracket.
// With this in mind, let's go through our table again. The rightmost colum is the final state for that index.

// Index | Operation | State   | IsEmpty? | Operation | State   | Comment
// 0       pop         []		 Yes	  | push(0)	    [0]	      After the initial pop we have an empty stack, push the index back in
// 1       pop         []        Yes      | push(1)     [1]       After the initial pop we have an empty stack, push the index back in
// 2       push(2)     [1,2]              |
// 3       push(3)     [1,2,3]            |
// 4       pop         [1,2]              |
// 5       pop         [1]                |
// 6       pop         []        Yes      | push(6)     [6]       After the initial pop we have an empty stack, push the index back in
// 7       push(7)     [6,7]              |
// 8	   pop         [6]                |
// 9       pop         []        Yes      | push(9)     [9]       After the initial pop we have an empty stack, push the index back in
// 10      push(10)    [9,10]             |
// 11      pop         [9]                |
// 12      push(12)    [9,12]             |
// 13      push(13)    [9,12,13]          |
// 14      pop         [9,12]             |
// 15      pop         [9]                |
// 16      pop         []        Yes      | push(16)    [16]
// 17      pop         []        Yes      | push(17)    [17]
// 18      push(18)    [17, 18]           |
// 19      pop         [17]               |

// We end up with all of the invalid closing brackets showing up in the stack at some point in time, in addition to all the indexes of the opening brackets.
// Note that when we pop, we are saying that we have necessarily hit a closing bracket index. But the trick is that we never store the indexes of valid closing brackets in our stack, only invalid ones.
// What if we have more opening brackets than we can close? Think about what pushing the index of invalid closing brackets ONLY when the stack is EMPTY means.
// If a stack is empty, it must mean that at least n pop operations happened after n push operations - otherwise we'd still have remaining indexes in the stack from the pushes.
// So the condition of the stack being empty AND popping from the empty stack is sufficient to ensure a balanced substring before the index of current.
// And since we only push closing brackets when they are invalid, when a pop results in an empty stack, what we just popped was the previous invalid closing bracket.
// So when a pop results in an empty stack, if we take the difference between the current index and the value top of stack just popped, this is index differnce between the two invalid closing brackets.

// Now if we wait until we actaully have an empty stack, after popping the last element, we'll always have an off by one error. So instead we look at top(stack).
// This allows us to look at the value of the index of the invalid closing bracket one iteration earlier, just before we pop it, therefore i is one less.
// That is we want the difference: i - top(stack). Let us see this in action. We calculate this invaraint after popping for any given index i, this is a clever trick so decrease our index by 1.
// You'll notice there are a lot of undef in the table below, this is because we can't read anything off the top empty stack. So we won't do the calculation in these instances and gate it with an if statement.

// Index | Invariant | Operation | State     | IsEmpty? | Operation | State   | Comment
// 0     | undef     | pop       | []		 | Yes	    | push(0)	   [0]	     After the initial pop we have an empty stack, push the index back in
// 1     | undef     | pop       | []        | Yes      | push(1)     [1]       After the initial pop we have an empty stack, push the index back in
// 2     |           | push(2)   | [1,2]     |          |
// 3     |           | push(3)   | [1,2,3]   |          |
// 4     | 4-2=2     | pop       | [1,2]     |          |
// 5     | 5-1=4     | pop       | [1]       |          |
// 6     | undef     | pop       | []        | Yes      | push(6)     [6]       After the initial pop we have an empty stack, push the index back in
// 7     |           | push(7)   | [6,7]     |          |
// 8	 | 8-6=2     | pop       | [6]       |          |
// 9     | undef     | pop       | []        | Yes      | push(9)     [9]       After the initial pop we have an empty stack, push the index back in
// 10    |           | push(10)  | [9,10]    |          |
// 11    | 1-9=2     | pop       | [9]       |          |
// 12    |           | push(12)  | [9,12]    |          |
// 13    |           | push(13)  | [9,12,13] |          |
// 14    | 14-12=2   | pop       | [9,12]    |          |
// 15    | 15-9=6    | pop       | [9]       |          |
// 16    | undef     | pop       | []        | Yes      | push(16)    [16]
// 17    | undef     | pop       | []        | Yes      | push(17)    [17]
// 18    |           | push(18)  | [17, 18]  |          |
// 19    | 19-17=2   | pop       | [17]      |          |

// What about when have too many opening brackets?
// So, by the time we get to the next invalid closing bracket and pop, the only index left in the stack is that of the last invalid closing bracket.
// Suppose we have some kind of string "()((((())))", it's clear in this case our stack will never be empty.
// i=6  [2,3,4,5,6]
// i=7  [2,3,4,5]    iv=7-5=2
// i=8  [2,3,4]      iv=8-4=4
// i=9  [2,3]        iv=9-3=6
// i=10 [2]          iv=10-2=8

// This question really hides a lot of complexity behind a very short solution.
// I think just using two stacks is so much more intuitive
