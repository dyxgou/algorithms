package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	for i := range 100 {
		s.Insert(i)
	}

	for range 95 {
		s.Pop()
	}

	s.Clear()

	s.Insert(100)
	if pk := s.Peek(); pk != 100 {
		t.Fatalf("peeked elem expected=%d. got=%d", 100, pk)
	}
}
