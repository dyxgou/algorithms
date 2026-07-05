package main

import (
	"fmt"
	"strings"
)

type Reverser string

func createMarker(i int) string {
	if i-1 <= 0 {
		i = 1
	}
	return strings.Repeat(" ", i-1) + "^"
}

func showIndeces(buf []byte, i, j int) {
	fmt.Println(string(buf))

	fmt.Println(createMarker(i - 1))
	fmt.Println(createMarker(j))
}

func swap(buf []byte, i, j int) {
	temp := buf[i]
	buf[i] = buf[j]
	buf[j] = temp
}

func (r *Reverser) isLetter(idx int) bool {
	c := (*r)[idx]
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func (r *Reverser) Reverse() string {
	buf := make([]byte, 0, len(*r))
	buf = append(buf, (*r)...)

	lastIdx := len(buf) - 1
	idx := 0

	for idx < lastIdx {
		if !r.isLetter(idx) && !r.isLetter(lastIdx) {
			idx++
			lastIdx--
			continue
		}

		if !r.isLetter(idx) && r.isLetter(lastIdx) {
			idx++
			continue
		}

		if !r.isLetter(lastIdx) && r.isLetter(idx) {
			lastIdx--
			continue
		}

		swap(buf, idx, lastIdx)
		lastIdx--
		idx++
	}

	return string(buf)
}
