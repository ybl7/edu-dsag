package main

import (
	"fmt"
	qarr "queue/queue-over-array"
	"queue/stack"
)

// We just dequeue the first k elements of the queue and push onto the stack, then we queue these back up
// However now we have the k elements reversed in the correct order, but they are after the first len - k elements that we want to keep in the same order
// so instead we need to dequeue on element at a time from the top and push it to the bottom
func reverseKElementInQueue(que *qarr.Queue, k int) {
	stk := new(stack.StackInt)
	i := 0
	var diff, temp int
	for que.Length() != 0 && i < k {
		stk.Push(que.Dequeue().(int)) // StackInt requires int, so cast the interface to int
		i++
	}
	for stk.Length() > 0 {
		que.Enqueue(stk.Pop())
	}
	diff = que.Length() - k
	for diff > 0 {
		temp = que.Dequeue().(int)
		que.Enqueue(temp)
		diff -= 1
	}
}

// Testing code
func main() {
	que := new(qarr.Queue)
	que.Enqueue(1)
	que.Enqueue(2)
	que.Enqueue(3)
	que.Enqueue(4)
	fmt.Print("Queue before Reversal: ")
	que.Print()
	fmt.Print("Queue after Reversal: ")
	reverseKElementInQueue(que, 2)
	que.Print()
}
