package JumpSearch

import "testing"

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95


 */
func Testjumpsearch825(t *testing.T) {

	type testData struct {
		name     string
		array    []int
		query    int
		expected int
	}

	tests := []testData{
		{name: "Search Element Exists at Beginning", array: []int{3, 5, 7, 9, 11}, query: 3, expected: 0},
		{name: "Search Element Exists in Middle", array: []int{2, 4, 6, 8, 10}, query: 6, expected: 2},
		{name: "Search Element Exists at End", array: []int{1, 3, 5, 7, 9}, query: 9, expected: 4},
		{name: "Search Element Not Present", array: []int{1, 2, 3, 4, 5}, query: 6, expected: -1},
		{name: "Empty Array", array: []int{}, query: 10, expected: -1},
		{name: "Single Element Matching", array: []int{7}, query: 7, expected: 0},
		{name: "Single Element Non-Matching", array: []int{5}, query: 8, expected: -1},
		{name: "Large Array with Element Present", array: generateOrderedArray(1, 10000), query: 4567, expected: 4566},
		{name: "Large Array with Element Absent", array: generateOrderedArray(1, 10000), query: 10001, expected: -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Logf("Running test scenario: %s", tc.name)
			result := jumpSearch(tc.array, tc.query)
			if result != tc.expected {
				t.Errorf("FAIL: %s - Expected %d, got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("PASS: %s", tc.name)
			}
		})
	}
}

func generateOrderedArray(start, end int) []int {

	arr := make([]int, end-start+1)
	for i := range arr {
		arr[i] = start + i
	}
	return arr
}

