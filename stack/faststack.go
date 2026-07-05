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

func (s *FastStack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, EmptyStackErr
	}

	v := s.vals[s.len-1]
	s.len--

	return v, nil
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
