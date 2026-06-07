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

	g.Insert(1)
	g.Insert(2)
	g.Insert(3)
	g.Insert(4)
	g.Insert(5)
	g.Insert(6)
	g.Insert(7)
	g.Insert(8)

	slog.Info("graph nodes", "ns", g.nodes)
}
