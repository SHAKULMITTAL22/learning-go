package Fibonacci

import (
	"reflect"
	"testing"
)

func fibonacciTest(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciTest(n-1) + fibonacciTest(n-2)
}

func fibonacciSequenceTest(num int) []int {
	var res = []int{}
	for n := 0; n <= num; n++ {
		result := fibonacciTest(n)
		res = append(res, result)
	}
	return res
}

func TestFibonacciSequence_bf4aa71a9c(t *testing.T) {
	testCases := []struct {
		num      int
		expected []int
	}{
		{5, []int{0, 1, 1, 2, 3, 5}},
		{10, []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55}},
	}

	for _, tc := range testCases {
		result := fibonacciSequenceTest(tc.num)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("fibonacciSequenceTest(%v) => %v, expected %v", tc.num, result, tc.expected)
		} else {
			t.Logf("Success: fibonacciSequenceTest(%v) => %v", tc.num, result)
		}
	}
}
