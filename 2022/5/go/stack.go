package main

import "errors"

type Stack struct {
	s []rune
}

func NewStack() *Stack {
	return &Stack{make([]rune, 0)}
}

func (s *Stack) Push(val rune) {
	s.s = append(s.s, val)
}

func (s *Stack) Pop() (rune, error) {
	if len(s.s) == 0 {
		return 0, errors.New("Empty Stack")
	}
	val := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return val, nil
}

func (s *Stack) Peek() (rune, error) {
	if len(s.s) == 0 {
		return 0, errors.New("Empty Stack")
	}
	return s.s[len(s.s)-1], nil
}
