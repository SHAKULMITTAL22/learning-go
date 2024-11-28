package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd
*/
func Testinterpolationsearch523(t *testing.T) {
	type testCase struct {
		name     string
		arr      []int
		query    int
		expected int
	}

	testCases := []testCase{
		{
			name:     "Scenario 1: Element exists in sorted array",
			arr:      []int{1, 3, 5, 7, 9, 11, 13},
			query:    7,
			expected: 3,
		},
		{
			name:     "Scenario 2: Element does not exist in sorted array",
			arr:      []int{1, 3, 5, 7, 9, 11, 13},
			query:    8,
			expected: -1,
		},
		{
			name:     "Scenario 3: Empty array",
			arr:      []int{},
			query:    5,
			expected: -1,
		},
		{
			name:     "Scenario 4: Identical elements where query is present",
			arr:      []int{2, 2, 2, 2, 2},
			query:    2,
			expected: 0,
		},
		{
			name:     "Scenario 5: Single element matching query",
			arr:      []int{4},
			query:    4,
			expected: 0,
		},
		{
			name:     "Scenario 6: Query larger than any element",
			arr:      []int{1, 3, 5, 7, 9},
			query:    10,
			expected: -1,
		},
		{
			name:     "Scenario 7: Query smaller than any element",
			arr:      []int{2, 3, 4, 5, 6},
			query:    1,
			expected: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := interpolationSearch(tc.arr, tc.query)
			if actual != tc.expected {
				t.Errorf("Failed %s: expected %d but got %d", tc.name, tc.expected, actual)
			} else {
				t.Logf("Passed %s: found element at index %d", tc.name, actual)
			}
		})
	}
}
