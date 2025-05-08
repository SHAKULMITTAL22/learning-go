package InterpolationSearch

import (
	"testing"
)

func TestInterpolationSearch_23d9bb25dd(t *testing.T) {
	testCases := []struct {
		name   string
		arr    []int
		query  int
		result int
	}{
		{
			name:   "Test case 1: Normal case",
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:  5,
			result: 4,
		},
		{
			name:   "Test case 2: Query not in array",
			arr:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:  11,
			result: -1,
		},
		{
			name:   "Test case 3: Empty array",
			arr:    []int{},
			query:  5,
			result: -1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := InterpolationSearchFunc(tc.arr, tc.query)

			if result != tc.result {
				t.Errorf("InterpolationSearchFunc(%v, %d) = %d; want %d", tc.arr, tc.query, result, tc.result)
			} else {
				t.Logf("Test %s passed", tc.name)
			}
		})
	}
}

func InterpolationSearchFunc(arr []int, query int) int {
	// Find indexes of two corners
	lo := 0
	hi := len(arr) - 1

	// Since array is sorted, an element present in array must be in range defined by corner
	for lo <= hi && query >= arr[lo] && query <= arr[hi] {
		// Probing the position with keeping uniform distribution in mind.
		midIndex := lo + (query-arr[lo])*(hi-lo)/(arr[hi]-arr[lo])
		midItem := arr[midIndex]

		// Target is found
		if midItem == query {
			return midIndex
		} else if midItem < query {
			// Target is in upper part
			lo = midIndex + 1
		} else if midItem > query {
			// Target is in lower part
			hi = midIndex - 1
		}
	}

	return -1
}
