package main

type Queue[T any] struct {
	s []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		s: make([]T, 0, 10),
	}
}

func (q *Queue[T]) Queue(elem T) {
	q.s = append(q.s, elem)
}

func (q *Queue[T]) Dequeue() T {
	elem := q.s[0]
	q.s = q.s[1:]

	return elem
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.s) == 0
}
