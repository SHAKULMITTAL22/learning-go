package JumpSearch

import "testing"

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95
*/
func Testjumpsearch140(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
	}{
		{
			name:     "Search Element Exists at Beginning",
			arr:      []int{3, 5, 7, 9, 11},
			query:    3,
			expected: 0,
		},
		{
			name:     "Search Element Exists in Middle",
			arr:      []int{2, 4, 6, 8, 10},
			query:    6,
			expected: 2,
		},
		{
			name:     "Search Element Exists at End",
			arr:      []int{1, 3, 5, 7, 9},
			query:    9,
			expected: 4,
		},
		{
			name:     "Search Element Not Present",
			arr:      []int{1, 2, 3, 4, 5},
			query:    6,
			expected: -1,
		},
		{
			name:     "Empty Array",
			arr:      []int{},
			query:    10,
			expected: -1,
		},
		{
			name:     "Single Element Matching",
			arr:      []int{7},
			query:    7,
			expected: 0,
		},
		{
			name:     "Single Element Non-Matching",
			arr:      []int{5},
			query:    8,
			expected: -1,
		},
		{
			name:     "Large Array with Element Present",
			arr:      createLargeArray(1, 10000),
			query:    4567,
			expected: 4566,
		},
		{
			name:     "Large Array with Element Absent",
			arr:      createLargeArray(1, 10000),
			query:    10001,
			expected: -1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := jumpSearch(test.arr, test.query)
			if result != test.expected {
				t.Errorf("Test %s failed. Expected %d, got %d", test.name, test.expected, result)
			} else {
				t.Logf("Test %s succeeded. Correctly returned %d", test.name, result)
			}
		})
	}
}

func createLargeArray(start, end int) []int {
	arr := make([]int, end-start+1)
	for i := range arr {
		arr[i] = start + i
	}
	return arr
}
