package main

import "fmt"

// ReverseString takes a string as input and returns its reverse
func ReverseString(s string) string {
	// Convert string to rune slice to handle Unicode characters correctly
	runes := []rune(s)
	
	// Get the length of the rune slice
	n := len(runes)
	
	// Swap characters from both ends until we reach the middle
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	
	// Convert rune slice back to string
	return string(runes)
}

func main() {
	// Test cases for string reversal
	testCases := []string{
		"Hello, World!",
		"上海自来水来自海上", // Chinese characters
		"12345",
		"",
		"A",
		"racecar",
	}

	fmt.Println("String Reversal Results:")
	fmt.Println("----------------------")
	for _, test := range testCases {
		reversed := ReverseString(test)
		fmt.Printf("Original: %q\nReversed: %q\n\n", test, reversed)
	}
}
