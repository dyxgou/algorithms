package main

import (
	"log/slog"
	"strconv"
)

type StringableInt int

func (i StringableInt) String() string {
	return strconv.Itoa(int(i))
}

func main() {
	g := New[StringableInt]()

	n1 := g.Insert(1)
	n2 := g.Insert(2)
	n3 := g.Insert(3)
	n4 := g.Insert(4)
	n5 := g.Insert(5)
	n6 := g.Insert(6)
	n7 := g.Insert(7)
	n8 := g.Insert(8)

	// Cycle
	g.Connect(n1, n7) // n1 <-> n7
	g.Connect(n1, n8) // n1 <-> n8
	g.Connect(n7, n8) // n7 <-> n8
	g.Connect(n8, n6) // n7 <-> n8
	g.Connect(n6, n2) // n7 <-> n8
	g.Connect(n2, n3) // n7 <-> n8
	g.Connect(n4, n5)

	isConnected := g.IsConnected(n1, n4)

	slog.Info("isNodeConnected", "v", isConnected)
}
