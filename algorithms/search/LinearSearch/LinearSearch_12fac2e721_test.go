// Test generated by RoostGPT for test math-go using AI Type Open AI and AI Model gpt-4

package LinearSearch

import (
	"testing"
)

func TestLinearSearch_12fac2e721(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		query  int
		expect int
	}{
		{
			name:   "Test Case 1: Element found in the array",
			arr:    []int{1, 2, 3, 4, 5},
			query:  4,
			expect: 3,
		},
		{
			name:   "Test Case 2: Element not found in the array",
			arr:    []int{1, 2, 3, 4, 5},
			query:  6,
			expect: -1,
		},
		{
			name:   "Test Case 3: Query on an empty array",
			arr:    []int{},
			query:  1,
			expect: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := linearSearchTestFunc(tc.arr, tc.query)
			if result != tc.expect {
				t.Errorf("linearSearch(%v, %v): expected %v, got %v", tc.arr, tc.query, tc.expect, result)
			} else {
				t.Logf("Success: %s", tc.name)
			}
		})
	}
}

func linearSearchTestFunc(arr []int, query int) int {
	for i, val := range arr {
		if val == query {
			return i
		}
	}
	return -1
}
