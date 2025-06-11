package main

import (
	"errors"
	"fmt"
)

// Stack represents a stack data structure
type Stack struct {
	items []int // slice to store stack elements
}


// NewStack creates and returns a new empty stack
func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
// Returns an error if the stack is empty
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	
	index := len(s.items) - 1 // Index of the last element
	item := s.items[index]    // Get the top element
	s.items = s.items[:index] // Remove the last element
	
	return item, nil
}

// Peek returns the top element without removing it
// Returns an error if the stack is empty
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return len(s.items)
}

// Display prints all elements in the stack (top to bottom)
func (s *Stack) Display() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	
	fmt.Print("Stack (top to bottom): ")
	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Printf("%d ", s.items[i])
	}
	fmt.Println()
}

// Example usage
func main() {
	// Create a new stack
	stack := NewStack()
	
	fmt.Println("=== Stack Operations Demo ===")
	
	// Check if empty
	fmt.Printf("Is stack empty? %t\n", stack.IsEmpty())
	fmt.Printf("Stack size: %d\n", stack.Size())
	
	// Try to peek empty stack
	if _, err := stack.Peek(); err != nil {
		fmt.Printf("Peek error: %s\n", err)
	}
	
	// Try to pop empty stack
	if _, err := stack.Pop(); err != nil {
		fmt.Printf("Pop error: %s\n", err)
	}
	
	fmt.Println("\n--- Pushing elements ---")
	// Push some elements
	stack.Push(10)
	fmt.Printf("Pushed: 10, Size: %d\n", stack.Size())
	stack.Display()
	
	stack.Push(20)
	fmt.Printf("Pushed: 20, Size: %d\n", stack.Size())
	stack.Display()
	
	stack.Push(30)
	fmt.Printf("Pushed: 30, Size: %d\n", stack.Size())
	stack.Display()
	
	stack.Push(40)
	fmt.Printf("Pushed: 40, Size: %d\n", stack.Size())
	stack.Display()
	
	fmt.Println("\n--- Peek operations ---")
	// Peek at top element
	if top, err := stack.Peek(); err == nil {
		fmt.Printf("Top element (peek): %d\n", top)
	}
	fmt.Printf("Size after peek: %d\n", stack.Size())
	stack.Display()
	
	fmt.Println("\n--- Pop operations ---")
	// Pop elements
	if item, err := stack.Pop(); err == nil {
		fmt.Printf("Popped: %d, Size: %d\n", item, stack.Size())
		stack.Display()
	}
	
	if item, err := stack.Pop(); err == nil {
		fmt.Printf("Popped: %d, Size: %d\n", item, stack.Size())
		stack.Display()
	}
	
	if item, err := stack.Pop(); err == nil {
		fmt.Printf("Popped: %d, Size: %d\n", item, stack.Size())
		stack.Display()
	}
	
	if item, err := stack.Pop(); err == nil {
		fmt.Printf("Popped: %d, Size: %d\n", item, stack.Size())
		stack.Display()
	}
	
	// Try to pop from empty stack
	if _, err := stack.Pop(); err != nil {
		fmt.Printf("Pop error: %s\n", err)
	}
	
	fmt.Printf("Final stack size: %d\n", stack.Size())
	fmt.Printf("Is stack empty? %t\n", stack.IsEmpty())
}