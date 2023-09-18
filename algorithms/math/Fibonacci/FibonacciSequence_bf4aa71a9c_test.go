// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

package Fibonacci

import (
	"testing"
	"reflect"
)

func TestFibonacciSequence_bf4aa71a9c(t *testing.T) {
	// Test case 1: Testing for positive number
	result := fibonacciSequence(5)
	expectedResult := []int{0, 1, 1, 2, 3, 5}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test failed! Args: %v, Want: %v, Got: %v", 5, expectedResult, result)
	} else {
		t.Logf("Test success! Args: %v, Want: %v, Got: %v", 5, expectedResult, result)
	}

	// Test case 2: Testing for zero
	result = fibonacciSequence(0)
	expectedResult = []int{0}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test failed! Args: %v, Want: %v, Got: %v", 0, expectedResult, result)
	} else {
		t.Logf("Test success! Args: %v, Want: %v, Got: %v", 0, expectedResult, result)
	}

	// Test case 3: Testing for negative number
	result = fibonacciSequence(-5)
	expectedResult = []int{}
	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Test failed! Args: %v, Want: %v, Got: %v", -5, expectedResult, result)
	} else {
		t.Logf("Test success! Args: %v, Want: %v, Got: %v", -5, expectedResult, result)
	}
}
