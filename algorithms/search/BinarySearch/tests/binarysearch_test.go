package BinarySearch

import "testing"

/*
ROOST_METHOD_HASH=binarySearch_5149337a0e
ROOST_METHOD_SIG_HASH=binarySearch_7d22ad2576


 */
func Testbinarysearch(t *testing.T) {
	tests := []struct {
		description	string
		arr		[]int
		query		int
		expected	int
	}{{description: "Scenario 1: Search for an Element Present in the Array", arr: []int{1, 2, 3, 4, 5}, query: 3, expected: 2}, {description: "Scenario 2: Search for an Element Not Present in the Array", arr: []int{1, 2, 3, 4, 5}, query: 6, expected: -1}, {description: "Scenario 3: Search in an Empty Array", arr: []int{}, query: 1, expected: -1}, {description: "Scenario 4: Search for the First Element", arr: []int{1, 2, 3, 4, 5}, query: 1, expected: 0}, {description: "Scenario 5: Search for the Last Element", arr: []int{1, 2, 3, 4, 5}, query: 5, expected: 4}, {description: "Scenario 6: Search for an Element in a Two-Element Array", arr: []int{1, 2}, query: 2, expected: 1}, {description: "Scenario 7: Search for an Element Not Present in a Two-Element Array", arr: []int{1, 2}, query: 3, expected: -1}}
	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			t.Logf("Running test: %s", test.description)
			result := binarySearch(test.arr, test.query)
			if result != test.expected {
				t.Errorf("Failed test - %s: expected %d, got %d", test.description, test.expected, result)
			} else {
				t.Logf("Success: expected %d, got %d", test.expected, result)
			}
		})
	}
}

/*
ROOST_METHOD_HASH=binarySearch1_e1c20abb75
ROOST_METHOD_SIG_HASH=binarySearch1_c36ccb5b2b


 */
func Testbinarysearch1(t *testing.T) {
	tests := []struct {
		name		string
		arr		[]int
		query		int
		expected	int
	}{{name: "Query Element is Present in Array", arr: []int{1, 3, 5, 7, 9}, query: 5, expected: 2}, {name: "Query Element is Not Present in Array", arr: []int{2, 4, 6, 8, 10}, query: 5, expected: -1}, {name: "Query is Below Range of Array Elements", arr: []int{2, 3, 4, 5}, query: 1, expected: -1}, {name: "Query is Above Range of Array Elements", arr: []int{3, 4, 5, 6}, query: 10, expected: -1}, {name: "Empty Array", arr: []int{}, query: 3, expected: -1}, {name: "Single Element Array - Element Present", arr: []int{7}, query: 7, expected: 0}, {name: "Single Element Array - Element Absent", arr: []int{8}, query: 10, expected: -1}, {name: "Middle Element is the Target", arr: []int{4, 6, 8, 10, 12}, query: 8, expected: 2}, {name: "Handling Duplicates - Query Matching Multiple", arr: []int{1, 2, 3, 3, 3, 4, 5}, query: 3, expected: 2}, {name: "Larger Array", arr: func() []int {
		arr := make([]int, 1000)
		for i := 0; i < 1000; i++ {
			arr[i] = i
		}
		return arr
	}(), query: 678, expected: 678}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log("Executing test case:", tt.name)
			got := binarySearch1(tt.arr, tt.query)
			if got != tt.expected {
				t.Errorf("binarySearch1(%v, %d) = %d; expected %d", tt.arr, tt.query, got, tt.expected)
			} else {
				t.Logf("Test succeeded as expected, result: %d", got)
			}
		})
	}
}

