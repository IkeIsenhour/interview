package stack

import (
	"errors"
	"math"
)

type Stack struct {
	items []int
}

var (
	ErrEmptyStack = errors.New("empty stack")
)

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return math.MinInt, ErrEmptyStack
	}

	idx := len(s.items) - 1
	data := s.items[idx]
	s.items = s.items[:idx]

	return data, nil
}
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return math.MinInt, ErrEmptyStack
	}

	return s.items[len(s.items)-1], nil
}
