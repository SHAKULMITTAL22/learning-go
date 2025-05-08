package Fibonacci

import (
	"testing"
)

func fibonacci(n int) int {
	if n <= 0 {
		if n == 0 {
			return 0
		}
		panic("Input should not be a negative integer")
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func TestFibonacci_3c2494e9fa(t *testing.T) {
	t.Run("Test with positive integer", func(t *testing.T) {
		result := fibonacci(10)
		if result != 55 {
			t.Errorf("Expected 55, but got %d", result)
		} else {
			t.Logf("Success: Expected 55 and got %d", result)
		}
	})

	t.Run("Test with zero", func(t *testing.T) {
		result := fibonacci(0)
		if result != 0 {
			t.Errorf("Expected 0, but got %d", result)
		} else {
			t.Logf("Success: Expected 0 and got %d", result)
		}
	})

	t.Run("Test with negative integer", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			} else {
				t.Logf("Success: The code panicked as expected")
			}
		}()
		fibonacci(-1)
	})
}
