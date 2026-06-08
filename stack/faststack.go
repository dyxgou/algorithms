package main

type FastStack[T any] struct {
	vals []T
	len  int
}

func NewFastStack[T any]() *FastStack[T] {
	return &FastStack[T]{
		vals: make([]T, 0, 10),
		len:  0,
	}
}

func (s *FastStack[T]) Insert(val T) {
	if s.len < len(s.vals) {
		s.vals[s.len] = val
		s.len++
		return
	}

	s.vals = append(s.vals, val)
	s.len++
}

func (s *FastStack[T]) Pop() T {
	if s.IsEmpty() {
		return s.vals[0]
	}

	v := s.vals[s.len-1]
	s.len--

	return v
}

func (s *FastStack[T]) Clear() {
	clear(s.vals[s.len:])
}

func (s *FastStack[T]) Peek() T {
	return s.vals[s.len-1]
}

func (s *FastStack[T]) IsEmpty() bool {
	return s.len == 0
}
