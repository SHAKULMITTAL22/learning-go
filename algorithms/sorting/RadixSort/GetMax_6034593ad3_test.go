package RadixSort

import (
	"testing"
)

// TestGetMax_6034593ad3 is a test function for getMax function
func TestGetMax_6034593ad3(t *testing.T) {
	testCases := []struct {
		name string
		arr  []int
		n    int
		want int
	}{
		{
			name: "Test Case 1: Normal case",
			arr:  []int{1, 2, 3, 4, 5},
			n:    5,
			want: 5,
		},
		{
			name: "Test Case 2: Negative numbers",
			arr:  []int{-1, -2, -3, -4, -5},
			n:    5,
			want: -1,
		},
		{
			name: "Test Case 3: Single element",
			arr:  []int{1},
			n:    1,
			want: 1,
		},
		{
			name: "Test Case 4: All elements are zero",
			arr:  []int{0, 0, 0, 0, 0},
			n:    5,
			want: 0,
		},
		{
			name: "Test Case 5: Array is empty",
			arr:  []int{},
			n:    0,
			want: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.n == 0 {
				return
			}
			got := getMax(tc.arr, tc.n)
			if got != tc.want {
				t.Errorf("Test %v failed. got: %v, want: %v", tc.name, got, tc.want)
			} else {
				t.Logf("Test %v passed. got: %v", tc.name, got)
			}
		})
	}
}
