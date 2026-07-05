package main

import (
	"context"
	"log/slog"
	"testing"
)

type InterA interface {
	method()
}

type InterB interface {
	method()
}

type StructA []int

func (s *StructA) method() {
	slog.Info("StructA method call")
}

type StructB int

func (s *StructB) method() {
	slog.Info("StructB method call")

	context.WithValue(parent Context, key any, val any)
}

func TestFifth(t *testing.T) {
	var a InterA
	a = new(StructA)

	var b InterB
	b = new(StructB)

	b.method()
	b = a
	b.method()

	if b != a {
		t.Fatalf("interfaces are not equal. a=%v. b=%v", a, b)
	}
}
