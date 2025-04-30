package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721


 */
func Testlinearsearch65(t *testing.T) {
	type testCase struct {
		name     string
		array    []int
		query    int
		expected int
	}

	testCases := []testCase{
		{
			name:     "NonExistentElement",
			array:    []int{1, 2, 3, 4, 5},
			query:    6,
			expected: -1,
		},
		{
			name:     "EmptyArray",
			array:    []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "MultipleIdenticalElements",
			array:    []int{1, 2, 2, 3, 4},
			query:    2,
			expected: 1,
		},
		{
			name:     "SortedArray",
			array:    []int{1, 2, 3, 4, 5},
			query:    4,
			expected: 3,
		},
		{
			name:     "NegativeNumbers",
			array:    []int{-5, -2, 0, 1, 3},
			query:    -2,
			expected: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := linearSearch(tc.array, tc.query)
			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %d but got %d. Array: %v, Query: %d", tc.name, tc.expected, result, tc.array, tc.query)
			} else {
				t.Logf("Test %s succeeded. Expected and got %d. Array: %v, Query: %d", tc.name, tc.expected, tc.array, tc.query)
			}
		})
	}
}

