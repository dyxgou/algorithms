package main

type Graph[T Stringable] struct {
	nodes []*Node[T]
}

func New[T Stringable]() *Graph[T] {
	return &Graph[T]{
		nodes: make([]*Node[T], 0, 10),
	}
}

func (g *Graph[T]) Insert(val T) {
	n := NewNode(val)

	g.nodes = append(g.nodes, n)
}
