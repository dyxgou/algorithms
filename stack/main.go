package main

import "log/slog"

func main() {
	s := NewSlowStack[int]()

	s.Insert(1)
	s.Insert(2)
	s.Insert(3)
	s.Insert(4)
	s.Insert(5)
	s.Insert(6)
	s.Insert(7)
	s.Insert(8)
	s.Insert(9)
	s.Insert(10)
	s.Insert(11)

	v := s.Pop()
	slog.Info("popped element", "v", v)
}
