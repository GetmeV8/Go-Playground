package main

import "fmt"

// MergeSortedSlices merges two sorted slices into a single sorted slice
func MergeSortedSlices(slice1, slice2 []int) []int {
	// Create result slice with combined length
	result := make([]int, len(slice1)+len(slice2))

	// Initialize pointers for both slices
	i, j, k := 0, 0, 0

	// Compare and merge elements from both slices
	for i < len(slice1) && j < len(slice2) {
		if slice1[i] <= slice2[j] {
			result[k] = slice1[i]
			i++
		} else {
			result[k] = slice2[j]
			j++
		}
		k++
	}

	// Copy remaining elements from slice1, if any
	for i < len(slice1) {
		result[k] = slice1[i]
		i++
		k++
	}

	// Copy remaining elements from slice2, if any
	for j < len(slice2) {
		result[k] = slice2[j]
		j++
		k++
	}

	return result
}

func main() {
	// Test cases
	testCases := []struct {
		slice1 []int
		slice2 []int
	}{
		{
			[]int{1, 3, 5, 7},
			[]int{2, 4, 6, 8},
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		},
		{
			[]int{1, 3, 5},
			[]int{2, 4},
		},
		{
			[]int{},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3},
			[]int{},
		},
		{
			[]int{},
			[]int{},
		},
	}

	fmt.Println("Merging Sorted Slices")
	fmt.Println("-------------------")

	for i, tc := range testCases {
		merged := MergeSortedSlices(tc.slice1, tc.slice2)
		fmt.Printf("Test Case %d:\n", i+1)
		fmt.Printf("Slice 1: %v\n", tc.slice1)
		fmt.Printf("Slice 2: %v\n", tc.slice2)
		fmt.Printf("Merged:  %v\n\n", merged)
	}
}
