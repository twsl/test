package main

import "fmt"

// Greet returns a greeting addressed to name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func main() {
	fmt.Println(Greet("World"))
}
