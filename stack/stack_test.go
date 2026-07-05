package main

import (
	"os"
	"testing"
)

var s *SlowStack[int]

func TestMain(m *testing.M) {
	s = NewSlowStack[int]()

	code := m.Run()

	os.Exit(code)
}

func BenchmarkPopFastStack(b *testing.B) {
	b.SetParallelism(4)
}

func BenchmarkPopSlowStack(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			s.Insert(1)
			s.Pop()
		}
	})

	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			ss := NewFastStack[int]()
			ss.Insert(1)
			ss.Pop()
		}
	})
}
