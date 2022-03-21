package main

import (
	"fmt"
)

// Recursive method. (This can be very slow)
func Recurse(n int) int {
	if n <= 1 {
		return n
	}
	return Recurse(n-1) + Recurse(n-2)
}

// Loop method.
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

// Array Method.
func Array(n int) int {
	if n <= 1 {
		return n
	}
	arr := make([]int, n)
	arr[0], arr[1] = 1, 1

	for i, _ := range arr {
		if i < 2 {
			continue
		}
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[len(arr)-1]
}

// Slice Method.
func Slice(n int) int {
	arr := make([]int, n+1, n+2)
	if n < 2 {
		arr = arr[0:2]
	}

	arr[0], arr[1] = 0, 1

	for i := 2; i <= n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}

func Prompt(method string) {
	var fibonacciLength int
	var result int
	var methodName string

	switch method {
	case "r":
		methodName = "Recursive"
	case "a":
		methodName = "Array"
	case "l":
		methodName = "Loop"
	case "s":
		methodName = "Slice"
	}
	fmt.Printf(`Provide the %v Fibonacci length: `, methodName)
	fmt.Scanln(&fibonacciLength)

	switch method {
	case "r":
		result = Recurse(fibonacciLength)
	case "a":
		result = Array(fibonacciLength)
	case "l":
		result = Loop(fibonacciLength)
	case "s":
		result = Slice(fibonacciLength)
	}

	fmt.Println(result)
	Prompt(method)
}

func main() {
	var fibonacciMethod string
	fmt.Printf(`

╔═╗┬┌┐ ┌─┐┌┐┌┌─┐┌─┐┌─┐┬
╠╣ │├┴┐│ ││││├─┤│  │  │
╚  ┴└─┘└─┘┘└┘┴ ┴└─┘└─┘┴
	
Choose a Fibonacci method

- Recursive 			(r)
- Loop				(l)
- Array				(a)
- Better Array			(s)
(Type r/l/a/s): `)

	fmt.Scanln(&fibonacciMethod)
	Prompt(fibonacciMethod)
}
