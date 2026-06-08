package main

import (
	"testing"
)

func BenchmarkPopFastStack(b *testing.B) {
	b.Skip()
	s := NewFastStack[int]()
	for range b.N {
		s.Insert(1)
	}
	b.ResetTimer()

	// for range b.N {
	// 	s.Pop()
	// }
}

func BenchmarkPopSlowStack(b *testing.B) {
	b.Skip()
	s := NewSlowStack[int]()
	for range b.N {
		s.Insert(1)
	}

	b.ResetTimer()

	for range b.N {
		s.Pop()
	}
}

func BenchmarkInsertFastStack(b *testing.B) {
	s := NewFastStack[int]()
	b.ResetTimer()

	for b.Loop() {
		s.Insert(1)
	}
}

func BenchmarkInsertSlowStack(b *testing.B) {
	s := NewSlowStack[int]()
	b.ResetTimer()

	for b.Loop() {
		s.Insert(1)
	}
}
