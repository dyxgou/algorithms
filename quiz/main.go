package main

import "fmt"

const (
	i = 8
	j
	k
)

func main() {

	a := []string{"a", "b"}
	a = a[:0]
	fmt.Print(a, len(a), cap(a))
}
