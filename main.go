package main

import (
	"datastructures/linkedlist"
	"datastructures/queueds"
	"fmt"
)

func main() {
	q := queueds.NewQueue(3)

	fmt.Println("Enqueue data")
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println(q.Elements)
	q.Enqueue(4)

	fmt.Println("Dequeue data")
	if len(q.Elements) > 0 {
		fmt.Println(q.Deque())
		fmt.Println(q.Deque())
		fmt.Println(q.Deque())
	}
	fmt.Println(q.Deque())

	// Single Linked List
	fmt.Println("Single Linked List")
	lis := linkedlist.List{}
	lis.Insert(10)
	lis.Insert(20)
	lis.Insert(30)
	linkedlist.Show(&lis)

	fmt.Println("Reverse Single Linked List")
	c := lis.ReverseLinkedList()
	linkedlist.ReverseTraverse(c)

	fmt.Println("Circular Linked List")
	cll := linkedlist.NewCircularLinkedList()
	cll.AddNode(10)
	cll.AddNode(20)
	cll.AddNode(30)
	cll.Traverse()

	fmt.Println("Double Linked List")
	head := linkedlist.NewDoubleLLNode(10)
	node2 := linkedlist.NewDoubleLLNode(20)
	node3 := linkedlist.NewDoubleLLNode(30)
	head.Next = node2
	node2.Prev = head
	node2.Next = node3
	node3.Prev = node2

	linkedlist.DLLTraverse(head)

	fmt.Println("Double Linked List from dynamic approach")
	dll := &linkedlist.ListDLL{}
	dll.Insert(10)
	dll.Insert(20)
	dll.Insert(30)
	fmt.Println("Traverse Forward : ")
	linkedlist.DLLTraverse(dll.Head)
	fmt.Println("Traverse Reverse : ")
	linkedlist.ReverseDLLTraverse(dll.Tail)
}
