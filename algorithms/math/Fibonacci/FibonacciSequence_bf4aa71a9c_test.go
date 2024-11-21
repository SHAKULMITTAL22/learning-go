package Fibonacci

import (
	"testing"
	"reflect"
)

// fibonacciTestHelper is a helper function for generating fibonacci number at a particular position
func fibonacciTestHelper(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciTestHelper(n-1) + fibonacciTestHelper(n-2)
}

// TestFibonacciSequence_bf4aa71a9c is a test function for fibonacciSequence
func TestFibonacciSequence_bf4aa71a9c(t *testing.T) {
	// Test case 1: Testing for num = 5
	expectedResult := []int{0, 1, 1, 2, 3, 5}
	result := fibonacciSequence(5)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Test Case 1 Failed: Expected ", expectedResult, " but got ", result)
	} else {
		t.Log("Test Case 1 Passed")
	}

	// Test case 2: Testing for num = 0
	expectedResult = []int{0}
	result = fibonacciSequence(0)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Test Case 2 Failed: Expected ", expectedResult, " but got ", result)
	} else {
		t.Log("Test Case 2 Passed")
	}

	// Test case 3: Testing for num = 1
	expectedResult = []int{0, 1}
	result = fibonacciSequence(1)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Test Case 3 Failed: Expected ", expectedResult, " but got ", result)
	} else {
		t.Log("Test Case 3 Passed")
	}

	// Test case 4: Testing for num = 2
	expectedResult = []int{0, 1, 1}
	result = fibonacciSequence(2)
	if !reflect.DeepEqual(result, expectedResult) {
		t.Error("Test Case 4 Failed: Expected ", expectedResult, " but got ", result)
	} else {
		t.Log("Test Case 4 Passed")
	}
}
