package problem

import "math"

type Stack struct {
	items []int
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() int {
	if s.isEmpty() {
		return 0
	}

	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]

	return item
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Peek() int {
	if s.isEmpty() {
		return 0
	}

	return s.items[len(s.items)-1]
}

func asteroidCollision(asteroids []int) []int {
	s := Stack{}
	for _, asteroid := range asteroids {
		flag := true

		for !s.isEmpty() && s.Peek() > 0 && asteroid < 0 {
			if math.Abs(float64(s.Peek())) < math.Abs(float64(asteroid)) {
				s.Pop()
				continue
			} else if math.Abs(float64(s.Peek())) == math.Abs(float64(asteroid)) {
				s.Pop()
			}

			flag = false
			break
		}

		if flag {
			s.Push(asteroid)
		}
	}

	return s.items
}
