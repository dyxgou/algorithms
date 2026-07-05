package main

import (
	"testing"
)

func TestReverseSimpleString(t *testing.T) {
	tests := []struct {
		name            string
		input, expected string
	}{
		{
			name:     "First Simple test",
			input:    "ab-cd",
			expected: "dc-ba",
		},
		{
			name:     "Complex test",
			input:    "a-bC-dEf=ghlj!!",
			expected: "j-lh-gfE=dCba!!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Reverser(tt.input)

			if rev := r.Reverse(); tt.expected != rev {
				t.Errorf("reverser input=%q. expected=%q. got=%q", tt.input, tt.expected, rev)
			}
		})
	}
}
