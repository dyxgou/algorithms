package main

import (
	"log/slog"
	"slices"
)

func Substring(s string) int {
	buf := make([]byte, 0, len(s))

upper:
	for i := range len(s) {
		ch := s[i]

		for _, bc := range buf {
			if ch == bc {
				continue upper
			}

			if len(buf) < len(s)-1-i {
				buf = slices.Delete(buf, 0, len(buf))
			}
		}

		buf = append(buf, ch)
	}

	slog.Info("lss", "buf", string(buf))
	return len(buf)
}
