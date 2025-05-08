package Fibonacci

import (
	"testing"
)

func fibonacci(n int) int {
	if n < 0 {
		panic("Negative arguments not allowed")
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func TestFibonacci_3c2494e9fa(t *testing.T) {
	t.Run("Test with positive integer", func(t *testing.T) {
		got := fibonacci(5)
		want := 5
		if got != want {
			t.Errorf("fibonacci(5) = %d; want %d", got, want)
		} else {
			t.Logf("Success: fibonacci(5) = %d", got)
		}
	})

	t.Run("Test with zero", func(t *testing.T) {
		got := fibonacci(0)
		want := 0
		if got != want {
			t.Errorf("fibonacci(0) = %d; want %d", got, want)
		} else {
			t.Logf("Success: fibonacci(0) = %d", got)
		}
	})

	t.Run("Test with negative integer", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()
		fibonacci(-1)
	})
}
