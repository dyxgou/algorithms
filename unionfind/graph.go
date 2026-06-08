package main

import (
	"log/slog"
	"slices"
)

type Connection struct {
	src, dst int
}

func (c Connection) IsDirect(src, dst int) bool {
	return c.src == src && c.dst == dst
}

func (c Connection) IsSameSrc(src int) bool {
	return c.src == src
}

func (c Connection) IsSameDst(dst int) bool {
	return c.dst == dst
}

func NewConn(src, dst int) Connection {
	return Connection{src, dst}
}

type Graph[T Stringable] struct {
	nodes []*Node[T]
	conns []Connection
}

func New[T Stringable]() *Graph[T] {
	return &Graph[T]{
		nodes: make([]*Node[T], 0, 10),
		conns: make([]Connection, 0, 10),
	}
}

func (g *Graph[T]) Insert(val T) *Node[T] {
	n := NewNode(val, len(g.nodes))

	g.nodes = append(g.nodes, n)

	return n
}

func (g *Graph[T]) getNodeConns(idx int) []int {
	return g.nodes[idx].conns
}

// This function assumes that the given nodes are already
// contained in the graph as it uses its `index` field
func (g *Graph[T]) Connect(src *Node[T], dst *Node[T]) {
	// A Node is not allowed to have a connection to itself
	// If it does, Connect() fails silently
	if src.index == dst.index {
		return
	}

	src.addConn(dst.index)
	dst.addConn(src.index)
}

func (g *Graph[T]) IsConnected(src *Node[T], dst *Node[T]) bool {
	if len(src.conns) == 0 {
		return false
	}

	if src.index == dst.index {
		return false
	}

	s := NewStack()
	for _, conn := range src.conns {
		s.Insert(conn)
	}

	for !s.IsEmpty() {
		c := s.Pop()
		slog.Info("conn popped", "c", c, "n", g.nodes[c])

		conns := g.getNodeConns(c)

		if slices.Contains(conns, dst.index) {
			return true
		}
	}

	return false
}
