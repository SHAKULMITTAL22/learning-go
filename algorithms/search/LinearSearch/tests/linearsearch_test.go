package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721


 */
func Testlinearsearch403(t *testing.T) {

	type testCase struct {
		name     string
		array    []int
		query    int
		expected int
	}

	testCases := []testCase{
		{
			name:     "Scenario 1: Test linearSearch with a Non-Existent Element",
			array:    []int{1, 2, 3, 4, 5},
			query:    6,
			expected: -1,
		},
		{
			name:     "Scenario 2: Test linearSearch with an Empty Array",
			array:    []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "Scenario 3: Test linearSearch with Multiple Identical Elements",
			array:    []int{1, 2, 3, 2, 1},
			query:    2,
			expected: 1,
		},
		{
			name:     "Scenario 4: Test linearSearch with a Sorted Array",
			array:    []int{1, 2, 3, 4, 5},
			query:    4,
			expected: 3,
		},
		{
			name:     "Scenario 5: Test linearSearch with Negative Numbers",
			array:    []int{-5, -4, -3, 0, 3, 4, 5},
			query:    -3,
			expected: 2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tc.name)

			result := linearSearch(tc.array, tc.query)

			if result != tc.expected {
				t.Errorf("Failed %s: expected %d, got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Passed %s: expected %d, got %d", tc.name, tc.expected, result)
			}

			t.Logf("Array: %v, Query: %d, Expected: %d, Result: %d", tc.array, tc.query, tc.expected, result)
		})
	}
}

