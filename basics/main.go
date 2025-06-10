package main

import "fmt"

func main() {
	testString := "Hello, 世界!"
	reversed := ReverseString(testString)
	fmt.Printf("Original: %s\nReversed: %s\n", testString, reversed)
}
