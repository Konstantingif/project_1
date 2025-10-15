package main

import (
	"fmt"
	"project/linkedlist"
	"project/queue"
	"project/stack"
)

func main() {

	var s stack.Stack[int]
	s.Push(100)
	s.Push(200)
	if val, ok := s.Pop(); ok {
		fmt.Println("Popped from stack:", val) 
	}

	var q queue.Queue[string]
	q.Enqueue("Go")
	q.Enqueue("is")
	q.Enqueue("great")
	if val, ok := q.Dequeue(); ok {
		fmt.Println("Dequeued from queue:", val) 
	}

	var l linkedlist.List[int]
	l.Append(1)
	l.Append(2)
	l.Prepend(0)
	fmt.Println("List values:", l.Values()) 
	fmt.Println("List length:", l.Len())    
}
