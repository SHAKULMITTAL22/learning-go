// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

package Factorial

import (
	"testing"
)

func TestFactorialRecursive_178a7b8974(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{
			name: "Test Case 1: Factorial of 0",
			num:  0,
			want: 1,
		},
		{
			name: "Test Case 2: Factorial of 5",
			num:  5,
			want: 120,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FactorialRecursiveTest(tt.num); got != tt.want {
				t.Errorf("FactorialRecursiveTest() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Success: FactorialRecursiveTest(%v) = %v", tt.num, got)
			}
		})
	}
}

func FactorialRecursiveTest(num int) int {
	if num == 0 {
		return 1
	}

	return num * FactorialRecursiveTest(num-1)
}
