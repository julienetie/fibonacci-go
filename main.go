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

func Loop(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	var c int

	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

func Prompt() {
	fmt.Println("Provide the Fibonacci length")
	var fibonacciLength int
	fmt.Scanln(&fibonacciLength)
	// fmt.Println(Recurse(fibonacciLength))
	fmt.Println(Loop(fibonacciLength))
	Prompt()
}

func main() {

	Prompt()
}
