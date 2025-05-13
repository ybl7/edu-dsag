package main

import "fmt"

type Node struct {
	value int
	next  *Node
}
type QueueLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (q *QueueLinkedList) Size() int {
	return q.size
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueLinkedList) Peek() int {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0
	}
	return q.head.value
}

func (q *QueueLinkedList) Add(value int) {
	temp := &Node{value, nil}
	if q.head == nil {
		q.head = temp
		q.tail = temp
		q.tail.next = q.head
	} else {
		temp.next = q.head
		q.tail.next = temp
		q.tail = temp
	}
	q.size++
}

func (q *QueueLinkedList) Remove() (int, bool) {
	if q.IsEmpty() {
		fmt.Println("QueueEmptyException")
		return 0, false
	}
	var value int
	value = q.head.value
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
		q.tail.next = q.head
	}
	q.size--
	return value, true
}

func (q *QueueLinkedList) Print() {
	temp := q.head
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Println()
}

// Testing Code
func main() {
	q := new(QueueLinkedList)
	q.Add(1)
	q.Add(2)
	q.Add(3)

	for q.IsEmpty() == false {
		val, _ := q.Remove()
		fmt.Println(val, " ")
	}
}
