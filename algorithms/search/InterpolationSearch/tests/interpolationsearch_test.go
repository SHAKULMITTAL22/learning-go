package InterpolationSearch

import (
	"fmt"
	"testing"
)

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd


 */
func Testinterpolationsearch464(t *testing.T) {
	tests := []struct {
		name           string
		arr            []int
		query          int
		expectedResult int
		description    string
	}{
		{
			name:           "ElementExists",
			arr:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:          6,
			expectedResult: 5,
			description:    "Verify that the function correctly identifies the index of an existing element within a sorted array.",
		},
		{
			name:           "ElementDoesNotExist",
			arr:            []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:          11,
			expectedResult: -1,
			description:    "Test that the function returns -1 when the queried element isn't in the array.",
		},
		{
			name:           "EmptyArray",
			arr:            []int{},
			query:          5,
			expectedResult: -1,
			description:    "Confirm the function handles an empty array input gracefully.",
		},
		{
			name:           "IdenticalElementsPresent",
			arr:            []int{4, 4, 4, 4, 4, 4, 4},
			query:          4,
			expectedResult: 0,
			description:    "Validate the function's behavior in an array composed entirely of the same value.",
		},
		{
			name:           "SingleElementArray",
			arr:            []int{3},
			query:          3,
			expectedResult: 0,
			description:    "Check the function's response when the array has only one element which is the query value.",
		},
		{
			name:           "QueryLargerThanAnyElement",
			arr:            []int{1, 2, 3, 4, 5},
			query:          10,
			expectedResult: -1,
			description:    "Assess the function's ability to identify an out-of-bounds query beyond the array's range.",
		},
		{
			name:           "QuerySmallerThanAnyElement",
			arr:            []int{5, 6, 7, 8, 9},
			query:          4,
			expectedResult: -1,
			description:    "Determine the function's behavior with a query less than the smallest element in the array.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualResult := interpolationSearch(tt.arr, tt.query)
			if actualResult != tt.expectedResult {
				t.Errorf("Test case %s failed: expected %d, got %d. %s", tt.name, tt.expectedResult, actualResult, tt.description)
			} else {
				t.Logf("Test case %s passed. %s", tt.name, tt.description)
			}
		})
	}
}

