package main

import (
	"testing"
)

func TestSetGetArray(t *testing.T) {
	da := New(0, 100)

	for i := range 100 {
		da.Set(i, i)
	}

	for i := range 100 {
		val := da.Get(i)

		if val != i {
			t.Fatalf("dynamic array expected=%d. got=%d", i, val)
		}
	}
}

func TestGrowArray(t *testing.T) {
	da := New(0, 100)

	daCap := da.cap
	da.grow()

	if da.cap == daCap {
		t.Fatal("array capability has not changed after grow")
	}

	for i := range da.cap {
		da.Set(i, i)
	}

	for i := range da.cap {
		val := da.Get(i)

		if val != i {
			t.Fatalf("array mem invalidated expected=%d. got=%d", i, val)
		}
	}
}

func BenchmarkPushArray(b *testing.B) {
	da := New(0, 1)
	b.ResetTimer()

	for i := range b.N {
		da.Push(i)
	}
}

func BenchmarkPushSlice(b *testing.B) {
	s := make([]int, 0, 10)

	b.ResetTimer()

	for i := range b.N {
		s = append(s, i)
	}
}

func TestPopArray(t *testing.T) {
	ran := 10_000
	da := New(0, 1)

	for i := range ran {
		da.Push(i)
	}

	for i := ran - 1; i > 0; i-- {
		if v := da.Pop(); v != i {
			t.Fatalf("array popped value expected=%d. got=%d", i, v)
		}
	}
}
