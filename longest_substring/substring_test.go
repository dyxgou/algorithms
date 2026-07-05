package main

import "testing"

func TestSubstring(t *testing.T) {
	tests := []struct {
		name, in string
		size     int
	}{
		{
			name: "Substring abc",
			in:   "abcabcbb",
			size: 3,
		},
		{
			name: "Substring just one letter",
			in:   "bbbbbbbb",
			size: 1,
		},
		{
			name: "Substring wke",
			in:   "pwwkew",
			size: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := Substring(tt.in)

			if ss != tt.size {
				t.Fatalf("Longest substring expected=%d. got=%d", tt.size, ss)
			}
		})
	}
}
