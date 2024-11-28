package JumpSearch

import (
	"testing"
	"fmt"
	"bytes"
)

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95


 */
func Testjumpsearch418(t *testing.T) {
	type test struct {
		name		string
		array		[]int
		query		int
		expected	int
	}
	tests := []test{{name: "Search for an Element in a Small Array", array: []int{1, 3, 5, 7}, query: 5, expected: 2}, {name: "Search for an Element at the Start of the Array", array: []int{2, 4, 6, 8}, query: 2, expected: 0}, {name: "Element Not Present in the Array", array: []int{1, 4, 6, 9}, query: 3, expected: -1}, {name: "Search for an Element in a Large Array", array: generateLargeArray(10000), query: 4500, expected: 4499}, {name: "Search in an Array with Duplicate Elements", array: []int{1, 3, 5, 5, 7}, query: 5, expected: 2}, {name: "Empty Array Check", array: []int{}, query: 100, expected: -1}, {name: "Element at the End of the Array", array: []int{1, 3, 5, 9, 11}, query: 11, expected: 4}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := jumpSearch(tt.array, tt.query)
			if result != tt.expected {
				t.Errorf("FAILED: %s | Expected: %d, Got: %d", tt.name, tt.expected, result)
			} else {
				t.Logf("PASSED: %s | Expected and got: %d", tt.name, tt.expected)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=jumpSearch_cf44201f2c
ROOST_METHOD_SIG_HASH=jumpSearch_45970a9b95


 */
func generateLargeArray(length int) []int {
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = i
	}
	return arr
}

