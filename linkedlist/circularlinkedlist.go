package linkedlist

import "fmt"

type CircularLinkedList struct {
	head *Node
	tail *Node
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{
		head: nil,
		tail: nil,
	}
}

func (cll *CircularLinkedList) IsEmpty() bool {
	return cll.head == nil
}

func (cll *CircularLinkedList) AddNode(data interface{}) {
	newNode := &Node{info: data, next: nil}
	if cll.IsEmpty() {
		cll.head = newNode
		cll.tail = newNode
		newNode.next = cll.head
	} else {
		cll.tail.next = newNode
		cll.tail = newNode
		newNode.next = cll.head
	}
}

func (cll *CircularLinkedList) Traverse() {
	if cll.IsEmpty() {
		fmt.Println("Circular Linked List is empty")
	} else {
		current := cll.head
		for {
			fmt.Println("data", current.info)
			current = current.next
			if current == cll.head {
				break
			}
		}
		fmt.Printf("data %d\n", cll.head.info)
	}
}
