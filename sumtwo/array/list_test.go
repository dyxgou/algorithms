package main

import (
	"iter"
	"log/slog"
	"testing"
)

func TestAddTwo1(t *testing.T) {
	tests := []struct {
		name   string
		l1Vals []int
		l2Vals []int
		res    []int
	}{
		{
			name:   "Add two numbers with single carriage",
			l1Vals: []int{5, 6, 4},
			l2Vals: []int{2, 4, 3},
			res:    []int{7, 0, 8},
		},
		{
			name:   "Add two empty lists",
			l1Vals: []int{0},
			l2Vals: []int{0},
			res:    []int{0},
		},
		{
			name:   "Add two list with multiple carriages",
			l1Vals: []int{9, 9, 9, 9, 9, 9, 9},
			l2Vals: []int{9, 9, 9, 9},
			res:    []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := new(ListNode)
			l1.Fill(tt.l1Vals)

			l2 := new(ListNode)
			l2.Fill(tt.l2Vals)

			lr := addTwoNumbers(l1, l2)

			next, stop := iter.Pull(lr.Iter())
			defer stop()

			for _, res := range tt.res {
				v, ok := next()
				if !ok {
					t.Fatalf("result list value doesn't exists. expected=%d. rl=%v", res, lr)
				}

				if res != v {
					t.Fatalf("result list value expected=%d. got=%d. lr=%v", res, v, lr)
				}
			}

			slog.Info("add two numbers", "lr", lr, "want", tt.res)
		})
	}

}
