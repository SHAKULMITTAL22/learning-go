package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721


 */
func Testlinearsearch803(t *testing.T) {
	type testCase struct {
		arr		[]int
		query		int
		expected	int
		desc		string
	}
	tests := []testCase{{arr: []int{}, query: 0, expected: -1, desc: "Scenario 1: Searching in an Empty Array"}, {arr: []int{5}, query: 5, expected: 0, desc: "Scenario 2: Search with Single Element Matching Query"}, {arr: []int{3}, query: 5, expected: -1, desc: "Scenario 3: Search with Single Element Not Matching Query"}, {arr: []int{1, 2, 3, 4, 5}, query: 1, expected: 0, desc: "Scenario 4: Search for the First Element in a Multi-Element Array"}, {arr: []int{1, 2, 3, 4, 5}, query: 5, expected: 4, desc: "Scenario 5: Search for the Last Element in a Multi-Element Array"}, {arr: []int{1, 2, 3, 4, 5}, query: 6, expected: -1, desc: "Scenario 6: Search for a Non-Existing Element"}, {arr: []int{1, 2, 3, 3, 4, 5}, query: 3, expected: 2, desc: "Scenario 7: Search for a Duplicate Element"}}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			result := linearSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Failed %s: expected %d, got %d", tt.desc, tt.expected, result)
			} else {
				t.Logf("Passed %s: expected %d, got %d", tt.desc, tt.expected, result)
			}
		})
	}
}

