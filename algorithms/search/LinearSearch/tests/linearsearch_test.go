package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721


 */
func Testlinearsearch423(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
	}{
		{
			name:     "Non-Existent Element",
			arr:      []int{1, 2, 3, 4, 5},
			query:    6,
			expected: -1,
		},
		{
			name:     "Empty Array",
			arr:      []int{},
			query:    1,
			expected: -1,
		},
		{
			name:     "Multiple Identical Elements",
			arr:      []int{5, 3, 5, 1, 2},
			query:    5,
			expected: 0,
		},
		{
			name:     "Sorted Array",
			arr:      []int{1, 2, 3, 4, 5},
			query:    4,
			expected: 3,
		},
		{
			name:     "Array with Negative Numbers",
			arr:      []int{-5, -1, 3, 2},
			query:    -1,
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := linearSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test %s failed: expected %v, got %v", tt.name, tt.expected, result)
			} else {
				t.Logf("Test %s succeeded: returned expected %v", tt.name, result)
			}
		})
	}
}

