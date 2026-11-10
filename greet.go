package main

import "fmt"

// Greet returns a greeting that includes the provided name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
