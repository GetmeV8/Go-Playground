package main

import (
	"fmt"
	"math"
)

// KthSmallest finds the kth smallest element using simple selection approach
func KthSmallest(arr []int, k int) (int, error) {
	// Input validation
	if k <= 0 || k > len(arr) {
		return 0, fmt.Errorf("k must be between 1 and %d", len(arr))
	}
	
	if len(arr) == 0 {
		return 0, fmt.Errorf("array cannot be empty")
	}
	
	// Create a copy of the array to avoid modifying the original
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	
	// Create a boolean slice to mark used elements
	used := make([]bool, len(arrCopy))
	
	var kthSmallest int
	
	// Find the smallest element K times
	for i := 0; i < k; i++ {
		minVal := math.MaxInt32
		minIndex := -1
		
		// Find the minimum element that hasn't been used yet
		for j := 0; j < len(arrCopy); j++ {
			if !used[j] && arrCopy[j] < minVal {
				minVal = arrCopy[j]
				minIndex = j
			}
		}
		
		// Mark this element as used
		used[minIndex] = true
		kthSmallest = minVal
		
		fmt.Printf("Step %d: Found %d smallest = %d\n", i+1, i+1, minVal)
	}
	
	return kthSmallest, nil
}

// Alternative implementation: Remove elements instead of marking them
func KthSmallestRemoval(arr []int, k int) (int, error) {
	// Input validation
	if k <= 0 || k > len(arr) {
		return 0, fmt.Errorf("k must be between 1 and %d", len(arr))
	}
	
	if len(arr) == 0 {
		return 0, fmt.Errorf("array cannot be empty")
	}
	
	// Create a copy of the array
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	
	var kthSmallest int
	
	// Find and remove the smallest element K times
	for i := 0; i < k; i++ {
		// Find minimum element and its index
		minVal := arrCopy[0]
		minIndex := 0
		
		for j := 1; j < len(arrCopy); j++ {
			if arrCopy[j] < minVal {
				minVal = arrCopy[j]
				minIndex = j
			}
		}
		
		kthSmallest = minVal
		fmt.Printf("Step %d: Found %d smallest = %d\n", i+1, i+1, minVal)
		
		// Remove the minimum element from the array
		arrCopy = append(arrCopy[:minIndex], arrCopy[minIndex+1:]...)
	}
	
	return kthSmallest, nil
}

// Helper function to find minimum in a slice
func findMin(arr []int, used []bool) (int, int) {
	minVal := math.MaxInt32
	minIndex := -1
	
	for i, val := range arr {
		if !used[i] && val < minVal {
			minVal = val
			minIndex = i
		}
	}
	
	return minVal, minIndex
}

// KthSmallestClean - cleaner version using helper function
func KthSmallestClean(arr []int, k int) (int, error) {
	// Input validation
	if k <= 0 || k > len(arr) {
		return 0, fmt.Errorf("k must be between 1 and %d", len(arr))
	}
	
	if len(arr) == 0 {
		return 0, fmt.Errorf("array cannot be empty")
	}
	
	// Create a copy to avoid modifying original
	arrCopy := make([]int, len(arr))
	copy(arrCopy, arr)
	
	used := make([]bool, len(arrCopy))
	var result int
	
	// Find kth smallest element
	for i := 0; i < k; i++ {
		minVal, minIndex := findMin(arrCopy, used)
		used[minIndex] = true
		result = minVal
	}
	
	return result, nil
}

func main() {
	fmt.Println("=== Kth Smallest Element - Simple Selection Approach ===\n")
	
	// Test cases
	testCases := []struct {
		arr []int
		k   int
		desc string
	}{
		{[]int{7, 10, 4, 3, 20, 15}, 3, "Basic case"},
		{[]int{1, 2, 3, 4, 5}, 3, "Already sorted"},
		{[]int{5, 4, 3, 2, 1}, 2, "Reverse sorted"},
		{[]int{4, 4, 4, 4}, 2, "All duplicates"},
		{[]int{1, 3, 2, 4, 5, 7, 6}, 4, "Random order"},
		{[]int{42}, 1, "Single element"},
		{[]int{9, 1, 8, 2, 7, 3, 6, 4, 5}, 5, "Large case"},
	}
	
	for i, tc := range testCases {
		fmt.Printf("Test Case %d: %s\n", i+1, tc.desc)
		fmt.Printf("Array: %v, K: %d\n", tc.arr, tc.k)
		
		// Test with marking approach
		result, err := KthSmallest(tc.arr, tc.k)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("Result: %d smallest element is %d\n", tc.k, result)
		}
		
		fmt.Println(stringss.Repeat("-", 50))
	}
	
	// Test error cases
	fmt.Println("\n=== Error Cases ===")
	
	errorCases := []struct {
		arr []int
		k   int
		desc string
	}{
		{[]int{1, 2, 3}, 0, "K = 0"},
		{[]int{1, 2, 3}, 5, "K > array length"},
		{[]int{}, 1, "Empty array"},
		{[]int{1, 2, 3}, -1, "Negative K"},
	}
	
	for _, tc := range errorCases {
		fmt.Printf("Test: %s - Array: %v, K: %d\n", tc.desc, tc.arr, tc.k)
		_, err := KthSmallest(tc.arr, tc.k)
		if err != nil {
			fmt.Printf("✓ Correctly caught error: %s\n", err)
		} else {
			fmt.Printf("✗ Should have thrown an error\n")
		}
		fmt.Println()
	}
	
	// Compare different implementations
	fmt.Println("=== Comparing Different Implementations ===")
	testArr := []int{64, 34, 25, 12, 22, 11, 90}
	k := 4
	
	fmt.Printf("Array: %v, Finding %d smallest\n\n", testArr, k)
	
	fmt.Println("Method 1 - Marking approach:")
	result1, _ := KthSmallest(testArr, k)
	fmt.Printf("Result: %d\n\n", result1)
	
	fmt.Println("Method 2 - Removal approach:")
	result2, _ := KthSmallestRemoval(testArr, k)
	fmt.Printf("Result: %d\n\n", result2)
	
	fmt.Println("Method 3 - Clean approach:")
	result3, _ := KthSmallestClean(testArr, k)
	fmt.Printf("Result: %d\n", result3)
}