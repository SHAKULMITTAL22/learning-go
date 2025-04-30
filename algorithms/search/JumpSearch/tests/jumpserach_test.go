package JumpSearch

import (
	"testing"
	"fmt"
)

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95


 */
func Testjumpsearch864(t *testing.T) {
	type testCase struct {
		description string
		arr         []int
		query       int
		expected    int
	}

	cases := []testCase{
		{
			description: "Search Element Exists at Beginning",
			arr:         []int{3, 5, 7, 9, 11},
			query:       3,
			expected:    0,
		},
		{
			description: "Search Element Exists in Middle",
			arr:         []int{2, 4, 6, 8, 10},
			query:       6,
			expected:    2,
		},
		{
			description: "Search Element Exists at End",
			arr:         []int{1, 3, 5, 7, 9},
			query:       9,
			expected:    4,
		},
		{
			description: "Search Element Not Present",
			arr:         []int{1, 2, 3, 4, 5},
			query:       6,
			expected:    -1,
		},
		{
			description: "Empty Array",
			arr:         []int{},
			query:       10,
			expected:    -1,
		},
		{
			description: "Single Element Matching",
			arr:         []int{7},
			query:       7,
			expected:    0,
		},
		{
			description: "Single Element Non-Matching",
			arr:         []int{5},
			query:       8,
			expected:    -1,
		},
		{
			description: "Large Array with Element Present",
			arr:         generateLargeArray(10000),
			query:       4567,
			expected:    4566,
		},
		{
			description: "Large Array with Element Absent",
			arr:         generateLargeArray(10000),
			query:       10001,
			expected:    -1,
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			t.Logf("Executing test: %s", c.description)
			result := jumpSearch(c.arr, c.query)
			if result != c.expected {
				t.Errorf("Test failed: %s. Expected %d, got %d", c.description, c.expected, result)
			} else {
				t.Logf("Test passed: %s", c.description)
			}
		})
	}
}

func generateLargeArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1
	}
	return arr
}

