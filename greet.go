package main

import "fmt"

// Greet returns a greeting containing the provided name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
