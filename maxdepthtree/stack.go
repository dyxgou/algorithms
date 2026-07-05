package main

import "errors"

type Stack[T any] struct {
	vals []T
}

var EmptyStackErr = errors.New("Empty Stack")

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		vals: make([]T, 0, 10),
	}
}

func (s *Stack[T]) Insert(val T) {
	s.vals = append(s.vals, val)
}

func (s *Stack[T]) Pop() T {
	v := s.vals[len(s.vals)-1]

	s.vals = s.vals[:len(s.vals)-1]

	return v
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.vals) == 0
}
