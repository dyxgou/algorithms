package main

import "log/slog"

func main() {
	// Two packets will come at a second
	const arrivalRate = 2
	packets := []int{1, 2, 3, 4, 5, 6, 7, 5, 5}
	s := New[int]()

	for i := 0; i < len(packets); i += 2 {
		s.Queue(packets[i])
		if i+1 <= len(packets)-1 {
			s.Queue(packets[i+1])
		}
	}

	prev := s.Dequeue()
	sec := 1
	for !s.IsEmpty() {
		cur := s.Dequeue()

		slog.Info("time", "prev", prev, "cur", cur)
		if prev == cur {
			break
		}

		if !s.IsEmpty() {
			prev = s.Dequeue()
			sec++
		}
	}

	slog.Info("halted at", "second", sec)
}
