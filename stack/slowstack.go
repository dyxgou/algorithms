package main

type SlowStack[T any] struct {
	vals []T
}

func NewSlowStack[T any]() *SlowStack[T] {
	return &SlowStack[T]{
		vals: make([]T, 0, 10),
	}
}

func (s *SlowStack[T]) Insert(val T) {
	s.vals = append(s.vals, val)
}

func (s *SlowStack[T]) Pop() T {
	if s.IsEmpty() {
		panic("The stack is empty")
	}

	v := s.vals[len(s.vals)-1]

	s.vals = s.vals[:len(s.vals)-1]

	return v
}

func (s *SlowStack[T]) IsEmpty() bool {
	return len(s.vals) == 0
}
