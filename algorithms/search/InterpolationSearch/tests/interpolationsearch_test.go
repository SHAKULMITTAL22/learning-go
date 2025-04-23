package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd
*/
func Testinterpolationsearch567(t *testing.T) {

	type testCase struct {
		name     string
		arr      []int
		query    int
		expected int
	}

	tests := []testCase{
		{
			name:     "element exists",
			arr:      []int{1, 2, 3, 4, 5, 6},
			query:    4,
			expected: 3,
		},
		{
			name:     "element does not exist",
			arr:      []int{1, 2, 3, 4, 5, 6},
			query:    7,
			expected: -1,
		},
		{
			name:     "empty array",
			arr:      []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "identical elements, element present",
			arr:      []int{7, 7, 7, 7, 7},
			query:    7,
			expected: 0,
		},
		{
			name:     "single matching element",
			arr:      []int{10},
			query:    10,
			expected: 0,
		},
		{
			name:     "query larger than any element",
			arr:      []int{1, 2, 3, 4, 5},
			query:    10,
			expected: -1,
		},
		{
			name:     "query smaller than any element",
			arr:      []int{10, 20, 30, 40, 50},
			query:    5,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Running test case: %s", tt.name)
			result := interpolationSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test failed: %s, expected %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("Test succeeded: %s, correctly returned %d", tt.name, result)
			}
		})
	}

}
