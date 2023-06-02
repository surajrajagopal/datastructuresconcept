package stackds

import (
	"errors"
	"sync"
)

type stack struct {
	Lock sync.Mutex // you don't have to do this if you don't want thread safety
	S    []int
}

func NewStack() *stack {
	return &stack{
		Lock: sync.Mutex{},
		S:    make([]int, 0),
	}
}

func (s *stack) IsEmpty() bool {
	return len(s.S) == 0
}
func (s *stack) Push(v int) {
	s.Lock.Lock()
	defer s.Lock.Unlock()
	s.S = append(s.S, v)
}

func (s *stack) Pop() (int, error) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if s.IsEmpty() {
		return 0, errors.New("empty stack")
	} else {
		res := s.S[len(s.S)-1]
		s.S = s.S[:len(s.S)-1]
		return res, nil
	}
}

/*
func main() {
	s := stackds.NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if len(s.S) > 0 {
		fmt.Println(s.Pop())
		fmt.Println(s.Pop())
		fmt.Println(s.Pop())
	}
}
*/
