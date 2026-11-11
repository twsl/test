package main

import (
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "simple name", input: "Go", expected: "Hello, Go!"},
		{name: "empty name", input: "", expected: "Hello, !"},
		{name: "name with spaces", input: "Ada Lovelace", expected: "Hello, Ada Lovelace!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Greet(tt.input)
			if got != tt.expected {
				t.Errorf("Greet(%q) = %q, want %q", tt.input, got, tt.expected)
			}
		})
	}
}

func TestGreetContainsName(t *testing.T) {
	const name = "Claude"
	got := Greet(name)
	if !strings.Contains(got, name) {
		t.Errorf("Greet(%q) = %q, want it to contain %q", name, got, name)
	}
}
