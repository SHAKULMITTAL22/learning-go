package InterpolationSearch

import "testing"

/*
ROOST_METHOD_HASH=interpolationSearch_d0bdb80f3c
ROOST_METHOD_SIG_HASH=interpolationSearch_23d9bb25dd
*/
func Testinterpolationsearch33(t *testing.T) {
	tests := []struct {
		description string
		arr         []int
		query       int
		expected    int
	}{{description: "Scenario 1: Search in a Sorted Array with Present Value", arr: []int{1, 3, 5, 7, 9}, query: 5, expected: 2}, {description: "Scenario 2: Search for a Value Not in the Array", arr: []int{2, 4, 6, 8, 10}, query: 5, expected: -1}, {description: "Scenario 3: Search in an Empty Array", arr: []int{}, query: 1, expected: -1}, {description: "Scenario 4: Search at the Array Boundaries (Start)", arr: []int{10, 20, 30, 40, 50}, query: 10, expected: 0}, {description: "Scenario 4: Search at the Array Boundaries (End)", arr: []int{10, 20, 30, 40, 50}, query: 50, expected: 4}, {description: "Scenario 5: Search with Identical Elements", arr: []int{7, 7, 7, 7, 7}, query: 7, expected: 0}, {description: "Scenario 6: Large Array Search Performance", arr: generateLargeArray(1, 10000), query: 9876, expected: 9875}, {description: "Scenario 7: Search for Minimum Integer Value", arr: []int{-2147483648, 0, 2147483647}, query: -2147483648, expected: 0}, {description: "Scenario 7: Search for Maximum Integer Value", arr: []int{-2147483648, 0, 2147483647}, query: 2147483647, expected: 2}}
	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			result := interpolationSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test failed: %s, expected %d but got %d", tt.description, tt.expected, result)
			} else {
				t.Logf("Test success: %s", tt.description)
			}
		})
	}
}

func generateLargeArray(start, end int) []int {
	size := end - start + 1
	arr := make([]int, size)
	for i := range arr {
		arr[i] = start + i
	}
	return arr
}
