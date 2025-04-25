package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721
*/
func Testlinearsearch34(t *testing.T) {
	type test struct {
		name     string
		arr      []int
		query    int
		expected int
	}

	tests := []test{
		{
			name:     "Non-Existent Element",
			arr:      []int{2, 4, 6, 8, 10},
			query:    5,
			expected: -1,
		},
		{
			name:     "Empty Array",
			arr:      []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "Multiple Identical Elements",
			arr:      []int{1, 3, 1, 7, 1, 8},
			query:    1,
			expected: 0,
		},
		{
			name:     "Sorted Array",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			query:    4,
			expected: 3,
		},
		{
			name:     "Array with Negative Numbers",
			arr:      []int{-3, -1, 0, 2, 3},
			query:    -1,
			expected: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := linearSearch(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("linearSearch(%v, %d) = %d; want %d", tc.arr, tc.query, result, tc.expected)
			} else {
				t.Logf("Success: %s - linearSearch(%v, %d) returned %d as expected", tc.name, tc.arr, tc.query, result)
			}
		})
	}
}
