package BinarySearch

import "testing"

/*
ROOST_METHOD_HASH=binarySearch_5149337a0e
ROOST_METHOD_SIG_HASH=binarySearch_7d22ad2576


 */
func Testbinarysearch192(t *testing.T) {
	type testCase struct {
		name     string
		arr      []int
		query    int
		expected int
	}

	testCases := []testCase{
		{
			name:     "Element present in the middle",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			query:    5,
			expected: 4,
		},
		{
			name:     "First element in the array",
			arr:      []int{1, 2, 3, 4, 5},
			query:    1,
			expected: 0,
		},
		{
			name:     "Last element in the array",
			arr:      []int{1, 2, 3, 4, 5},
			query:    5,
			expected: 4,
		},
		{
			name:     "Element not present in the array",
			arr:      []int{1, 2, 3, 4, 5},
			query:    6,
			expected: -1,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "Non-integer element query",
			arr:      []int{1, 2, 3, 4, 5},
			query:    'a',
			expected: -1,
		},
		{
			name:     "Single-element array",
			arr:      []int{1},
			query:    1,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := binarySearch(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed", tc.name)
			}
		})
	}
}

