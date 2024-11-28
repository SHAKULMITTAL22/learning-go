package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721


 */
func Testlinearsearch687(t *testing.T) {

	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
	}{

		{
			name:     "NonExistentElement",
			arr:      []int{1, 2, 3, 4, 5},
			query:    100,
			expected: -1,
		},

		{
			name:     "EmptyArray",
			arr:      []int{},
			query:    42,
			expected: -1,
		},

		{
			name:     "MultipleIdenticalElements",
			arr:      []int{2, 3, 4, 3, 5},
			query:    3,
			expected: 1,
		},

		{
			name:     "SortedArray",
			arr:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			query:    7,
			expected: 6,
		},

		{
			name:     "NegativeNumbers",
			arr:      []int{-5, -2, 0, 2, 5},
			query:    -2,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := linearSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test %s FAILED: expected %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("Test %s PASSED: correctly returned %d", tt.name, result)
			}
		})
	}

}

