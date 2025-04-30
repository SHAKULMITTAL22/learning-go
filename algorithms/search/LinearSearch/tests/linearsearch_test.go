package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721
*/
func Testlinearsearch367(t *testing.T) {
	type testCase struct {
		description string
		array       []int
		query       int
		expected    int
	}

	testCases := []testCase{
		{
			description: "Non-Existent Element",
			array:       []int{1, 2, 3, 4, 5},
			query:       6,
			expected:    -1,
		},
		{
			description: "Empty Array",
			array:       []int{},
			query:       1,
			expected:    -1,
		},
		{
			description: "Multiple Identical Elements",
			array:       []int{9, 3, 9, 3, 1, 9},
			query:       9,
			expected:    0,
		},
		{
			description: "Sorted Array",
			array:       []int{1, 2, 3, 4, 5},
			query:       3,
			expected:    2,
		},
		{
			description: "Array with Negative Numbers",
			array:       []int{-1, -2, -3, -4, 0, 1, 2},
			query:       -3,
			expected:    2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			t.Logf("Testing case: %s", tc.description)
			result := linearSearch(tc.array, tc.query)
			if result != tc.expected {
				t.Errorf("Failed %s: expected %d, got %d", tc.description, tc.expected, result)
			} else {
				t.Logf("Passed %s: expected %d, got %d", tc.description, tc.expected, result)
			}
		})
	}

}
