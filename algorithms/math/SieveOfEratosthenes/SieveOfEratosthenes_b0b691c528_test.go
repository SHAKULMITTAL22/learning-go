package SieveOfEratosthenes

import (
	"reflect"
	"testing"
)

func sieveOfEratosthenes(n int) []int {
	if n < 2 {
		return []int{}
	}
	numbers := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if numbers[i] == false {
			for j := i * i; j <= n; j += i {
				numbers[j] = true
			}
		}
	}
	var primes []int
	for i := 2; i <= n; i++ {
		if numbers[i] == false {
			primes = append(primes, i)
		}
	}
	return primes
}

func TestSieveOfEratosthenes_b0b691c528(t *testing.T) {
	testCases := []struct {
		maxNumber int
		expected  []int
	}{
		{
			maxNumber: 10,
			expected:  []int{2, 3, 5, 7},
		},
		{
			maxNumber: 0,
			expected:  []int{},
		},
		{
			maxNumber: 1,
			expected:  []int{},
		},
		{
			maxNumber: 2,
			expected:  []int{2},
		},
		{
			maxNumber: 30,
			expected:  []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29},
		},
	}

	for i, testCase := range testCases {
		result := sieveOfEratosthenes(testCase.maxNumber)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Test case %d failed, expected: %v, got: %v", i, testCase.expected, result)
		} else {
			t.Logf("Test case %d passed", i)
		}
	}
}
