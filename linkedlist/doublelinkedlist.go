package linkedlist

import "fmt"

type DoubleLLNode struct {
	Prev  *DoubleLLNode
	Value interface{}
	Next  *DoubleLLNode
}

type ListDLL struct {
	Head *DoubleLLNode
	Tail *DoubleLLNode
}

// Static Doubly-Linked-List Implemenentation
func NewDoubleLLNode(data interface{}) *DoubleLLNode {
	return &DoubleLLNode{
		Prev:  nil,
		Value: data,
		Next:  nil,
	}
}

// Dynamic Doubly-Linked-List Implemenentation
func (dll *ListDLL) Insert(data interface{}) {
	nextNode := &DoubleLLNode{Prev: nil, Value: data, Next: nil}

	if dll.Head == nil {
		dll.Head = nextNode
		dll.Tail = nextNode
	} else {
		p := dll.Head
		for p.Next != nil {
			p = p.Next
		}
		nextNode.Prev = p
		p.Next = nextNode
		dll.Tail = nextNode
	}
}

func DLLTraverse(head *DoubleLLNode) {
	ptr := head

	for ptr != nil {
		fmt.Println("data", ptr.Value)
		ptr = ptr.Next
	}
}

func ReverseDLLTraverse(tail *DoubleLLNode) {
	p := tail

	for p != nil {
		fmt.Println("data", p.Value)
		p = p.Prev
	}
}
