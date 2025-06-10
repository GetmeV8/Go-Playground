package main

import "fmt"

// IsPalindrome checks if a string is a palindrome
// Parameters:
//   - s: the string to check
//   - caseSensitive: if true, the check will be case-sensitive
//
// Returns true if the string is a palindrome, false otherwise
func IsPalindrome(s string, caseSensitive bool) bool {
	// Convert string to rune slice to handle Unicode characters correctly
	runes := []rune(s)
	n := len(runes)

	// If case-insensitive, convert to lowercase
	if !caseSensitive {
		for i := 0; i < n; i++ {
			if runes[i] >= 'A' && runes[i] <= 'Z' {
				runes[i] += 32 // Convert uppercase to lowercase
			}
		}
	}

	// Check characters from both ends
	for i := 0; i < n/2; i++ {
		if runes[i] != runes[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	// Test cases for palindrome checker
	testCases := []struct {
		str           string
		caseSensitive bool
	}{
		{"racecar", true},
		{"Racecar", false},
		{"hello", true},
		{"A man, a plan, a canal: Panama", false},
		{"Was it a car or a cat I saw?", false},
		{"12321", true},
		{"上海自来水来自海上", true}, // Chinese palindrome
	}

	fmt.Println("Palindrome Checker Results:")
	fmt.Println("-------------------------")
	for _, tc := range testCases {
		result := IsPalindrome(tc.str, tc.caseSensitive)
		fmt.Printf("String: %q\nCase Sensitive: %v\nIs Palindrome: %v\n\n",
			tc.str, tc.caseSensitive, result)
	}
}
