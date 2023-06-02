package linkedlist

import "fmt"

type Node struct {
	info interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(data interface{}) {
	list := &Node{info: data, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		ptr := l.head
		for ptr.next != nil {
			ptr = ptr.next
		}
		ptr.next = list
	}
}

func (l *List) ReverseLinkedList() *Node {
	var prevptr *Node
	currptr := l.head

	for currptr != nil {
		next := currptr.next
		currptr.next = prevptr
		prevptr = currptr
		currptr = next
	}
	return prevptr
}

func Show(l *List) {
	ptr := l.head
	for ptr != nil {
		fmt.Println(ptr.info)
		ptr = ptr.next
	}
}

func ReverseTraverse(n *Node) {
	for n != nil {
		fmt.Println(n.info)
		n = n.next
	}
}
