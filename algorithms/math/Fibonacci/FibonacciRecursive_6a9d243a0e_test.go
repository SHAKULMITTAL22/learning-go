// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

package Fibonacci

import (
	"testing"
)

func TestFibonacciRecursive_6a9d243a0e(t *testing.T) {
	t.Run("Test with positive number", func(t *testing.T) {
		result := FibonacciRecursive(6)
		if result != 8 {
			t.Errorf("FibonacciRecursive(6) = %d; want 8", result)
		} else {
			t.Logf("Success: FibonacciRecursive(6) = %d", result)
		}
	})

	t.Run("Test with zero", func(t *testing.T) {
		result := FibonacciRecursive(0)
		if result != 0 {
			t.Errorf("FibonacciRecursive(0) = %d; want 0", result)
		} else {
			t.Logf("Success: FibonacciRecursive(0) = %d", result)
		}
	})

	t.Run("Test with negative number", func(t *testing.T) {
		result := FibonacciRecursive(-4)
		if result != -4 {
			t.Errorf("FibonacciRecursive(-4) = %d; want -4", result)
		} else {
			t.Logf("Success: FibonacciRecursive(-4) = %d", result)
		}
	})
}
