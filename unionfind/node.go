package main

type Stringable interface {
	String() string
}

type Node[T Stringable] struct {
	val   T
	conns []int
}

func NewNode[T Stringable](val T) *Node[T] {
	return &Node[T]{
		val:   val,
		conns: make([]int, 0, 1),
	}
}

func (n *Node[T]) String() string {
	return n.val.String()
}
