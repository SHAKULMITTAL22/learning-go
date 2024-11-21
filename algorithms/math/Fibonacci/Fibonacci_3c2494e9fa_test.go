package Fibonacci

import (
	"reflect"
	"testing"
)

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{0}
	}

	fibSeq := make([]int, n+1)
	fibSeq[0], fibSeq[1] = 0, 1

	for i := 2; i <= n; i++ {
		fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
	}

	return fibSeq
}

func TestFibonacciSequence_bf4aa71a9c(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []int
	}{
		{
			name: "Test Case 1",
			args: 5,
			want: []int{0, 1, 1, 2, 3, 5},
		},
		{
			name: "Test Case 2",
			args: 10,
			want: []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
