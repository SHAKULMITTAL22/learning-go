package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd


 */
func Testinterpolationsearch579(t *testing.T) {
	type testCase struct {
		name        string
		arr         []int
		query       int
		expected    int
		description string
	}

	tests := []testCase{
		{
			name:        "Test 01 - Search for an element in a sorted array that exists",
			arr:         []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:       5,
			expected:    4,
			description: "Verify that the function correctly identifies the index of an existing element within a sorted array.",
		},
		{
			name:        "Test 02 - Search for an element in a sorted array that does not exist",
			arr:         []int{1, 2, 3, 4, 6, 7, 8, 9, 10},
			query:       5,
			expected:    -1,
			description: "Ensure the function returns -1 when the queried element isn't in the array.",
		},
		{
			name:        "Test 03 - Search in an empty array",
			arr:         []int{},
			query:       5,
			expected:    -1,
			description: "Confirm the function handles an empty array input gracefully.",
		},
		{
			name:        "Test 04 - Search in an array with identical elements where the element is present",
			arr:         []int{5, 5, 5, 5, 5},
			query:       5,
			expected:    0,
			description: "Validate the function's behavior in an array composed entirely of the same value.",
		},
		{
			name:        "Test 05 - Search in an array with a single element that matches the search query",
			arr:         []int{5},
			query:       5,
			expected:    0,
			description: "Check the function's response when the array has only one element which is the query value.",
		},
		{
			name:        "Test 06 - Search in an array where the query is larger than any element",
			arr:         []int{1, 2, 3, 4, 5},
			query:       10,
			expected:    -1,
			description: "Assess the function's ability to identify an out-of-bounds query beyond the array's range.",
		},
		{
			name:        "Test 07 - Search in an array where the query is smaller than any element",
			arr:         []int{5, 6, 7, 8, 9},
			query:       1,
			expected:    -1,
			description: "Determine the function's behavior with a query less than the smallest element in the array.",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Log(tc.description)
			result := interpolationSearch(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("Test: %s - Failed, expected %d, got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Test: %s - Success, correctly returned %d", tc.name, result)
			}
		})
	}
}

