package main

import (
	"sync"
	"testing"
)

func returnPointer() *int {
	x := 50

	return &x
}

func TestFirstA(t *testing.T) {
	x := returnPointer()

	if *x != 50 {
		t.Fatalf("x value expected=%d. got=%d", 50, *x)
	}
}

func TestFirstD(t *testing.T) {
	x := 50

	var wg sync.WaitGroup
	wg.Add(1)

	go func(num *int) {
		wg.Done()
		t.Logf("go routine number. x=%d", *num)
	}(&x)

	wg.Wait()
}

type StructFirst struct {
	val int
}

func (s StructFirst) ValueMethod() int {
	return s.val
}

func TestFirstE(t *testing.T) {
	s := StructFirst{val: 1}

	x := s.ValueMethod()

	if x != 1 {
		t.Fatalf("x value method expected=%d. got=%d", 1, x)
	}
}
