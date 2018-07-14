package stack

import (
	"errors"
)

var errOutRange error = errors.New("Ouf of range!")
var errNullTop error = errors.New("top is nil!")

type Stack struct {
	data []int
}

func NewStack() *Stack {
	r := new(Stack)
	return r
}

func (s *Stack) StackPush(d int) {
	s.data = append(s.data, d)
}

func (s *Stack) StackPop() (int, error) {
	if len(s.data) == 0 {
		return 0, errOutRange
	}
	r := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]

	return r, nil
}

func (s *Stack) StackTop() (int, error) {
	if len(s.data) == 0 {
		return 0, errNullTop
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) StackSize() int {
	return len(s.data)
}
