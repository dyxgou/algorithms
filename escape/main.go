package main

func main() {
	esc := escapingValue()
	_ = esc
}

func escapingValue() *int {
	x := 19

	return &x
}
