package main

import (
	"fmt"
	"math"
)

// SumDigitsIterative calculates the sum of digits in a number using iteration
func SumDigitsIterative(num int) int {
	// Handle negative numbers
	num = int(math.Abs(float64(num)))

	sum := 0
	for num > 0 {
		sum += num % 10 // Get the last digit
		num /= 10       // Remove the last digit
	}
	return sum
}

// SumDigitsRecursive calculates the sum of digits in a number using recursion
func SumDigitsRecursive(num int) int {
	// Handle negative numbers
	num = int(math.Abs(float64(num)))

	// Base case: single digit number
	if num < 10 {
		return num
	}

	// Recursive case: last digit + sum of remaining digits
	return num%10 + SumDigitsRecursive(num/10)
}

func main() {
	// Test cases for digit sum
	testCases := []int{
		123,    // 1 + 2 + 3 = 6
		456,    // 4 + 5 + 6 = 15
		789,    // 7 + 8 + 9 = 24
		0,      // 0
		1,      // 1
		-123,   // 1 + 2 + 3 = 6 (handles negative numbers)
		999999, // 9 + 9 + 9 + 9 + 9 + 9 = 54
	}

	fmt.Println("Sum of Digits Results:")
	fmt.Println("--------------------")
	for _, num := range testCases {
		iterativeSum := SumDigitsIterative(num)
		recursiveSum := SumDigitsRecursive(num)

		fmt.Printf("Number: %d\n", num)
		fmt.Printf("Iterative Sum: %d\n", iterativeSum)
		fmt.Printf("Recursive Sum: %d\n", recursiveSum)
		fmt.Println()
	}
}
