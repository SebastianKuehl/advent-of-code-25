package main

import (
	"fmt"
)

// Example algorithm: find the n-th Fibonacci number
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println("Fibonacci(10):", Fibonacci(10))
}
