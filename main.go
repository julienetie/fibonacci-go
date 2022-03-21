package main

import (
	"fmt"
)

// Recursive method
func Recurse(n int) int {
	if n <= 1 {
		return n
	}
	return Recurse(n-1) + Recurse(n-2)
}

func Prompt() {
	fmt.Println("Provide the Fibonacci length")
	var fibonacciLength int
	fmt.Scanln(&fibonacciLength)
	fmt.Println(Recurse(fibonacciLength))
	Prompt()
}

func main() {

	Prompt()
}
