package LinearSearch

import "testing"

/*
ROOST_METHOD_HASH=linearSearch_076b2ca8c0
ROOST_METHOD_SIG_HASH=linearSearch_12fac2e721
*/
func Testlinearsearch879(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		query    int
		expected int
	}{{name: "Scenario 1: Searching in an Empty Array", arr: []int{}, query: 5, expected: -1}, {name: "Scenario 2: Search with Single Element Matching Query", arr: []int{5}, query: 5, expected: 0}, {name: "Scenario 3: Search with Single Element Not Matching Query", arr: []int{4}, query: 5, expected: -1}, {name: "Scenario 4: Search for the First Element in a Multi-Element Array", arr: []int{5, 3, 7, 9, 10}, query: 5, expected: 0}, {name: "Scenario 5: Search for the Last Element in a Multi-Element Array", arr: []int{3, 7, 9, 10, 5}, query: 5, expected: 4}, {name: "Scenario 6: Search for a Non-Existing Element", arr: []int{1, 2, 3, 4, 6}, query: 5, expected: -1}, {name: "Scenario 7: Search for a Duplicate Element", arr: []int{1, 5, 5, 7, 9}, query: 5, expected: 1}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := linearSearch(tt.arr, tt.query)
			if result != tt.expected {
				t.Errorf("Test case %q failed: expected %d, got %d", tt.name, tt.expected, result)
			} else {
				t.Logf("Test case %q passed: got expected result %d", tt.name, tt.expected)
			}
		})
	}
}
