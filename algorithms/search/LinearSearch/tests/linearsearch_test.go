package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721
*/
func Testlinearsearch721(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
	}{{name: "Empty array", arr: []int{}, query: 5, expected: -1}, {name: "Single element matching query", arr: []int{7}, query: 7, expected: 0}, {name: "Single element not matching query", arr: []int{7}, query: 5, expected: -1}, {name: "First element in multi-element array", arr: []int{3, 8, 12, 15}, query: 3, expected: 0}, {name: "Last element in multi-element array", arr: []int{3, 8, 12, 15}, query: 15, expected: 3}, {name: "Non-existing element", arr: []int{3, 8, 12}, query: 15, expected: -1}, {name: "Duplicate elements in array", arr: []int{3, 8, 12, 8, 15}, query: 8, expected: 1}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := linearSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test %s failed: got %d, expected %d", tt.name, result, tt.expected)
			} else {
				t.Logf("Test %s succeeded: got %d, expected %d", tt.name, result, tt.expected)
			}
		})
	}
}
