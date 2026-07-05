package main

import (
	"iter"
	"log/slog"
)

type Node[T any] struct {
	val  T
	next *Node[T]
}

func (n *Node[T]) Next() *Node[T] {
	return n.next
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{
		val: val,
	}
}

type LinkedList[T any] struct {
	head *Node[T]
	len  int
}

func New[T any]() *LinkedList[T] {
	return new(LinkedList[T])
}

func (l *LinkedList[T]) InsertHead(val T) {
	n := NewNode(val)
	l.len++

	if l.head == nil {
		l.head = n
		return
	}

	n.next = l.head
	l.head = n
}

func (l *LinkedList[T]) InsertTail(val T) {
	if l.head == nil {
		l.head = NewNode(val)
		l.len++
		return
	}

	for n := l.head; n != nil; n = n.Next() {
		if n.next == nil {
			n.next = NewNode(val)
			l.len++
			return
		}
	}
}

func (l *LinkedList[T]) Get(i int) *Node[T] {
	var idx int

	for n := range l.Iter() {
		if idx == i {
			return n
		}

		idx++
	}

	return nil
}

func (l *LinkedList[T]) Remove(i int) bool {
	if i == 0 {
		if l.len == 1 {
			l.head = nil
		} else {
			l.head = l.head.next
		}

		l.len--
		return true
	}

	if i+1 > l.len-1 {
		return false
	}

	next, stop := iter.Pull(l.Iter())
	defer stop()

	for range i - 1 {
		next()
	}

	prev, ok := next()
	if !ok {
		return false
	}

	prev.next = prev.next.next
	l.len--
	slog.Info("removing elements", "prev", prev, "i", i, "vals", l.GetValues())

	return true
}

func (l *LinkedList[T]) GetValues() []T {
	buf := make([]T, 0, l.len)

	for n := range l.Iter() {
		buf = append(buf, n.val)
	}

	return buf
}

func (l *LinkedList[T]) Iter() iter.Seq[*Node[T]] {
	return func(yield func(*Node[T]) bool) {
		for n := l.head; n != nil; n = n.Next() {
			if !yield(n) {
				return
			}
		}
	}
}
