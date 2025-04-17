package main

import (
	"fmt"
	"stack/stack-over-array"
)

func main() {
	stack := &MinStack{}
	stack.Push(5)
	stack.Push(10)
	stack.Push(4)
	stack.Push(9)
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	fmt.Println("Minimum:", stack.Min())
	stack.s.Print()
	stack.Pop()

	fmt.Println("\nMinimum: ", stack.Min())
	stack.s.Print()
	stack.Pop()

	fmt.Println("\nMinimum: ", stack.Min())
	stack.s.Print()
	stack.Pop()

	fmt.Println("\nMinimum: ", stack.Min())
	stack.s.Print()
	stack.Pop()

	fmt.Println("\nMinimum: ", stack.Min())
	stack.s.Print()
	stack.Pop()

	fmt.Println("\nMinimum: ", stack.Min())
	stack.s.Print()
}

type MinStack struct {
	s   stack.StackInt
	min stack.StackInt
}

func (m *MinStack) Pop() int {
	m.min.Pop()
	return m.s.Pop()
}

func (m *MinStack) Push(value int) {
	// If the value we are adding is a new minimum, then add it to the min stack
	// As soon as we pop off the current value from both stacks, it stops being the minimum by virtue of no longer existing in the min stack
	if value < m.min.Top() || m.min.IsEmpty() {
		m.min.Push(value)
	} else {
		// otherwise use the old minimum as this is further down the main stack s
		m.min.Push(m.min.Top())
	}
	m.s.Push(value)
}

func (m *MinStack) Min() int {
	return m.min.Top()
}

// The idea here is that each value in the MinStack will have a corresponding minimum value stored in another stack.
// This way we outsource the tracking of the minimum value to the other stack and when we pop off the main stack, we just pop off the other stack.
// We don't need to traverse the MinStack to determine what the new minimum is, we've already stored this.
