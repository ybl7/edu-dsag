package main

import (
	"fmt"
	"queue/stack"
)

type QueueUsingStack struct {
	stk1 stack.StackInt
	stk2 stack.StackInt
}

func (que *QueueUsingStack) Add(value int) {
	que.stk1.Push(value)
}

// The idea here is that the first stack contains the elements, when we want to remove we loop over it and push all the elements to the second stack
// The second stack is the reverse of the first stack so we can pop off it while maintining LIFO
// We only ever reach the for loop to populate stk2 when we have exhausted all elements in stk2
func (que *QueueUsingStack) Remove() int {
	var value int
	if que.stk2.IsEmpty() == false {
		value = que.stk2.Pop()
		return value
	}
	for que.stk1.IsEmpty() == false {
		value = que.stk1.Pop()
		que.stk2.Push(value)
	}
	value = que.stk2.Pop()
	return value
}

func (que *QueueUsingStack) Length() int {
	return (que.stk1.Length() + que.stk2.Length())
}

func (que *QueueUsingStack) IsEmpty() bool {
	return (que.stk1.Length() + que.stk2.Length()) == 0
}

// Testing Code
func main() {
	que := new(QueueUsingStack)
	que.Add(1)
	que.Add(2)
	que.Add(3)
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
	fmt.Println(que.Remove())
}
