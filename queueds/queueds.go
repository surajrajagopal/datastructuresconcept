package queueds

import "fmt"

type Queue struct {
	Elements []int
	Size     int
}

func NewQueue(size int) *Queue {
	return &Queue{
		Elements: make([]int, 0),
		Size:     size,
	}
}

func (q *Queue) Getlength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Enqueue(val int) {
	if q.Getlength() == q.Size {
		fmt.Println("queue overflow")
		return
	}
	q.Elements = append(q.Elements, val)
}

func (q *Queue) Deque() int {
	if q.IsEmpty() {
		fmt.Println("queue underflow")
		return 0
	}

	res := q.Elements[0]

	if q.Getlength() == 1 {
		q.Elements = nil
		return res
	}

	q.Elements = q.Elements[1:]
	return res
}
