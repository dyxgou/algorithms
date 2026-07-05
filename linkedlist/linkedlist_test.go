package main

import (
	"iter"
	"log/slog"
	"testing"
)

func TestInsertHead(t *testing.T) {
	ll := New[int]()

	for i := range 100 {
		ll.InsertHead(i)
	}

	next, stop := iter.Pull(ll.Iter())
	defer stop()

	for i := 99; i >= 0; i-- {
		n, ok := next()
		if !ok {
			t.Fatal("LinkedList expected a value")
		}

		if n.val != i {
			t.Fatalf("LinkedList value expected=%d. got=%d", i, n.val)
		}
	}
}

func TestInsertElements(t *testing.T) {
	ll := New[int]()

	//["insertHead", 1, "insertHead", 2, "insertTail", 3, "insertTail", 4, "insertHead", 5, "get", 0, "get", 2, "get", 4, "remove", 2, "remove", 0, "insertHead", 6, "insertTail", 7, "getValues", "get", 5]

	ll.InsertHead(1)
	ll.InsertHead(2)
	ll.InsertTail(3)
	ll.InsertTail(4)
	ll.InsertHead(5)

	if !ll.Remove(2) {
		t.Fatalf("LinkedList element=%d expected to be deleted", ll.Get(2).val)
	}

	if !ll.Remove(0) {
		t.Fatalf("LinkedList element=%d expected to be deleted", ll.Get(0).val)
	}

	// 5 -> 2 -> 3 -> 4

	ll.InsertHead(6)
	ll.InsertTail(7)

	slog.Info("insert elements", "vals", ll.GetValues(), "want", []int{6, 2, 3, 4, 7})
}

func TestInsertTail(t *testing.T) {
	ll := New[int]()

	for i := range 100 {
		ll.InsertTail(i + 1)
	}

	next, stop := iter.Pull(ll.Iter())
	defer stop()

	for i := range 100 {
		n, ok := next()
		if !ok {
			t.Fatalf("LinkedList value expected=%d", i)
		}

		if n.val != i {
			t.Fatalf("LinkedList value expected=%d. got=%d", i, n.val)
		}
	}
}

func TestGetElement(t *testing.T) {
	tt := struct {
		ran, idx, want int
	}{
		ran:  100,
		idx:  50,
		want: 49,
	}

	ll := New[int]()

	for i := range tt.ran {
		ll.InsertHead(i)
	}

	n := ll.Get(tt.idx)
	if n == nil {
		t.Fatal("LinkedList element does not exists")
	}

	if n.val != tt.want {
		t.Fatalf("LinkedList get element expected=%d. got=%d", tt.want, n.val)
	}
}

func TestRemoveElement2(t *testing.T) {
	ll := New[int]()

	ll.InsertHead(1)
	ll.InsertTail(2)
	ll.InsertHead(0)

	slog.Info("LinkedList values", "vals", ll.GetValues())

	if !ll.Remove(1) {
		t.Fatal("element expected to be removed")

	}

	slog.Info("LinkedList values", "vals", ll.GetValues())

}

func TestRemoveElement(t *testing.T) {
	tt := struct {
		ran, idx int
	}{
		ran: 100,
		idx: 50,
	}

	ll := New[int]()

	for i := range tt.ran {
		ll.InsertHead(i)
	}

	n1 := ll.Get(tt.idx)
	if n1 == nil {
		t.Fatal("LinkedList n1 element does not exists")
	}

	if !ll.Remove(tt.idx) {
		t.Fatal("LinkedList element expected to be deleted")
	}

	n2 := ll.Get(tt.idx)
	if n2 == nil {
		t.Fatal("LinkedList n2 element does not exists")
	}

	if n1.val == n2.val {
		t.Fatalf("LinkedList removed element is still on the list")
	}
}

func TestGetElements(t *testing.T) {
	ll := New[int]()
	ran := 100

	for i := range ran {
		ll.InsertTail(i)
	}

	vals := ll.GetValues()

	for i := range len(vals) {
		if i != vals[i] {
			t.Fatalf("LinkedList current value expected=%d. got=%d", i, vals[i])
		}
	}
}
