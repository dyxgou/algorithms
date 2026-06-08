package main

import (
	"strconv"
	"strings"
)

type Stringable interface {
	String() string
}

type Node[T Stringable] struct {
	val   T
	index int
	conns []int
}

func NewNode[T Stringable](val T, index int) *Node[T] {
	return &Node[T]{
		val:   val,
		index: index,
		conns: make([]int, 0, 2),
	}
}

func (n *Node[T]) addConn(dst int) {
	n.conns = append(n.conns, dst)
}

func (n *Node[T]) String() string {
	var sb strings.Builder

	sb.WriteByte('{')
	sb.WriteString("val=")
	sb.WriteString(n.val.String())
	sb.WriteByte(',')
	sb.WriteString("idx=")
	sb.WriteString(strconv.Itoa(n.index))
	sb.WriteByte('}')

	return sb.String()
}
