package JumpSearch

import "testing"

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95
*/
func Testjumpsearch397(t *testing.T) {
	type test struct {
		name     string
		array    []int
		query    int
		expected int
	}
	tests := []test{{name: "Small array, element present", array: []int{1, 3, 5, 7}, query: 5, expected: 2}, {name: "Element at start of array", array: []int{2, 4, 6, 8}, query: 2, expected: 0}, {name: "Element not in array", array: []int{1, 4, 6, 9}, query: 3, expected: -1}, {name: "Large array test", array: createLargeTestArray(10000), query: 9999, expected: 9999}, {name: "Array with duplicate elements", array: []int{1, 3, 5, 5, 7}, query: 5, expected: 2}, {name: "Empty array test", array: []int{}, query: 10, expected: -1}, {name: "Element at end of array", array: []int{1, 3, 5, 9, 11}, query: 11, expected: 4}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := jumpSearch(tc.array, tc.query)
			if result != tc.expected {
				t.Errorf("Test %s failed. Expected %d, got %d", tc.name, tc.expected, result)
			} else {
				t.Logf("Test %s succeeded. Correct index %d returned for query %d", tc.name, result, tc.query)
			}
		})
	}
}

func createLargeTestArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i
	}
	return arr
}
