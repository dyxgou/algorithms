package main

import (
	"iter"
	"log/slog"
	"testing"
)

func pairs[Slice ~[]E, E any](s Slice) iter.Seq2[E, E] {
	return func(yield func(E, E) bool) {
		for i := 0; i < len(s); i += 2 {
			v1 := s[i]

			if i+1 == len(s) {
				var zero E
				if !yield(v1, zero) {
					return
				}

				return
			}

			v2 := s[i+1]

			if !yield(v1, v2) {
				return
			}
		}
	}
}

func FuzzGraphInsert(f *testing.F) {
	tt := []int{10, 100, 1000, 1, 2, 3, 4}
	g := New[StringerInt]()

	for v1, v2 := range pairs(tt) {
		slog.Info("pairs iterator vals", "v1", v1, "v2", v2)
		f.Add(v1, v2)
	}

	f.Fuzz(func(t *testing.T, v1 int, v2 int) {
		n1 := g.Insert(StringerInt(v1))
		n2 := g.Insert(StringerInt(v2))

		g.Connect(n1, n2)

		if int(n1.val) != v1 {
			t.Fatalf("inserted node value expected=%d. got=%d", v1, n1.val)
		}

		if int(n2.val) != v2 {
			t.Fatalf("inserted node value expected=%d. got=%d", v1, n2.val)
		}

		if ok := g.IsConnected(n1, n2); !ok {
			t.Fatalf("is connected graph expected=%t. got=%t", true, ok)
		}
	})
}

func BenchmarkInsertGraph(b *testing.B) {
	g := New[StringerInt]()
	b.ResetTimer()

	for i := range b.N {
		g.Insert(StringerInt(i))
	}
}
