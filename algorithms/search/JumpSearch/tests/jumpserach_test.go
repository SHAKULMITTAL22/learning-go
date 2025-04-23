package JumpSearch

import "testing"

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95
*/
func Testjumpsearch248(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
	}{
		{
			name:     "Scenario 1: Search Element Exists at Beginning",
			arr:      []int{3, 5, 7, 9, 11},
			query:    3,
			expected: 0,
		},
		{
			name:     "Scenario 2: Search Element Exists in Middle",
			arr:      []int{2, 4, 6, 8, 10},
			query:    6,
			expected: 2,
		},
		{
			name:     "Scenario 3: Search Element Exists at End",
			arr:      []int{1, 3, 5, 7, 9},
			query:    9,
			expected: 4,
		},
		{
			name:     "Scenario 4: Search Element Not Present",
			arr:      []int{1, 2, 3, 4, 5},
			query:    6,
			expected: -1,
		},
		{
			name:     "Scenario 5: Empty Array",
			arr:      []int{},
			query:    10,
			expected: -1,
		},
		{
			name:     "Scenario 6: Single Element Matching",
			arr:      []int{7},
			query:    7,
			expected: 0,
		},
		{
			name:     "Scenario 7: Single Element Non-Matching",
			arr:      []int{5},
			query:    8,
			expected: -1,
		},
		{
			name: "Scenario 8: Large Array with Element Present",
			arr: func() []int {
				arr := make([]int, 10000)
				for i := 1; i <= 10000; i++ {
					arr[i-1] = i
				}
				return arr
			}(),
			query:    4567,
			expected: 4566,
		},
		{
			name: "Scenario 9: Large Array with Element Absent",
			arr: func() []int {
				arr := make([]int, 10000)
				for i := 1; i <= 10000; i++ {
					arr[i-1] = i
				}
				return arr
			}(),
			query:    10001,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jumpSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Failed %s: expected index %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("Passed %s: returned expected index %d", tt.name, result)
			}
		})
	}
}
