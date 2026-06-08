package main

type Stack struct {
	cons []int
	len  int
}

func NewStack() *Stack {
	return &Stack{
		cons: make([]int, 0, 10),
		len:  0,
	}
}

func (s *Stack) Insert(conn int) {
	if s.len < len(s.cons) {
		s.cons[s.len] = conn
		s.len++
		return
	}

	s.cons = append(s.cons, conn)
	s.len++
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack doesn't have any values")
	}

	v := s.cons[s.len-1]
	s.len--

	return v
}

func (s *Stack) Clear() {
	clear(s.cons[s.len:])
}

func (s *Stack) Peek() int {
	return s.cons[s.len-1]
}

func (s *Stack) IsEmpty() bool {
	return s.len <= 0
}
