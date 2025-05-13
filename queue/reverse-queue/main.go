package main

import (
	"fmt"
	qarr "queue/queue-over-array"
	"queue/stack"
)

// Empty the queue sequentially onto a stack then repopulate the queueu in reverse order by pushing off the stack
func reverseQueue(que *qarr.Queue) {
	stk := new(stack.StackInt)
	for que.Length() != 0 {
		stk.Push(que.Dequeue().(int))
	}
	for stk.Length() > 0 {
		que.Enqueue(stk.Pop())
	}
}

// Testing code
func main() {
	que := new(qarr.Queue)
	que.Enqueue(1)
	que.Enqueue(2)
	que.Enqueue(3)
	que.Enqueue(4)
	fmt.Print("Queue before reversal : ")
	que.Print()
	reverseQueue(que)
	fmt.Print("Queue after reversal : ")
	que.Print()
}
