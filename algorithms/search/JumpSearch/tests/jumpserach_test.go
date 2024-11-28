package JumpSearch

import "testing"

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95


 */
// Testjumpsearch662 tests the jumpSearch function with various scenarios.
func Testjumpsearch662(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
	}{

		{name: "Element at Beginning", arr: []int{3, 5, 7, 9, 11}, query: 3, expected: 0},

		{name: "Element in Middle", arr: []int{2, 4, 6, 8, 10}, query: 6, expected: 2},

		{name: "Element at End", arr: []int{1, 3, 5, 7, 9}, query: 9, expected: 4},

		{name: "Element Not Present", arr: []int{1, 2, 3, 4, 5}, query: 6, expected: -1},

		{name: "Empty Array", arr: []int{}, query: 10, expected: -1},

		{name: "Single Element Match", arr: []int{7}, query: 7, expected: 0},

		{name: "Single Element Non-Match", arr: []int{5}, query: 8, expected: -1},

		{name: "Large Array Element Present", arr: generateLargeArray(10000), query: 4567, expected: 4566},

		{name: "Large Array Element Absent", arr: generateLargeArray(10000), query: 10001, expected: -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := jumpSearch(tc.arr, tc.query)
			if result != tc.expected {
				t.Errorf("Test %s failed: expected %d, got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s passed", tc.name)
			}
		})
	}
}

// Helper function to generate a large array
// TODO: User must adjust range for different test setups if needed.
func generateLargeArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i + 1
	}
	return arr
}

