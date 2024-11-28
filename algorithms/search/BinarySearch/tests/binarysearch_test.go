package BinarySearch

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

/*
ROOST_METHOD_HASH=binarySearch_5149337a0e
ROOST_METHOD_SIG_HASH=binarySearch_7d22ad2576


 */
func Testbinarysearch(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
	}{
		{
			name:     "Query Exists in Array",
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			query:    4,
			expected: 3,
		},
		{
			name:     "Query at the Last Position",
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			query:    7,
			expected: 6,
		},
		{
			name:     "Query at the First Position",
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			query:    1,
			expected: 0,
		},
		{
			name:     "Query Does Not Exist",
			arr:      []int{1, 2, 3, 4, 5, 6, 7},
			query:    10,
			expected: -1,
		},
		{
			name:     "Empty Array",
			arr:      []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "Single Element Matching Query",
			arr:      []int{1},
			query:    1,
			expected: 0,
		},
		{
			name:     "Single Element Not Matching Query",
			arr:      []int{1},
			query:    2,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := binarySearch(tt.arr, tt.query)

			assert.Equal(t, tt.expected, result)

			t.Logf("Test case '%s': expected index %d, got %d", tt.name, tt.expected, result)
		})
	}
}

/*
ROOST_METHOD_HASH=binarySearch1_e1c20abb75
ROOST_METHOD_SIG_HASH=binarySearch1_c36ccb5b2b


 */
// Testbinarysearch1 is designed to rigorously test the binarySearch1 function using a table-driven approach.
// This approach enhances robustness by efficiently accommodating a variety of scenarios, as specified in the instructions.
func Testbinarysearch1(t *testing.T) {
	// Define a struct to encapsulate test scenarios and expected outcomes
	type test struct {
		name     string
		arr      []int
		query    int
		expected int
	}

	tests := []test{

		{
			name:     "Element exists in the middle",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:    5,
			expected: 4,
		},

		{
			name:     "Element not present in array",
			arr:      []int{1, 2, 3, 4, 6, 7, 8, 9, 10},
			query:    5,
			expected: -1,
		},

		{
			name:     "Element at the start",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:    1,
			expected: 0,
		},

		{
			name:     "Element at the end",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:    10,
			expected: 9,
		},

		{
			name:     "Empty array",
			arr:      []int{},
			query:    5,
			expected: -1,
		},

		{
			name:     "Single-element array, element present",
			arr:      []int{5},
			query:    5,
			expected: 0,
		},

		{
			name:     "Single-element array, element absent",
			arr:      []int{3},
			query:    5,
			expected: -1,
		},

		{
			name:     "Duplicate elements, element present",
			arr:      []int{1, 2, 3, 4, 4, 4, 5, 6, 7, 8, 9, 10},
			query:    4,
			expected: 3,
		},

		{
			name:     "Reverse-sorted array",
			arr:      []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			query:    5,
			expected: -1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := binarySearch1(tc.arr, tc.query)
			if result != tc.expected {
				t.Logf("Failed: %s - Expected index: %d, got: %d", tc.name, tc.expected, result)
				t.Fail()
			} else {
				t.Logf("Success: %s - Correctly found element at index: %d", tc.name, result)
			}
		})
	}
}

