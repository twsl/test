package main

import (
	"strings"
	"testing"
)

func TestGreetIncludesName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "simple name", in: "Alice", want: "Hello, Alice!"},
		{name: "empty name", in: "", want: "Hello, !"},
		{name: "unicode name", in: "Zoë", want: "Hello, Zoë!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.in)
			if got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.in, got, tt.want)
			}
			if !strings.Contains(got, tt.in) {
				t.Errorf("Greet(%q) = %q, does not contain name", tt.in, got)
			}
		})
	}
}
