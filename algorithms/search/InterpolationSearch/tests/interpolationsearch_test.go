package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd
*/
func Testinterpolationsearch298(t *testing.T) {
	type testScenario struct {
		description string
		arr         []int
		query       int
		expected    int
	}

	tests := []testScenario{
		{
			description: "Search for an element in a sorted array that exists",
			arr:         []int{1, 3, 5, 7, 9, 11},
			query:       7,
			expected:    3,
		},
		{
			description: "Search for an element in a sorted array that does not exist",
			arr:         []int{1, 3, 5, 7, 9, 11},
			query:       4,
			expected:    -1,
		},
		{
			description: "Search in an empty array",
			arr:         []int{},
			query:       5,
			expected:    -1,
		},
		{
			description: "Search in an array with identical elements where the element is present",
			arr:         []int{2, 2, 2, 2, 2},
			query:       2,
			expected:    0,
		},
		{
			description: "Search in an array with a single element that matches the search query",
			arr:         []int{10},
			query:       10,
			expected:    0,
		},
		{
			description: "Search in an array where the query is larger than any element",
			arr:         []int{1, 2, 3, 4, 5},
			query:       10,
			expected:    -1,
		},
		{
			description: "Search in an array where the query is smaller than any element",
			arr:         []int{5, 7, 9, 11},
			query:       1,
			expected:    -1,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			result := interpolationSearch(test.arr, test.query)
			if result != test.expected {
				t.Errorf("Failed: %s\nExpected: %d, Got: %d", test.description, test.expected, result)
			} else {
				t.Logf("Passed: %s\nExpected: %d, Got: %d", test.description, test.expected, result)
			}
		})
	}
}
