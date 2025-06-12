package main

import (
	"fmt"
)

// Find Duplicates in a Slice
// Write a function to find duplicate elements in a slice of integers or strings.

// Generic function to find duplicates (requires Go 1.18+)
func FindDuplicates[T comparable](items []T) []T {
	seen := make(map[T]bool)
	duplicates := make(map[T]bool)
	result := []T{}

	for _, item := range items {
		if seen[item] {
			if !duplicates[item] {
				result = append(result, item)
				duplicates[item] = true
			}
		} else {
			seen[item] = true
		}
	}

	return result
}

func main() {
	// Example 1: Integer slice
	nums := []int{1, 2, 3, 2, 4, 5, 1, 6, 3}
	dupNums := FindDuplicates(nums)
	fmt.Println("Duplicate integers:", dupNums)

	// Example 2: String slice
	words := []string{"go", "java", "go", "python", "ruby", "java"}
	dupWords := FindDuplicates(words)
	fmt.Println("Duplicate strings:", dupWords)
}
