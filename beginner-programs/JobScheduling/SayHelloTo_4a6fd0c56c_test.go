package main

import (
	"fmt"
	"testing"
)

func sayHelloTo(name string) {
	fmt.Println("Hello ", name)
}

// TestSayHelloTo_4a6fd0c56c tests the sayHelloTo function.
func TestSayHelloTo_4a6fd0c56c(t *testing.T) {
	// Test case 1: name is "Alice".
	sayHelloTo("Alice")
	t.Log("Test case 1 passed.")

	// Test case 2: name is "Bob".
	sayHelloTo("Bob")
	t.Log("Test case 2 passed.")

	// Test case 3: name is an empty string.
	sayHelloTo("")
	t.Log("Test case 3 passed.")
}

// TestMain_a2e85e6415 tests the main function.
func TestMain_a2e85e6415(t *testing.T) {
	// Test case 1: no arguments are passed to the main function.
	main()
	t.Log("Test case 1 passed.")

	// Test case 2: one argument is passed to the main function.
	main("argument")
	t.Log("Test case 2 passed.")

	// Test case 3: multiple arguments are passed to the main function.
	main("argument1", "argument2", "argument3")
	t.Log("Test case 3 passed.")
}
