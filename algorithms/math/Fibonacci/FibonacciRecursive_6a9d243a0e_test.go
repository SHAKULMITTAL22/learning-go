// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

package Fibonacci

import (
	"testing"
)

func TestFibonacciRecursive_6a9d243a0e(t *testing.T) {
	testCases := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Test Case 1: Fibonacci of 0",
			n:    0,
			want: 0,
		},
		{
			name: "Test Case 2: Fibonacci of 1",
			n:    1,
			want: 1,
		},
		{
			name: "Test Case 3: Fibonacci of 5",
			n:    5,
			want: 5,
		},
		{
			name: "Test Case 4: Fibonacci of 10",
			n:    10,
			want: 55,
		},
		{
			name: "Test Case 5: Fibonacci of negative number",
			n:    -5,
			want: -5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := FibonacciRecursive_6a9d243a0e(tc.n); got != tc.want {
				t.Errorf("FibonacciRecursive_6a9d243a0e(%v) = %v, want %v", tc.n, got, tc.want)
			} else {
				t.Logf("Success: FibonacciRecursive_6a9d243a0e(%v) = %v", tc.n, got)
			}
		})
	}
}

func FibonacciRecursive_6a9d243a0e(n int) int {
	if n <= 1 {
		return n
	}

	return FibonacciRecursive_6a9d243a0e(n-1) + FibonacciRecursive_6a9d243a0e(n-2)
}
