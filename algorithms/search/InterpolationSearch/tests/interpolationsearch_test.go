package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd


 */
func Testinterpolationsearch84(t *testing.T) {
	type test struct {
		name        string
		arr         []int
		query       int
		expected    int
		description string
	}

	tests := []test{
		{
			name:        "Scenario 1: Element exists",
			arr:         []int{10, 20, 30, 40, 50},
			query:       30,
			expected:    2,
			description: "Verify that the function correctly identifies the index of an existing element within a sorted array.",
		},
		{
			name:        "Scenario 2: Element does not exist",
			arr:         []int{10, 20, 30, 40, 50},
			query:       35,
			expected:    -1,
			description: "Test that the function returns -1 when the queried element isn't in the array.",
		},
		{
			name:        "Scenario 3: Empty array",
			arr:         []int{},
			query:       10,
			expected:    -1,
			description: "Confirm the function handles an empty array input gracefully.",
		},
		{
			name:        "Scenario 4: Identical elements, element present",
			arr:         []int{5, 5, 5, 5, 5},
			query:       5,
			expected:    0,
			description: "Validate the function's behavior in an array composed entirely of the same value.",
		},
		{
			name:        "Scenario 5: Single element array matching query",
			arr:         []int{42},
			query:       42,
			expected:    0,
			description: "Check the function's response when the array has only one element which is the query value.",
		},
		{
			name:        "Scenario 6: Query larger than any array element",
			arr:         []int{10, 20, 30, 40, 50},
			query:       60,
			expected:    -1,
			description: "Assess the function's ability to identify an out-of-bounds query beyond the array's range.",
		},
		{
			name:        "Scenario 7: Query smaller than any array element",
			arr:         []int{10, 20, 30, 40, 50},
			query:       5,
			expected:    -1,
			description: "Determine the function's behavior with a query less than the smallest element in the array.",
		},
	}

	for _, tc := range tests {
		t.Log(tc.description)
		t.Run(tc.name, func(t *testing.T) {
			result := interpolationSearch(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("failed %s: expected %d, got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("success %s: query %d found at index %d", tc.name, tc.query, result)
			}
		})
	}
}

