package main

import "fmt"

// FibonacciGenerator returns a function that generates the next Fibonacci number
func FibonacciGenerator() func() int {
	// Initialize the first two numbers in the sequence
	prev, curr := 0, 1

	// Return a closure that maintains state between calls
	return func() int {
		// Calculate the next number
		next := prev + curr
		// Update the state for the next call
		prev, curr = curr, next
		// Return the current number
		return prev
	}
}

func main() {
	// Create a new Fibonacci generator
	fib := FibonacciGenerator()

	// Generate and print the first 10 Fibonacci numbers
	fmt.Println("First 10 Fibonacci numbers:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fib())
	}
	fmt.Println()
}
