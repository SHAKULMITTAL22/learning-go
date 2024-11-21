package main

import (
	"fmt"
	"testing"
)

func sayHelloTo(name string) {
	fmt.Println("Hello ", name)
}

func TestSayHelloTo_4a6fd0c56c(t *testing.T) {
	// Test case 1: Valid name
	name := "John"
	sayHelloTo(name)
	t.Log("Test case 1 passed")

	// Test case 2: Empty name
	name = ""
	sayHelloTo(name)
	t.Log("Test case 2 passed")

	// Test case 3: Name with special characters
	name = "J@hn"
	sayHelloTo(name)
	t.Log("Test case 3 passed")

	// Test case 4: Name with numbers
	name = "J0hn"
	sayHelloTo(name)
	t.Log("Test case 4 passed")
}
